package fp

import (
	"encoding/hex"
	"io"
	"math/big"
	"math/bits"
)

// inp = (-p^{-1} mod 2^b) where b = 64
var inp uint64
var modulus FieldElement

type Field struct {
	// p2  = p-2
	// rN1 = r^1 modp
	// r1  = r modp
	// r2  = r^2 modp
	// r3  = r^3 modp
	p2  *FieldElement
	rN1 *FieldElement
	r1  *FieldElement
	r2  *FieldElement
	r3  *FieldElement
}

// Given prime number as big.Int,
// field constants are precomputed
func NewField(pBig *big.Int) *Field {
	modulus = *new(FieldElement).Unmarshal(pBig.Bytes())
	inp = bn().ModInverse(bn().Neg(pBig), bn().Exp(big2, big64, nil)).Uint64()
	r1Big := bn().Exp(big2, big256, nil)
	r1 := new(FieldElement).Unmarshal(bn().Mod(r1Big, pBig).Bytes())
	r2 := new(FieldElement).Unmarshal(bn().Exp(r1Big, big2, pBig).Bytes())
	r3 := new(FieldElement).Unmarshal(bn().Exp(r1Big, big3, pBig).Bytes())
	rN1 := new(FieldElement).Unmarshal(bn().ModInverse(r1Big, pBig).Bytes())
	p2 := new(FieldElement).Unmarshal(bn().Sub(pBig, big2).Bytes())
	return &Field{
		p2:  p2,
		r1:  r1,
		rN1: rN1,
		r2:  r2,
		r3:  r3,
	}
}

// Returns new element in Montgomery domain
func (f *Field) NewElement(in []byte) *FieldElement {
	fe := new(FieldElement).Unmarshal(in)
	f.Mul(fe, fe, f.r2)
	return fe
}

func (f *Field) NewElementFromUint(in uint64) *FieldElement {
	fe := &FieldElement{in, 0, 0, 0}
	if in == 0 {
		return fe
	}
	montmul(fe, fe, f.r2)
	return fe
}

func (f *Field) NewElementFromString(in string) (*FieldElement, error) {
	if in[:2] == "0x" {
		in = in[2:]
	}
	h, err := hex.DecodeString(in)
	if err != nil {
		return nil, err
	}
	fe := new(FieldElement).Unmarshal(h)
	montmul(fe, fe, f.r2)
	return fe, nil
}

// Adapted from https://github.com/golang/go/blob/master/src/crypto/rand/util.go
func (f *Field) RandElement(fe *FieldElement, r io.Reader) error {
	// assuming p > 2^192
	bitLen := bits.Len64(modulus[3]) + 64 + 64 + 64
	// k is the maximum byte length needed to encode a value < max.
	k := (bitLen + 7) / 8
	// b is the number of bits in the most significant byte of max-1.
	b := uint(bitLen % 8)
	if b == 0 {
		b = 8
	}
	bytes := make([]byte, k)
	for {
		_, err := io.ReadFull(r, bytes)
		if err != nil {
			return err
		}
		// Clear bits in the first byte to increase the probability
		// that the candidate is < max.
		bytes[0] &= uint8(int(1<<b) - 1)
		fe.Unmarshal(bytes)

		if fe.Cmp(&modulus) < 0 {
			break
		}
	}
	return nil
}

func (f *Field) Mont(c, a *FieldElement) {
	montmul(c, a, f.r2)
}

func (f *Field) Demont(c, a *FieldElement) {
	//mont(c, &[8]uint64{a[0], a[1], a[2], a[3], 0, 0, 0, 0})
	montmul(c, a, &FieldElement{1, 0, 0, 0})
}

// c = (a + b) modp
func (f *Field) Add(c, a, b *FieldElement) {
	add(c, a, b)
}

// c = (a + a) modp
func (f *Field) Double(c, a *FieldElement) {
	double(c, a)
}

// c = (a - b) modp
func (f *Field) Sub(c, a, b *FieldElement) {
	sub(c, a, b)
}

func (f *Field) Neg(c, a *FieldElement) {
	neg(c, a)
}

// Sets c as a^2(R^-1) modp
func (f *Field) Square(c, a *FieldElement) {
	montsquare(c, a)
}

// Sets c as ab(R^-1) modp
func (f *Field) Mul(c, a, b *FieldElement) {
	montmul(c, a, b)
}

// Sets c as (a^e) modp
func (f *Field) Exp(c, a, x *FieldElement) {
	z := new(FieldElement).Set(f.r1) // A
	var i uint64
	for i = 255; i != 0xffffffffffffffff; i-- {
		montmul(z, z, z)
		if x.bit(i) {
			montmul(z, z, a)
		}
	}
	c.Set(z)
}

// Guide to Elliptic Curve Cryptography Algorithm
// Hankerson, Menezes, Vanstone
// Algoritm 2.22 Binary algorithm for inversion in Fp
// Input: a
// Output: a^-1
func (f *Field) InvEEA(inv, fe *FieldElement) {
	u := new(FieldElement).Set(fe)
	v := new(FieldElement).Set(&modulus)
	p := new(FieldElement).Set(&modulus)
	x1 := &FieldElement{1, 0, 0, 0}
	x2 := &FieldElement{0, 0, 0, 0}
	var e uint64

	for !u.IsOne() && !v.IsOne() {
		for u.IsEven() {
			u.rightShift(0)
			if x1.IsEven() {
				x1.rightShift(0)
			} else {
				e = addn(x1, p)
				x1.rightShift(e)
			}
		}
		for v.IsEven() {
			v.rightShift(0)
			if x2.IsEven() {
				x2.rightShift(0)
			} else {
				addn(x2, p)
				x2.rightShift(e)
			}
		}
		if u.Cmp(v) == -1 {
			subn(v, u)
			sub(x2, x2, x1)
		} else {
			subn(u, v)
			sub(x1, x1, x2)
		}
	}
	if u.IsOne() {
		inv.Set(x1)
		return
	}
	inv.Set(x2)
}

// Two phase Montgomery Modular Inverse
// The Montgomery Modular Inverse - Revisited
// Savas, Koc
// &
// Guide to Elliptic Curve Cryptography Algorithm
// Hankerson, Menezes, Vanstone
// Algoritm 2.23 Partial Montgomery inversion in Fp
//
// Input : a
// Output : (a^-1)R
// or
// Input : aR
// Output : (a^-1)
func (f *Field) InvMontDown(inv, fe *FieldElement) {

	u := new(FieldElement).Set(fe)
	v := new(FieldElement).Set(&modulus)
	x1 := &FieldElement{1, 0, 0, 0}
	x2 := &FieldElement{0, 0, 0, 0}
	var k int
	// Phase 1
	for !v.IsZero() {
		if v.IsEven() {
			v.rightShift(0)
			x1.leftShift()
		} else if u.IsEven() {
			u.rightShift(0)
			x2.leftShift()
		} else if v.Cmp(u) == -1 {
			subn(u, v)
			u.rightShift(0)
			addn(x1, x2)
			x2.leftShift()
		} else {
			subn(v, u)
			v.rightShift(0)
			addn(x2, x1)
			x1.leftShift()
		}
		k = k + 1
	}
	// Phase2
	k = k - 256
	var e uint64
	for i := 0; i < k; i++ {
		if x1.IsEven() {
			x1.rightShift(0)
		} else {
			e = addn(x1, &modulus)
			x1.rightShift(e)
		}
	}
	inv.Set(x1)
}

// Inverse value stays in Montgomery space
// Two phase Montgomery Modular Inverse
// The Montgomery Modular Inverse - Revisited
// Savas, Koc
// &
// Guide to Elliptic Curve Cryptography Algorithm
// Hankerson, Menezes, Vanstone
// Algoritm 2.23 Partial Montgomery inversion in Fp
// Input : aR
// Output : (a^-1)R
func (f *Field) InvMontUp(inv, fe *FieldElement) {

	u := new(FieldElement).Set(fe)
	v := new(FieldElement).Set(&modulus)
	x1 := &FieldElement{1, 0, 0, 0}
	x2 := &FieldElement{0, 0, 0, 0}
	var k int

	// Phase 1
	for !v.IsZero() {
		if v.IsEven() {
			v.rightShift(0)
			x1.leftShift()
		} else if u.IsEven() {
			u.rightShift(0)
			x2.leftShift()
		} else if v.Cmp(u) == -1 {
			subn(u, v)
			u.rightShift(0)
			addn(x1, x2)
			x2.leftShift()
		} else {
			subn(v, u)
			v.rightShift(0)
			addn(x2, x1)
			x1.leftShift()
		}
		k = k + 1
	}
	// Phase2
	sub(x1, x1, &modulus)
	for i := k; i < 512; i++ {
		double(x1, x1)
	}
	inv.Set(x1)
}
