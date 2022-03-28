package zetaclient

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/zetaclient/config"
	"math/big"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/zetacore/x/zetacore/types"
)

type TxStatus int64

const (
	Unprocessed TxStatus = iota
	Pending
	Mined
	Confirmed
	Error
)

type CoreObserver struct {
	sendQueue  []*types.Send
	sendMap    map[string]*types.Send // send.Index => send
	sendStatus map[string]TxStatus    // send.Index => status
	recvMap    map[string]*types.Receive
	recvStatus map[string]TxStatus
	bridge     *MetachainBridge
	signerMap  map[common.Chain]*Signer
	clientMap  map[common.Chain]*ChainObserver
	lock       sync.Mutex
	httpServer *HTTPServer
}

func NewCoreObserver(bridge *MetachainBridge, signerMap map[common.Chain]*Signer, clientMap map[common.Chain]*ChainObserver, server *HTTPServer) *CoreObserver {
	co := CoreObserver{}
	co.bridge = bridge
	co.signerMap = signerMap
	co.sendQueue = make([]*types.Send, 0)
	co.sendMap = make(map[string]*types.Send)
	co.sendStatus = make(map[string]TxStatus)
	co.recvMap = make(map[string]*types.Receive)
	co.recvStatus = make(map[string]TxStatus)
	co.clientMap = clientMap
	co.httpServer = server
	return &co
}

func (co *CoreObserver) MonitorCore() {
	myid := co.bridge.keys.GetSignerInfo().GetAddress().String()
	log.Info().Msgf("MonitorCore started by signer %s", myid)
	go co.observeSend()

	go co.observeReceive()

	// Pull items from queue
	go co.processOutboundQueue()

}

func (co *CoreObserver) observeSend() {
	for {
		zetaHeight, err := co.bridge.GetMetaBlockHeight()
		if err != nil {
			log.Warn().Err(err).Msgf("GetMetaBlockHeight error")
			continue
		}

		sendList, err := co.bridge.GetAllPendingSend()
		if err != nil {
			fmt.Println("error requesting sends from zetacore")
			time.Sleep(5 * time.Second)
			continue
		}

		sendListMap := make(map[string]bool)
		for _, send := range sendList {
			sendListMap[send.Index] = true
		}

		// clean up sendMap and sendStatus and sendQueue
		co.lock.Lock()
		for key, _ := range co.sendMap {
			if _, ok := sendListMap[key]; !ok {
				log.Info().Msgf("removing %s from sendMap", key)
				delete(co.sendMap, key)
				delete(co.sendStatus, key)
			}
		}
		co.lock.Unlock()

		for _, send := range sendList {
			if send.Status == types.SendStatus_Finalized || send.Status == types.SendStatus_Revert {
				co.lock.Lock()
				oldSend, found := co.sendMap[send.Index]
				co.lock.Unlock()
				if !found || oldSend.Status != send.Status { // new send or send status changed; needs to process
					co.lock.Lock()
					if !found {
						log.Debug().Msgf("New send queued with finalized block %d", send.FinalizedMetaHeight)
					}
					if found && oldSend.Status != send.Status {
						log.Debug().Msgf("Old send status updated from %s to %s", types.SendStatus_name[int32(oldSend.Status)], types.SendStatus_name[int32(send.Status)])
					}
					co.sendMap[send.Index] = send
					co.sendQueue = append(co.sendQueue, send)
					co.sendStatus[send.Index] = Unprocessed
					co.lock.Unlock()
				} else {
					co.lock.Lock()
					status, found := co.sendStatus[send.Index]
					co.lock.Unlock()
					if !found {
						log.Error().Msgf("status of send: %s not found", send.Index)
						continue
					}
					// the send is not successfully process; re-process is needed
					if zetaHeight-send.FinalizedMetaHeight > config.TIMEOUT_THRESHOLD_FOR_RETRY &&
						(zetaHeight-send.FinalizedMetaHeight)%config.TIMEOUT_THRESHOLD_FOR_RETRY == 0 &&
						status != Unprocessed {
						log.Warn().Msgf("Zeta block %d: Timeout send: sendHash %s chain %s nonce %d; re-process...", zetaHeight, send.Index, send.ReceiverChain, send.Nonce)
						co.lock.Lock()
						co.sendStatus[send.Index] = Unprocessed
						co.lock.Unlock()
					}
				}
			} else if send.Status == types.SendStatus_Mined || send.Status == types.SendStatus_Reverted || send.Status == types.SendStatus_Aborted {
				co.lock.Lock()
				send, found := co.sendMap[send.Index]
				delete(co.sendMap, send.Index)
				co.lock.Unlock()
				if found {
					co.lock.Lock()
					if co.sendStatus[send.Index] != Mined {
						log.Info().Msgf("Send status changed to Mined")
					}
					co.sendStatus[send.Index] = Mined
					co.lock.Unlock()
				}
			}

		}
		time.Sleep(5 * time.Second)
	}
}

func (co *CoreObserver) observeReceive() {
	//for {
	//	recvList, err := co.bridge.GetAllReceive()
	//	if err != nil {
	//		fmt.Println("error requesting receives from zetacore")
	//		time.Sleep(5 * time.Second)
	//		continue
	//	}
	//	for _, recv := range recvList {
	//		if recv.Status == common.ReceiveStatus_Created {
	//			if _, found := co.recvMap[recv.Index]; !found {
	//				co.lock.Lock()
	//				log.Debug().Msgf("New recv created with finalized block")
	//				co.recvMap[recv.Index] = recv
	//				co.recvStatus[recv.Index] = Unprocessed
	//				chain, err := common.ParseChain(recv.Chain)
	//				if err != nil {
	//					fmt.Printf("recv chain invalid: %s\n", recv.Chain)
	//					continue
	//				}
	//				co.clientMap[chain].AddTxToWatchList(recv.OutTxHash, recv.SendHash)
	//				co.lock.Unlock()
	//			}
	//		}
	//	}
	//	time.Sleep(5 * time.Second)
	//}
}

func (co *CoreObserver) processOutboundQueue() {
	myid := co.bridge.keys.GetSignerInfo().GetAddress().String()
	for {
		startTime := time.Now()
		for idx, send := range co.sendQueue {
			co.lock.Lock()
			nPendingSend := len(co.sendMap)
			status, ok := co.sendStatus[send.Index]
			co.lock.Unlock()
			co.httpServer.mu.Lock()
			co.httpServer.pendingTx = uint64(nPendingSend)
			co.httpServer.mu.Unlock()
			if status != Unprocessed || !ok {
				continue
			}

			log.Info().Msgf("# of Pending send %d", nPendingSend)
			amount, ok := new(big.Int).SetString(send.MMint, 10)
			if !ok {
				log.Error().Msg("error converting MBurnt to big.Int")
				//time.Sleep(5 * time.Second)
				continue
			}

			var to ethcommon.Address
			var err error
			var toChain common.Chain
			if send.Status == types.SendStatus_Revert {
				to = ethcommon.HexToAddress(send.Sender)
				toChain, err = common.ParseChain(send.SenderChain)
				log.Info().Msgf("Abort: reverting inbound")
			} else {
				to = ethcommon.HexToAddress(send.Receiver)
				toChain, err = common.ParseChain(send.ReceiverChain)
			}
			if err != nil {
				log.Err(err).Msg("ParseChain fail; skip")
				time.Sleep(5 * time.Second)
				continue
			}

			processed, err := co.clientMap[toChain].IsSendOutTxProcessed(send.Index)
			if err != nil {
				log.Err(err).Msg("IsSendOutTxProcessed error")
				//time.Sleep(1 * time.Second)
				continue
			}
			if processed {
				log.Info().Msgf("sendHash %s already processed; skip it", send.Index)
				time.Sleep(1 * time.Second)
				continue
			}

			signer := co.signerMap[toChain]
			//message := []byte(send.Message)
			message, err := base64.StdEncoding.DecodeString(send.Message)
			if err != nil {
				log.Err(err).Msgf("decode send.Message %s error", send.Message)
			}
			var gasLimit uint64 = 250_000

			log.Info().Msgf("chain %s minting %d to %s, nonce %d, finalized %d, in queue %d", toChain, amount, to.Hex(), send.Nonce, send.FinalizedMetaHeight, idx)
			sendHash, err := hex.DecodeString(send.Index[2:]) // remove the leading 0x
			if err != nil || len(sendHash) != 32 {
				log.Err(err).Msgf("decode sendHash %s error", send.Index)
			}
			var sendhash [32]byte
			copy(sendhash[:32], sendHash[:32])
			gasprice, ok := new(big.Int).SetString(send.GasPrice, 10)
			if !ok {
				log.Err(err).Msgf("cannot convert gas price  %s ", send.GasPrice)
			}
			var tx *ethtypes.Transaction
			//func (signer *Signer) SignOutboundTx(sender ethcommon.Address, srcChainID uint16, to ethcommon.Address, amount *big.Int,, gasLimit uint64, message []byte, sendHash [32]byte, nonce uint64, gasPrice *big.Int) (*ethtypes.Transaction, error) {

			srcChainID := config.Chains[send.SenderChain].ChainID
			tx, err = signer.SignOutboundTx(ethcommon.HexToAddress(send.Sender), srcChainID, to, amount, gasLimit, message, sendhash, send.Nonce, gasprice)
			if err != nil {
				log.Warn().Err(err).Msgf("SignOutboundTx error: nonce %d chain %s", send.Nonce, send.ReceiverChain)
				break
			}

			outTxHash := tx.Hash().Hex()
			log.Info().Msgf("nonce %d, sendHash: %s, outTxHash %s signer %s", send.Nonce, send.Index[:6], outTxHash, myid)

			if myid == send.Signers[send.Broadcaster] || myid == send.Signers[int(send.Broadcaster+1)%len(send.Signers)] {
				log.Info().Msgf("broadcasting tx %s to chain %s: mint amount %d, nonce %d", outTxHash, toChain, amount, send.Nonce)
				err = signer.Broadcast(tx)
				if err != nil {
					log.Err(err).Msgf("Broadcast error: nonce %d chain %s", send.Nonce, toChain)
				}
			}

			co.lock.Lock()
			_, ok = co.sendStatus[send.Index]
			if ok {
				co.sendStatus[send.Index] = Pending
			}
			co.lock.Unlock()
			_, err = co.bridge.PostReceiveConfirmation(send.Index, outTxHash, 0, amount.String(), common.ReceiveStatus_Created, send.ReceiverChain)
			if err != nil {
				log.Err(err).Msgf("PostReceiveConfirmation of just created receive")
			}
			co.clientMap[toChain].AddTxToWatchList(outTxHash, send.Index)

			break

		}
		endTime := time.Now()
		if endTime.Sub(startTime).Milliseconds() < 1000 {
			time.Sleep(time.Duration(1000-endTime.Sub(startTime).Milliseconds()) * time.Millisecond)
		}
	}
}