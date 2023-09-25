package sdk

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/e-chain-net/echain-go-sdk-721/tx"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct{
	Address string
	Private string
}


func NewAccount() Account{
	privKey, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)

	// Convert the private key to a 32-byte hexadecimal string
	privBytes := privKey.D.Bytes()
	privHex := hex.EncodeToString(privBytes)
	address,_ := tx.PrivateKeyToAddress(privHex)

	return Account{
		Address: address,
		Private: privHex,
	}
}
