package util

import "math/big"

func EncodeBase62(i int64) string {
	bi := new(big.Int)
	bi.SetInt64(i)
	return bi.Text(62)
}

func DecodeBase62(s string) int64 {
	bi := new(big.Int)
	bi.SetString(s, 62)
	return bi.Int64()
}
