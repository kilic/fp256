package fp

import (
	"encoding/binary"
	"math/big"
	"reflect"
	"testing"
)

func TestFieldElementFromOneByte(t *testing.T) {
	bytes := []byte{
		1,
	}
	a := new(FieldElement).Unmarshal(bytes)
	e := &FieldElement{1, 0, 0, 0}
	if *a != *e {
		t.Errorf("cannot unmarshal bytes have: %x, want %x", a, e)
	}
}

func TestFieldElementFromOneBytes1(t *testing.T) {
	bytes := []byte{
		255, 255,
	}
	a := new(FieldElement).Unmarshal(bytes)
	e := &FieldElement{0xffff, 0, 0, 0}
	if *a != *e {
		t.Errorf("cannot unmarshal bytes have: %s, want %s", a.String(), e.String())
	}
}

func TestFieldElementFromOneBytes2(t *testing.T) {
	bytes := []byte{
		255, 255, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	a := new(FieldElement).Unmarshal(bytes)
	e := &FieldElement{0, 0xffff, 0, 0}
	if *a != *e {
		t.Errorf("cannot unmarshal bytes have: %s, want %s", a.String(), e.String())
	}
}

func TestFieldElementFromOneBytes3(t *testing.T) {
	x0, x64, x128, x192 := uint64(0x11bbbbbbbbbbbb22), uint64(0x33bbbbbbbbbbbb44),
		uint64(0x55bbbbbbbbbbbb66), uint64(0x77bbbbbbbbbbbb88)
	bytes := make([]byte, 32)
	binary.BigEndian.PutUint64(bytes[:], x0)
	binary.BigEndian.PutUint64(bytes[8:], x64)
	binary.BigEndian.PutUint64(bytes[16:], x128)
	binary.BigEndian.PutUint64(bytes[24:], x192)
	a := new(FieldElement).Unmarshal(bytes)
	e := &FieldElement{x192, x128, x64, x0}
	if *a != *e {
		t.Errorf("cannot unmarshal bytes have: %s, want %s", a.String(), e.String())
	}
}

func TestFieldElementToBytes1(t *testing.T) {
	bytes := []byte{
		100, 200,
	}
	a := new(FieldElement).Unmarshal(bytes)
	out := make([]byte, 32)
	a.Marshal(out)
	zeros := make([]byte, 30)
	if !reflect.DeepEqual(zeros[:], out[:30]) || !reflect.DeepEqual(bytes[:], out[30:]) {
		t.Errorf("cannot marshal field element")
	}
}

func TestFieldElementToBytes2(t *testing.T) {
	in := []byte{
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x10, 0x20,
		0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80, 0x90,
		0xfa, 0xfa, 0xfa, 0xfa, 0x24, 0x23, 0x22, 0x21,
	}
	a := new(FieldElement).Unmarshal(in)
	out := make([]byte, 32)
	a.Marshal(out)
	if !reflect.DeepEqual(in[:], out[:]) {
		t.Errorf("cannot marshal field element, have: %x, want: %x", out, in)
	}
}

func TestFieldElementToBytes3(t *testing.T) {
	in := []byte{
		0x11, 0x22, 0x33, 0x55, 0x66, 0x77, 0x88,
		0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x10, 0x20,
		0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80, 0x90,
		0xfa, 0xfa, 0xfa, 0xfa, 0x24, 0x23, 0x22, 0x21,
	}
	a := new(FieldElement).Unmarshal(in)
	out := make([]byte, 32)
	a.Marshal(out)
	if !reflect.DeepEqual(in[:], out[1:]) || out[0] != 0 {
		t.Errorf("cannot marshal field element, have: %x, want: %x", out, in)
	}
}

func TestFieldExponentiation(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x66ffeeeeddddccccffffeeeeddddcccc99aa99aa88bb88bb1919191928282828")
	x := fe(nil, "0x27c28f49cabcf02ec28a6a44d07436e062d004894bffeeefa73ab2abc10f487f")
	field.Exp(a, a, x)
	montmul(a, a, &FieldElement{1, 0, 0, 0})
	e := fe(nil, "0x028c8f7ec1cebad4afe67cbfb965e72cf8f26f4219b47dd44f5990359c486c23")
	if !e.eq(a) {
		t.Errorf("exponentiation fails , have %s, want %s", a.String(), e.String())
	}
}

func TestFieldExponentiation2(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x66ffeeeeddddccccffffeeeeddddcccc99aa99aa88bb88bb1919191928282828")
	ai1 := new(FieldElement)
	ai2 := new(FieldElement)
	p2 := new(FieldElement)
	sub(p2, new(FieldElement).Unmarshal(p.Bytes()), &FieldElement{2, 0, 0, 0})
	field.Exp(ai1, a, p2)
	field.InvMontUp(ai2, a)
	if !ai1.eq(ai2) {
		t.Errorf("exponentiation fails , have %s, want %s", ai2.String(), ai1.String())
	}
}

func TestFieldInverseEuclid(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(nil, "0x66ffeeeeddddccccffffeeeeddddcccc99aa99aa88bb88bb1919191928282828")
	e := fe(nil, "0x27c28f49cabcf02ec28a6a44d07436e062d004894bffeeefa73ab2abc10f487f")
	inv := new(FieldElement)
	field.InvEEA(inv, a)
	if !inv.eq(e) {
		t.Errorf("inversion fails (euclid), have %s, want %s", inv.String(), e.String())
	}
}

func TestFieldInverseMontgomeryDown(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x66ffeeeeddddccccffffeeeeddddcccc99aa99aa88bb88bb1919191928282828")
	e := fe(nil, "0x27c28f49cabcf02ec28a6a44d07436e062d004894bffeeefa73ab2abc10f487f")
	inv := new(FieldElement)
	field.InvMontDown(inv, a)
	if !inv.eq(e) {
		t.Errorf("inversion fails (montgomery down), have %s, want %s", inv.String(), e.String())
	}
	a = fe(nil, "0x66ffeeeeddddccccffffeeeeddddcccc99aa99aa88bb88bb1919191928282828")
	e = fe(field, "0x27c28f49cabcf02ec28a6a44d07436e062d004894bffeeefa73ab2abc10f487f")
	inv = new(FieldElement)
	field.InvMontDown(inv, a)
	if !inv.eq(e) {
		t.Errorf("inversion fails (montgomery down), have %s, want %s", inv.String(), e.String())
	}
}

func TestFieldInverseMontgomeryUp(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x66ffeeeeddddccccffffeeeeddddcccc99aa99aa88bb88bb1919191928282828")
	e := fe(field, "0x27c28f49cabcf02ec28a6a44d07436e062d004894bffeeefa73ab2abc10f487f")
	inv := new(FieldElement)
	field.InvMontUp(inv, a)
	if !inv.eq(e) {
		t.Errorf("inversion fails (montgomery up), have %s, want %s", inv.String(), e.String())
	}
}

func BenchmarkExponentiation(t *testing.B) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x66ffeeeeddddccccffffeeeeddddcccc99aa99aa88bb88bb1919191928282828")
	x := fe(nil, "0x27c28f49cabcf02ec28a6a44d07436e062d004894bffeeefa73ab2abc10f487f")
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		field.Exp(a, a, x)
	}
}
func BenchmarkFieldInverseEuclid(t *testing.B) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(nil, "0x1aaaaaaaaaaaaaa8ccccccccccccccc8777777777777777dffffffffddddddd0")
	inv := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		field.InvEEA(inv, a)
	}
}

func BenchmarkFieldInverseMontgomeryDown(t *testing.B) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(nil, "0x1aaaaaaaaaaaaaa8ccccccccccccccc8777777777777777dffffffffddddddd0")
	inv := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		field.InvMontDown(inv, a)
	}
}

func BenchmarkFieldInverseMontgomeryUp(t *testing.B) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(nil, "0x1aaaaaaaaaaaaaa8ccccccccccccccc8777777777777777dffffffffddddddd0")
	inv := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		field.InvMontUp(inv, a)
	}
}

func BenchmarkBigIntInversion(t *testing.B) {
	p := bigFromStr16(testmodulus)
	a := bigFromStr16("0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	c := new(big.Int)
	for i := 0; i < t.N; i++ {
		c.ModInverse(a, p)
	}
}

func BenchmarkBigIntExponentiation(t *testing.B) {
	p := bigFromStr16(testmodulus)
	a := bigFromStr16("0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := bigFromStr16("0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(big.Int).Mul(a, b)
	for i := 0; i < t.N; i++ {
		c.Exp(a, b, p)
	}
}
