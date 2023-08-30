package sdk

import (
	"github.com/e-chain-net/echain-go-sdk-721/abi"
	tx2 "github.com/e-chain-net/echain-go-sdk-721/tx"
)

const BLOCK_LIMIT_GAP int64 = 900

func SignMint(toAddress string,tokenId string,contractAddress string,privateHex string,blockNumber int64) (txHash string, txHex string, err error){
	input := abi.EncodeMint(toAddress,tokenId)
	blockLimit := blockNumber + BLOCK_LIMIT_GAP
	return tx2.CreateSignedTransaction(privateHex,"group0","chain0",contractAddress,input,"",blockLimit,0)
}

func SignTransferFrom(fromAddress string,toAddress string,tokenId string,contractAddress string,privateHex string,blockNumber int64) (txHash string, txHex string, err error){
	input := abi.EncodeTransferFrom(fromAddress,toAddress,tokenId)
	blockLimit := blockNumber + BLOCK_LIMIT_GAP
	return tx2.CreateSignedTransaction(privateHex,"group0","chain0",contractAddress,input,"",blockLimit,0)
}

func SignBurn(tokenId string,contractAddress string,privateHex string,blockNumber int64) (txHash string, txHex string, err error){
	input := abi.EncodeBurn(tokenId)
	blockLimit := blockNumber + BLOCK_LIMIT_GAP
	return tx2.CreateSignedTransaction(privateHex,"group0","chain0",contractAddress,input,"",blockLimit,0)
}