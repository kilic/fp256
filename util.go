package fp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
)

func fe(field *Field, s string) *FieldElement {
	if s[:2] == "0x" {
		s = s[2:]
	}
	h, _ := hex.DecodeString(s)
	if field == nil {
		return new(FieldElement).Unmarshal(h)
	}
	return field.NewElement(h)
}

var (
	big1   = new(big.Int).SetUint64(1)
	big2   = new(big.Int).SetUint64(2)
	big3   = new(big.Int).SetUint64(3)
	big64  = new(big.Int).SetUint64(64)
	big256 = new(big.Int).SetUint64(256)
)

func bn() *big.Int {
	return new(big.Int)
}

func bigFromStr10(s string) *big.Int {
	n, _ := new(big.Int).SetString(s, 10)
	return n
}

func bigFromStr16(s string) *big.Int {
	if s[:2] == "0x" {
		s = s[2:]
	}
	n, _ := new(big.Int).SetString(s, 16)
	return n
}

func bigFromInt64(i int64) *big.Int {
	return new(big.Int).SetInt64(i)
}

func toBytes(s string) []byte {
	h, _ := hex.DecodeString(s)
	return h
}

func toUint(s string) [4]uint64 {
	var i int64
	var bigTwo = bigFromInt64(2)
	value := bigFromStr16(s[2:])
	b := bn().Exp(bigTwo, bigFromInt64(64), nil)
	digits := [4]uint64{0, 0, 0, 0}
	for i < 4 {
		digits[i] = bn().Mod(value, b).Uint64()
		value.Div(value, b)
		i++
	}
	return digits
}

type hexstr string

func (s hexstr) uints4() [4]uint64 {
	if s[:2] == "0x" {
		s = s[2:]
	}
	r := new([4]uint64)
	for i := 0; i < 4; i++ {
		bs, _ := hex.DecodeString(string(s[(3-i)*16 : (4-i)*16]))
		r[i] = binary.BigEndian.Uint64(bs)
	}
	return *r
}

func (s hexstr) uints8() [8]uint64 {
	if s[:2] == "0x" {
		s = s[2:]
	}
	r := new([8]uint64)
	for i := 0; i < 8; i++ {
		bs, _ := hex.DecodeString(string(s[(7-i)*16 : (8-i)*16]))
		r[i] = binary.BigEndian.Uint64(bs)
	}
	return *r
}

func hexString(e []uint64) (r string) {
	for i := len(e); i > 0; i-- {
		r = r + fmt.Sprintf("%16.16x", e[i-1])
	}
	return
}
