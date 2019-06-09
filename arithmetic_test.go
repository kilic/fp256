package fp

import (
	"math/big"
	"testing"
)

var testmodulus = "0x73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001"

// var testmodulus = "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"

func TestAddition(t *testing.T) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	e := fe(nil, "0x46bd0357810d2d61ef4e4a7fc390f52dc65c76170001a3129aac34358b358d34")
	c := &FieldElement{}
	add(c, a, b)
	if !e.eq(c) {
		t.Errorf("field element addition fails, have %s, want %s", c.String(), e.String())
	}
	a = fe(nil, "0x00")
	e = b
	add(c, a, b)
	if !e.eq(c) {
		t.Errorf("field element zero addition fails, have %s, want %s", c.String(), e.String())
	}
}

func TestDoubling(t *testing.T) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	e := fe(nil, "0x6167ae022bb7d80c561ab14c7fb2b14fcf657f210001a20100233334ff540353")
	c := &FieldElement{}
	double(c, a)
	if !e.eq(c) {
		t.Errorf("field element doubling fails, have %s, want %s", c.String(), e.String())
	}
	a = fe(nil, "0x00")
	e = a
	double(c, a)
	if !e.eq(c) {
		t.Errorf("field element zero addition fails, have %s, want %s", c.String(), e.String())
	}
}

func TestNegate(t *testing.T) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	e := fe(nil, "0x0942fca87ef2d29dee8f935dc4f7935ac22c1270fffe5cfeffee66650055fe57")
	c := &FieldElement{}
	neg(c, a)
	if !e.eq(c) {
		t.Errorf("field element negation fails, have %s, want %s", c.String(), e.String())
	}
}

func TestSubtraction(t *testing.T) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	e := fe(nil, "0x5942fca87ef2d29dcc6d713b4d801be34ab49af8fffe5d109a8900ff8be189e2")
	c := &FieldElement{}
	sub(c, b, a)
	if !e.eq(c) {
		t.Errorf("field element subtraction fails, have %s, want %s", c.String(), e.String())
	}
	e = fe(nil, "0x1aaaaaaaaaaaaaaa66cc66ccbc21bc2209090909fffffeee6576feff741e761f")
	sub(c, a, b)
	if !e.eq(c) {
		t.Errorf("field element subtraction fails, have %s, want %s", c.String(), e.String())
	}
	e = fe(nil, "0x00")
	sub(c, a, a)
	if !e.eq(c) {
		t.Errorf("field element subtraction fails, have %s, want %s", c.String(), e.String())
	}
	a = fe(nil, "0x00")
	sub(c, a, b)
	neg(e, b)
	if !e.eq(c) {
		t.Errorf("field element subtraction fails, have %s, want %s", c.String(), e.String())
	}
	sub(c, b, a)
	e = b
	if !e.eq(c) {
		t.Errorf("field element subtraction fails, have %s, want %s", c.String(), e.String())
	}
}

func TestMontgomeryReduction1(t *testing.T) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	T := [8]uint64{
		0x22bbccdd55558888,
		0xaabbccdd55558888,
		0xaabbccdd55558888,
		0x11bbccdd55558888,
		0x22bbccdd55558888,
		0xaabbccdd55558888,
		0xaabbccdd55558888,
		0x22bbccdd55558888,
	}
	e := fe(nil, "0x0ac1b4094057dae42dab79d6693ee71d832ffa2bb7648e3884a7d38f035dceed")
	r := new(FieldElement)
	mont(r, &T)
	if !r.eq(e) {
		t.Errorf("montgomerry reduction fails, have %s, want %s", r.String(), e.String())
	}
}

func TestMontgomeryReduction2(t *testing.T) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	T := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}
	e := &FieldElement{0, 0, 0, 0}
	r := new(FieldElement)
	mont(r, &T)
	if !r.eq(e) {
		t.Errorf("montgomerry reduction fails, have %s, want %s", r.String(), e.String())
	}
}

func TestMontgomeryReduction3(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	T := [8]uint64{1, 0, 0, 0, 0, 0, 0, 0}
	e := field.rN1
	r := new(FieldElement)
	mont(r, &T)
	if !r.eq(e) {
		t.Errorf("montgomerry reduction fails, have %s, want %s", r.String(), e.String())
	}
}

func TestMultiplication(t *testing.T) {
	var e hexstr = "0x2155555555555555273c51e6d9036e58e950083f02ccfc1de185217fc929459fc0467f564e5ba377b716c90e4c730d1a2078bc4c1a3544fdf1ed68e95584354e"
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	var c [8]uint64
	mul(&c, a, b)
	if c != e.uints8() {
		t.Errorf("multi precision multiplication fails for first limb\n, have %s, want %s", hexString(c[:]), e)
	}
}

func TestMontgomeryMultiplication1(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x1aaaaaaaaaaaaaa8ccccccccccccccc8777777777777777dffffffffddddddd0")
	b := fe(field, "0x6cccccccccccccc911111111555555559393939393939393ffffffffeeeeeeef")
	e := fe(field, "0x678f0b264343979944fb3663d336c345b4347ed20629e7de98aa44cc4aa681bb")
	montmul(a, a, b)
	if !a.eq(e) {
		t.Errorf("mont multiplication fails, have %s, want %s", a.String(), e.String())
	}
}

func TestMontgomeryMultiplication2(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(nil, "0x00")
	e := fe(nil, "0x00")
	montmul(a, a, field.r1)
	if !a.eq(e) {
		t.Errorf("modular reduction fails, have %s, want %s", a.String(), e.String())
	}
}

func TestMontgomeryMultiplication3(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(nil, "0x01")
	e := fe(nil, "0x01")
	montmul(a, a, field.r1)
	if !a.eq(e) {
		t.Errorf("modular reduction fails, have %s, want %s", a.String(), e.String())
	}
}

func TestMontgomeryMultiplication4(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(nil, "0x01")
	b := fe(nil, "0x01")
	e := field.rN1
	montmul(a, a, b)
	if !a.eq(e) {
		t.Errorf("modular reduction fails, have %s, want %s", a.String(), e.String())
	}
}

func TestMontgomeryMultiplication5(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	field := NewField(p)
	a := fe(field, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x01")
	e := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	montmul(a, a, b)
	if !a.eq(e) {
		t.Errorf("field element addition fails, have %s, want %s", a.String(), e.String())
	}
}

func TestMontgomeryMultiplication6(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x01")
	b := fe(field, "0x6cccccccccccccc911111111555555559393939393939393ffffffffeeeeeeef")
	e := b
	c := new(FieldElement)
	montmul(c, a, b)
	if !c.eq(e) {
		t.Errorf("mont multiplication fails, have %s, want %s", c.String(), e.String())
	}
}

func TestSquaring(t *testing.T) {
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	var e hexstr = "0x2c71c71c71c71c71721c1cc6c771721c85812e25b2d02ff41c33f3811d3fa1bce41719ece0c67ff282b914d2dc9b6186872d471e6e9a10d1071b516ae1cac4e4"
	b := [8]uint64{}
	square(&b, a)
	if e.uints8() != b {
		t.Errorf("multi precision multiplication fails for carry limb\n, have %s, want %s", hexString(b[:]), e)
	}
}

func TestMontgomerySquaring1(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x1aaaaaaaaaaaaaa8ccccccccccccccc8777777777777777dffffffffddddddd0")
	e := fe(field, "0x39b25f1073641793dcb0007efc52a272b06514bd57562d45292d141976190f71")
	c := new(FieldElement)
	montsquare(c, a)
	if !c.eq(e) {
		t.Errorf("mont multiplication fails, have %s, want %s", c.String(), e.String())
	}
}

func TestMontgomerySquaring2(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := fe(field, "0x01")
	e := field.r1
	c := new(FieldElement)
	montsquare(c, a)
	if !c.eq(e) {
		t.Errorf("mont squaring one fails, have %s, want %s", c.String(), e.String())
	}
}

func TestMontgomerySquaring3(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	a := &FieldElement{0, 0, 0, 0}
	e := fe(field, "0x00")
	c := new(FieldElement)
	montsquare(c, a)
	if !c.eq(e) {
		t.Errorf("mont squaring one fails, have %s, want %s", c.String(), e.String())
	}
}

func BenchmarkAddition(t *testing.B) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		add(c, a, b)
	}
}

func BenchmarkDoubling(t *testing.B) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	c := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		double(c, a)
	}
}

func BenchmarkSubtraction(t *testing.B) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sub(c, a, b)
	}
}

func BenchmarkMultiplication(t *testing.B) {
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	w := new([8]uint64)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mul(w, a, b)
	}
}

func BenchmarkSquaring(t *testing.B) {
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	w := new([8]uint64)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		square(w, a)
	}
}

func BenchmarkMontgomeryReduction(t *testing.B) {
	p := bigFromStr16(testmodulus)
	_ = NewField(p)
	T := [8]uint64{
		0x22bbccdd55558888,
		0xaabbccdd55558888,
		0xaabbccdd55558888,
		0x11bbccdd55558888,
		0x22bbccdd55558888,
		0xaabbccdd55558888,
		0xaabbccdd55558888,
		0x22bbccdd55558888,
	}
	var b = new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mont(b, &T)
	}
	_ = b
}

func BenchmarkMontgomerryMultiplication(t *testing.B) {
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := fe(nil, "0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	w := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		montmul(w, a, b)
	}
}

func BenchmarkMontgomerySquaring(t *testing.B) {
	a := fe(nil, "0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	w := new(FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		montsquare(w, a)
	}
}

func BenchmarkBigIntAddition(t *testing.B) {
	a := bigFromStr16("0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := bigFromStr16("0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(big.Int)
	for i := 0; i < t.N; i++ {
		c.Add(a, b)
	}
}

func BenchmarkBigIntMultiplication(t *testing.B) {
	a := bigFromStr16("0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := bigFromStr16("0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(big.Int)
	for i := 0; i < t.N; i++ {
		c.Mul(a, b)
	}
}

func BenchmarkBigIntMod(t *testing.B) {
	p := bigFromStr16(testmodulus)
	a := bigFromStr16("0x6aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b := bigFromStr16("0x4fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(big.Int).Mul(a, b)
	for i := 0; i < t.N; i++ {
		c.Mod(c, p)
	}
}
