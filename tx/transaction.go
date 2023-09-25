package tx

import (
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/e-chain-net/echain-go-sdk-721/tars-protocol/bcostars"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"math/rand"
	"strings"
	"time"
)


// @todo
func GetBlockLimit(groupId string) (int64, error) {
	return 3333, nil
}

/*
1
*/
func CreateTransactionData(groupId string, chainId string, to string, dataHex string, abiJson string, blockLimit int64) (*bcostars.TransactionData, error) {
	if blockLimit == 0 {
		var err error
		blockLimit, err = GetBlockLimit(groupId)
		if err != nil {
			return nil, err
		}
	}

	if len(abiJson) > 0 {
		// 合约部署
		return &bcostars.TransactionData{
			Version:    0,
			ChainID:    chainId,
			GroupID:    groupId,
			BlockLimit: blockLimit,
			Nonce:      Nonce(),
			//To:    "0x0000000000000000000000000000000000000000",
			Input: HexByte2Int8(common.FromHex(dataHex)),
			Abi:   abiJson,
		}, nil
	}

	// 方法调用
	return &bcostars.TransactionData{
		Version:    0,
		ChainID:    chainId,
		GroupID:    groupId,
		BlockLimit: blockLimit,
		Nonce:      Nonce(),
		To:         strings.ToLower(to),
		//Input:      HexByte2Int8(dataHex),
		Input: HexByte2Int8(common.FromHex(dataHex)),
		Abi:   "",
	}, nil
}
func hash(buf []byte) string {
	// https://github.com/FISCO-BCOS/bcos-crypto/blob/main/bcos-crypto/hash/Keccak256.cpp
	return crypto.Keccak256Hash(buf).Hex()
}
func CalculateTransactionDataHash(txData *bcostars.TransactionData) (string, error) {
	// Keccak256 hash
	buf := codec.NewBuffer()
	//if err := txData.WriteTo(buf); err != nil {
	//	return "", err
	//}
	versionBigEndian := make([]byte,4)
	binary.BigEndian.PutUint32(versionBigEndian,uint32(txData.Version))
	buf.WriteBytes(versionBigEndian)
	buf.WriteBytes([]byte(txData.ChainID))
	buf.WriteBytes([]byte(txData.GroupID))
	blockLimitBigEndian := make([]byte,8)
	binary.BigEndian.PutUint64(blockLimitBigEndian,uint64(txData.BlockLimit))
	buf.WriteBytes(blockLimitBigEndian)
	buf.WriteBytes([]byte(txData.Nonce))
	buf.WriteBytes([]byte(txData.To))
	buf.WriteSliceInt8(txData.Input)
	buf.WriteBytes([]byte(txData.Abi))
	//return HexStringWithPrefix(hash(buf.ToBytes())), nil
	//fmt.Println(buf.ToBytes())
	return hash(buf.ToBytes()), nil
}

func Nonce() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var buf []byte
	for i := 0; i < 32; i++ {
		it := r.Intn(255)
		buf = append(buf, byte(it))
	}

	u256 := uint256.NewInt(0).SetBytes(buf)
	return fmt.Sprint(u256)
}

func PrivateKeyToAddress(privateKey string) (string, error) {
	privKey, err := ParseKeyPairFromPrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	//// 获取地址
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	//fmt.Printf("地址为: %s\n", addr.Hex())

	return strings.ToLower(addr.Hex()), nil
}

func ParseKeyPairFromPrivateKey(privateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKey)
}

func SignTransactionDataHash(privateKey string, txDataHash string) (string, error) {
	privKey, err := ParseKeyPairFromPrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	//Sign a message using the private key.
	hash,err := hexutil.Decode(txDataHash)
	if err != nil {
		return "",err
	}

	signature,err := crypto.Sign(hash,privKey)
	if err != nil{
		return "",err
	}
	return hexutil.Encode(signature),nil
}

func CreateTransaction(from string, txData *bcostars.TransactionData, txDataHash string, signedTxDataHash string, attribute int32) (*bcostars.Transaction, error) {
	return &bcostars.Transaction{
		Data:       *txData,
		DataHash:   HexByte2Int8(common.FromHex(txDataHash)),
		Signature:  HexByte2Int8(common.FromHex(signedTxDataHash)),
		ImportTime: 0,
		Attribute:  attribute,
		ExtraData:  "",
		//Sender:     HexByte2Int8(common.FromHex(strings.ToLower(from))),
	}, nil
}
func EncodeTransaction(tx *bcostars.Transaction) (string, error) {
	buf := codec.NewBuffer()
	if err := tx.WriteTo(buf); err != nil {
		return "", err
	}
	return hexutil.Encode(buf.ToBytes()), nil
}
func CreateSignedTransaction(privateKey string, groupId string, chainId string, to string, dataHex string, abiJson string, blockLimit int64, attribute int32) (txHash string, txHex string, err error) {
	txData, err := CreateTransactionData(groupId, chainId, to, dataHex, abiJson, blockLimit)
	if err != nil {
		return "", "", err
	}
	//fmt.Println("txData", txData)

	txDataHash, err := CalculateTransactionDataHash(txData)
	if err != nil {
		return "", "", err
	}
	//fmt.Println("txDataHash", txDataHash)

	signedTxDataHash, err := SignTransactionDataHash(privateKey, txDataHash)
	if err != nil {
		return "", "", err
	}
	//fmt.Println("signedTxDataHash", signedTxDataHash)

	from, err := PrivateKeyToAddress(privateKey)
	if err != nil {
		return "", "", err
	}
	tx, err := CreateTransaction(from, txData, txDataHash, signedTxDataHash, attribute)
	if err != nil {
		return "", "", err
	}

	_txHex, err := EncodeTransaction(tx)
	if err != nil {
		return "", "", err
	}
	//fmt.Println("_txHex", _txHex)

	return txDataHash, _txHex, nil
}

func HexByte2Int8(dataHex []byte) []int8 {
	var dataHexInt8 []int8
	for _, d := range dataHex {
		dataHexInt8 = append(dataHexInt8, int8(d))
	}
	return dataHexInt8
}

func Int82Byte(data []int8) []byte{
	var dataHexByte []byte
	for _, d := range data {
		dataHexByte = append(dataHexByte, byte(d))
	}
	return dataHexByte
}

func EncodeTransactionDataToHex(tx *bcostars.TransactionData) (string, error) {
	buf := codec.NewBuffer()
	if err := tx.WriteTo(buf); err != nil {
		return "", err
	}

	return hexutil.Encode(buf.ToBytes()), nil
}
func DecodeTransactionDataFromHex(txDataHex string) (*bcostars.TransactionData, error) {
	decode, err := hexutil.Decode(txDataHex)
	if err != nil {
		return nil, err
	}
	buf := codec.NewReader(decode)

	txData := &bcostars.TransactionData{}
	if err := txData.ReadFrom(buf); err != nil {
		return nil, err
	}

	return txData, nil
}

func DecodeTransactionFromHex(txHex string)(*bcostars.Transaction,error){
	decode,err := hexutil.Decode(txHex);
	if err != nil {
		return nil, err
	}
	buf := codec.NewReader(decode)

	tx := &bcostars.Transaction{}
	if err := tx.ReadFrom(buf); err != nil {
		return nil, err
	}

	return tx, nil
}
