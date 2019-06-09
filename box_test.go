package fp

import (
	"crypto/rand"
	"testing"
)

var nBox = 100000

func TestBoxFieldELementByteInOut(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b = new(FieldElement), new(FieldElement)
	bytes := make([]byte, 32)
	for i := 0; i < nBox; i++ {
		field.RandElement(a, rand.Reader)
		a.Marshal(bytes)
		field.Mul(a, a, field.r2)
		b = field.NewElement(bytes)
		if !b.eq(a) {
			t.Errorf("bad byte conversion in:%s, out:%s",
				a.String(), b.String())
			return
		}
	}
}

func TestBoxAdditiveAssoc(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c, u, v FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.RandElement(&b, rand.Reader)
		field.RandElement(&c, rand.Reader)
		field.Add(&u, &a, &b)
		field.Add(&u, &u, &c)
		field.Add(&v, &b, &c)
		field.Add(&v, &v, &a)
		if !u.eq(&v) {
			t.Errorf("additive associativity does not hold a:%s, b:%s, c:%s, u:%s, v:%s",
				a.String(), b.String(), c.String(), u.String(), v.String())
			return
		}
	}
}

func TestBoxSubractiveAssoc(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c, u, v FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.RandElement(&b, rand.Reader)
		field.RandElement(&c, rand.Reader)
		field.Sub(&u, &a, &c)
		field.Sub(&u, &u, &b)
		field.Sub(&v, &a, &b)
		field.Sub(&v, &v, &c)
		if !u.eq(&v) {
			t.Errorf("subtractive associativity does not hold a:%s, b:%s, c:%s, u:%s, v:%s",
				a.String(), b.String(), c.String(), u.String(), v.String())
			return
		}
	}
}

func TestBoxMultiplicativeAssoc(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c, u, v FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.RandElement(&b, rand.Reader)
		field.RandElement(&c, rand.Reader)
		field.Mul(&u, &a, &b)
		field.Mul(&u, &u, &c)
		field.Mul(&v, &b, &c)
		field.Mul(&v, &v, &a)
		if !u.eq(&v) {
			t.Errorf("multiplicative associativity does not hold a:%s, b:%s, c:%s, u:%s, v:%s",
				a.String(), b.String(), c.String(), u.String(), v.String())
			return
		}
	}
}

func TestBoxAdditiveCommutativity(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, u, v FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.RandElement(&b, rand.Reader)
		field.Add(&u, &a, &b)
		field.Add(&v, &b, &a)
		if !u.eq(&v) {
			t.Errorf("additive commutativity  does not hold a:%s, b:%s, u:%s",
				a.String(), b.String(), u.String())
			return
		}
	}
}

func TestBoxMultiplicativeCommutativity(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, u, v FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.RandElement(&b, rand.Reader)
		field.Mul(&u, &a, &b)
		field.Mul(&v, &b, &a)
		if !u.eq(&v) {
			t.Errorf("multiplicative commutativity does not hold a:%s, b:%s, u:%s",
				a.String(), b.String(), u.String())
			return
		}
	}
}

func TestBoxSquare(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.Square(&b, &a)
		field.Mul(&c, &a, &a)
		if !c.eq(&b) {
			t.Errorf("bad squaring, have: %s, want: %s", c.String(), b.String())
			return
		}
	}
}

func TestBoxNegation(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, u, v FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.RandElement(&b, rand.Reader)
		field.Sub(&u, &a, &b)
		field.Neg(&a, &a)
		field.Neg(&b, &b)
		field.Sub(&v, &b, &a)
		if !u.eq(&v) {
			t.Errorf("subtraction check does not hold a:%s, b:%s, u:%s",
				a.String(), b.String(), u.String())
			return
		}
	}
}

func TestBoxNegation2(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c FieldElement
	var zero = &FieldElement{0, 0, 0, 0}
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.Neg(&b, &a)
		field.Add(&c, &a, &b)
		if !zero.eq(&c) {
			t.Errorf("bad negation a:%s, b:%s",
				a.String(), b.String())
			return
		}
	}
}

func TestBoxDoubling(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c, monttwo FieldElement
	field.Mul(&monttwo, &FieldElement{2, 0, 0, 0}, field.r2)
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.Double(&b, &a)
		field.Mul(&c, &a, &monttwo)
		if !b.eq(&c) {
			t.Errorf("bad doubling c:%s, b:%s",
				c.String(), b.String())
			return
		}
	}
}

func TestBoxAdditiveIdentity(t *testing.T) {

	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, c FieldElement
	identity := &FieldElement{0, 0, 0, 0}
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.Add(&c, &a, identity)
		if !c.eq(&a) {
			t.Errorf("additive identity does not hold, have: %s, want: %s", c.String(), a.String())
			return
		}
		field.Add(&a, &c, identity)
		if !c.eq(&a) {
			t.Errorf("additive identity does not hold, have: %s, want: %s", c.String(), a.String())
			return
		}
	}
}

func TestBoxMultiplicativeIdentity(t *testing.T) {

	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, c FieldElement
	identity := field.r1
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.Mul(&c, &a, identity)
		if !c.eq(&a) {
			t.Errorf("multiplicative identity does not hold, have: %s, want: %s", c.String(), a.String())
			return
		}
		field.Mul(&a, &c, identity)
		if !c.eq(&a) {
			t.Errorf("multiplicative identity does not hold, have: %s, want: %s", c.String(), a.String())
			return
		}
	}
}

func TestBoxInverseDown(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c FieldElement
	e := &FieldElement{1, 0, 0, 0}
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.InvMontDown(&b, &a)
		field.Mul(&c, &b, &a)
		if !c.eq(e) {
			t.Errorf("bad montgomery downgrade inversion have: %s, want: %s", c.String(), e.String())
			return
		}
	}
}

func TestBoxInverse(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	var a, b, c FieldElement
	e := field.r1
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.InvMontUp(&b, &a)
		field.Mul(&c, &b, &a)
		if !c.eq(e) {
			t.Errorf("bad montgomery upgrade inversion, have: %s, want: %s", c.String(), e.String())
			return
		}
	}
}

func TestBoxExp(t *testing.T) {
	p := bigFromStr16(testmodulus)
	field := NewField(p)
	p2 := new(FieldElement)
	sub(p2, new(FieldElement).Unmarshal(p.Bytes()), &FieldElement{2, 0, 0, 0})
	var ai1, ai2, a FieldElement
	for i := 0; i < nBox; i++ {
		field.RandElement(&a, rand.Reader)
		field.Exp(&ai1, &a, p2)
		field.InvMontUp(&ai2, &a)
		if !ai1.eq(&ai2) {
			t.Errorf("exponentiation fails , have %s, want %s", ai2.String(), ai1.String())
			return
		}
	}
}
