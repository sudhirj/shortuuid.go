package shortuuid

import (
	"math"
	"strings"
)

const BASE62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(num int64) string {
	return EncodeCustom(num, BASE62)
}

func EncodeCustom(num int64, alphabet string) string {
	if num == 0 {
		return string(alphabet[0])
	}
	var parts []byte
	for num > 0 {
		base := int64(len(alphabet))
		remainder := math.Mod(float64(num), float64(base))
		num /= base
		parts = append([]byte{alphabet[int(remainder)]}, parts...)
	}
	return string(parts)
}

func Decode(code string) int64 {
	return DecodeCustom(code, BASE62)
}

func DecodeCustom(code string, alphabet string) int64 {
	base := int64(len(alphabet))
	var num int64
	for _, i := range []byte(code) {
		i := strings.IndexByte(alphabet, i)
		if i < 0 {
			return 0
		}
		num = base*num + int64(i)
	}
	return num
}
