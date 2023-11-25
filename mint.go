package main

import (
	"fmt"
	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/crypto"
	"time"
)

func mint(web3Client *web3.Web3, data string) {

	nonce, err := web3Client.Eth.GetNonce(web3Client.Eth.Address(), nil)
	if err != nil {
		fmt.Println("Error fetching nonce:", err)
		return
	}

	tipCap, err := web3Client.Eth.SuggestGasTipCap()
	if err != nil {
		fmt.Println("Error fetching tip cap:", err)
		return
	}

	fee, err := web3Client.Eth.EstimateFee()
	if err != nil {
		fmt.Println("Error fetching estimate:", err)
		return
	}
	transaction, err := web3Client.Eth.SyncSendEIP1559RawTransaction(
		crypto.PubkeyToAddress(web3Client.Eth.GetPrivateKey().PublicKey),
		web3Client.Utils.ToWei("0"),
		nonce,
		28000,
		tipCap,
		fee.MaxFeePerGas,
		[]byte(data),
	)
	if err != nil {
		fmt.Println("Error sending transaction:", err)
		return
	}

	fmt.Println(fmt.Sprintf("%s mint successfully, nonce:%d, hash:%s", time.Now().Format(time.DateTime), nonce, transaction.TxHash))
	//addNonce()
	if transaction.CumulativeGasUsed >= 292901036620 {
		return
	}
}

func newWeb3(rpcURL string, privateKeyStr string, chainId int64) (*web3.Web3, error) {
	web3Client, err := web3.NewWeb3(rpcURL)
	if err != nil {
		fmt.Println("Error connecting to Ethereum client:", err)
		return nil, err
	}

	err = web3Client.Eth.SetAccount(privateKeyStr)
	if err != nil {
		fmt.Println("Error setting account:", err)
		return nil, err
	}

	web3Client.Eth.SetChainId(chainId)
	chainID, err := web3Client.Eth.ChainID()
	if err != nil {
		fmt.Println("Error fetching chain ID:", err)
		return nil, err
	}
	fmt.Println("chain ID: ", chainID)
	fmt.Println("address: ", crypto.PubkeyToAddress(web3Client.Eth.GetPrivateKey().PublicKey))

	return web3Client, nil

}

func main() {
	privateKey := "xxx"

	//pols
	rpc := "https://polygon.drpc.org"
	chainId := int64(137)
	data := "data:,{\"p\":\"prc-20\",\"op\":\"mint\",\"tick\":\"poli\",\"amt\":\"1000\"}"
	loopCount := 1000000


	
	fmt.Println("https://okx.com/join/1992197 OKX的web3钱包支持60+公链，提供多链交易，提币不用等，走我链接注册可以获取任何的免费技术指导。")
	w, err := newWeb3(rpc, privateKey, chainId)
	if err != nil {
		fmt.Println("Error creating web3: ", err)
		return
	}

	for i := 0; i < loopCount; i++ {
		mint(w, data)
	}
}
