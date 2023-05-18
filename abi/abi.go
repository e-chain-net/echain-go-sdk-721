package abi

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"math"
	"math/big"
	"strings"
)

//vscode- codegpt 生成的代码，不需要go-ethereum依赖

//func IsAddress(address string) bool {
//	addr := common.HexToAddress(address)
//
//	return addr.Hex() == address || addr.String() == address
//}
func IsZeroPrefixed(address string) bool {
	if len(address) <= 2 {
		return false
	}

	return address[0:2] == "0x"
}

func ToBn(value interface{}) *big.Int {
	switch val := value.(type) {
	case string:
		if strings.HasPrefix(val, "0x") || strings.HasPrefix(val, "0X") {
			val = val[2:]
			i, ok := new(big.Int).SetString(val, 16)
			if !ok {
				i = new(big.Int)
			}
			return i
		} else {
			i, ok := new(big.Int).SetString(val, 10)
			if !ok {
				i = new(big.Int)
			}
			return i
		}
	case int64:
		return big.NewInt(val)
	case uint64:
		return new(big.Int).SetUint64(val)
	case float64:
		split := strings.Split(fmt.Sprintf("%f", val), ".")
		result, _ := new(big.Int).SetString(split[0], 10)

		if len(split) > 1 {
			decimal := new(big.Int)
			decimal.SetString(split[1], 10)
			exp := big.NewInt(int64(math.Pow10(len(split[1]))))
			decimal.Mul(decimal, exp)
			result.Add(result, decimal)
		}

		return result
	case *big.Int:
		return val
	default:
		return big.NewInt(0)
	}
}


type Formatter struct {}

func (f *Formatter) ToMethodFormat(method string) string {
	return "0x" + strings.TrimLeft(ToSha3(method)[:8], "0")
}

func (f *Formatter) ToAddressFormat(address string) string {
	//if IsAddress(address) {
		address = strings.ToLower(address)

		if IsZeroPrefixed(address) {
			address = strings.TrimPrefix(address, "0x")
		}
	//}
	return strings.Repeat("0", 64-len(address)) + address
}

func (f *Formatter)ToIntegerFormat(value interface{}, digit int) string {
	bn := ToBn(value)
	bnHex := bn.Text(16)

	padded := "0"
	if string(bnHex[0]) != "f" {
		padded = "0"
	}

	return strings.Repeat(padded, digit-len(bnHex)) + bnHex
}


func ToSha3(input string) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(input))
	return string(hex.EncodeToString(hash.Sum(nil)))
}

