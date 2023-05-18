package sdk

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/e-chain-net/echain-go-sdk-721/internal/crypto"
)

type Account struct{
	Address string
	Private string
}


func NewAccount() Account{
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// Convert the private key to a 32-byte hexadecimal string
	privBytes := privKey.D.Bytes()
	privHex := hex.EncodeToString(privBytes)
	address,_ := PrivateKeyToAddress(privHex)

	return Account{
		Address: address,
		Private: privHex,
	}
}

func PrivateKeyToAddress(privateKey string) (string, error) {
	privKey, err := ParseKeyPairFromPrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	//// 获取公钥并去除头部0x04
	//compressed := privKey.PubKey().SerializeUncompressed()[1:]
	//fmt.Printf("公钥为x: %s\n", hex.EncodeToString(compressed))

	//// 获取地址
	addr := crypto.PubkeyToAddress(*privKey.PubKey())
	//fmt.Printf("地址为: %s\n", addr.Hex())

	return addr.Hex(), nil
}

func ParseKeyPairFromPrivateKey(privateKey string) (*secp256k1.PrivateKey, error) {
	// Decode a hex-encoded private key.
	pkBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	privKey := secp256k1.PrivKeyFromBytes(pkBytes)

	return privKey, nil
}