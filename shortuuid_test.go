package shortuuid

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		num  int64
		code string
	}{
		{num: 0, code: "0"},
		{num: 1234, code: "Ju"},
		{num: 7654321, code: "W7En"},
		{num: 7654321847908756, code: "Z3Wcy65oK"},
	}
	for _, test := range tests {
		t.Run(test.code, func(t *testing.T) {
			if Encode(test.num) != test.code {
				t.Error("Encoding failed : " + Encode(test.num))
			}
			if Decode(test.code) != test.num {
				t.Error("Decoding failed : " + strconv.Itoa(int(Decode(test.code))))
			}
		})
	}
}

func TestLexicographicSort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Int63()
		b := rand.Int63()
		encA := fmt.Sprintf("%020s", Encode(a))
		encB := fmt.Sprintf("%020s", Encode(b))
		if (a < b) && (encA > encB) {
			t.Error("Lexicographic codes failed ", a, b, encA, encB)
		}
	}
}
