// +build !amd64, pure_go

package fp

func add(c, a, b *FieldElement) {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	p0 := modulus[0]
	p1 := modulus[1]
	p2 := modulus[2]
	p3 := modulus[3]
	var e, e2, ne uint64

	u0 := a0 + b0
	e = (a0&b0 | (a0|b0)&^u0) >> 63
	u1 := a1 + b1 + e
	e = (a1&b1 | (a1|b1)&^u1) >> 63
	u2 := a2 + b2 + e
	e = (a2&b2 | (a2|b2)&^u2) >> 63
	u3 := a3 + b3 + e
	e = (a3&b3 | (a3|b3)&^u3) >> 63

	v0 := u0 - p0
	e2 = (^u0&p0 | (^u0|p0)&v0) >> 63
	v1 := u1 - p1 - e2
	e2 = (^u1&p1 | (^u1|p1)&v1) >> 63
	v2 := u2 - p2 - e2
	e2 = (^u2&p2 | (^u2|p2)&v2) >> 63
	v3 := u3 - p3 - e2
	e2 = (^u3&p3 | (^u3|p3)&v3) >> 63

	e = e - e2
	ne = ^e

	c[0] = (u0 & e) | (v0 & ne)
	c[1] = (u1 & e) | (v1 & ne)
	c[2] = (u2 & e) | (v2 & ne)
	c[3] = (u3 & e) | (v3 & ne)
}

func double(c, a *FieldElement) {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	p0 := modulus[0]
	p1 := modulus[1]
	p2 := modulus[2]
	p3 := modulus[3]

	e := a3 >> 63
	u3 := a3<<1 | a2>>63
	u2 := a2<<1 | a1>>63
	u1 := a1<<1 | a0>>63
	u0 := a0 << 1

	v0 := u0 - p0
	e2 := (^u0&p0 | (^u0|p0)&v0) >> 63
	v1 := u1 - p1 - e2
	e2 = (^u1&p1 | (^u1|p1)&v1) >> 63
	v2 := u2 - p2 - e2
	e2 = (^u2&p2 | (^u2|p2)&v2) >> 63
	v3 := u3 - p3 - e2
	e2 = (^u3&p3 | (^u3|p3)&v3) >> 63

	e = e - e2
	ne := ^e

	c[0] = (u0 & e) | (v0 & ne)
	c[1] = (u1 & e) | (v1 & ne)
	c[2] = (u2 & e) | (v2 & ne)
	c[3] = (u3 & e) | (v3 & ne)
}

func sub(c, a, b *FieldElement) {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	p0 := modulus[0]
	p1 := modulus[1]
	p2 := modulus[2]
	p3 := modulus[3]

	var e, e2, ne uint64

	u0 := a0 - b0
	e = (^a0&b0 | (^a0|b0)&u0) >> 63
	u1 := a1 - b1 - e
	e = (^a1&b1 | (^a1|b1)&u1) >> 63
	u2 := a2 - b2 - e
	e = (^a2&b2 | (^a2|b2)&u2) >> 63
	u3 := a3 - b3 - e
	e = (^a3&b3 | (^a3|b3)&u3) >> 63

	v0 := u0 + p0
	e2 = (u0&p0 | (u0|p0)&^v0) >> 63
	v1 := u1 + p1 + e2
	e2 = (u1&p1 | (u1|p1)&^v1) >> 63
	v2 := u2 + p2 + e2
	e2 = (u2&p2 | (u2|p2)&^v2) >> 63
	v3 := u3 + p3 + e2

	e--
	ne = ^e
	c[0] = (u0 & e) | (v0 & ne)
	c[1] = (u1 & e) | (v1 & ne)
	c[2] = (u2 & e) | (v2 & ne)
	c[3] = (u3 & e) | (v3 & ne)
}

func neg(c, a *FieldElement) {
	sub(c, modulus, a)
}

func montmul(c, a, b *FieldElement) {
	w := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}
	mul256(&w, a, b)
	mont(c, &w)
}

func montsquare(c, a *FieldElement) {
	w := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}
	square256(&w, a)
	mont(c, &w)
}

// https://github.com/golang/go/blob/master/src/math/bits/bits.go
func mul64(a, b uint64) (hi, lo uint64) {
	const mask32 = 1<<32 - 1
	a0 := a & mask32
	a1 := a >> 32
	b0 := b & mask32
	b1 := b >> 32
	w0 := a0 * b0
	t := a1*b0 + w0>>32
	w1 := t & mask32
	w2 := t >> 32
	w1 += a0 * b1
	hi = a1*b1 + w2 + w1>>32
	lo = a * b
	return
}

func square64(a uint64) (hi, lo uint64) {
	return mul64(a, a)
}

// Handbook of Applied Cryptography
// Hankerson, Menezes, Vanstone
// 14.12 Algorithm Multiple-precision multiplication
func mul(w *[8]uint64, a, b *FieldElement) {

	var w0, w1, w2, w3, w4, w5, w6, w7 uint64
	var a0 = a[0]
	var a1 = a[1]
	var a2 = a[2]
	var a3 = a[3]
	var b0 = b[0]
	var b1 = b[1]
	var b2 = b[2]
	var b3 = b[3]
	var u, v, c, t uint64

	// i = 0, j = 0
	c, w0 = mul64(a0, b0)

	// i = 0, j = 1
	u, v = mul64(a1, b0)
	w1 = v + c
	c = u + (v&c|(v|c)&^w1)>>63

	// i = 0, j = 2
	u, v = mul64(a2, b0)
	w2 = v + c
	c = u + (v&c|(v|c)&^w2)>>63

	// i = 0, j = 3
	u, v = mul64(a3, b0)
	w3 = v + c
	w4 = u + (v&c|(v|c)&^w3)>>63

	//
	// i = 1, j = 0
	c, v = mul64(a0, b1)
	t = v + w1
	c += (v&w1 | (v|w1)&^t) >> 63
	w1 = t

	// i = 1, j = 1
	u, v = mul64(a1, b1)
	t = v + w2
	u += (v&w2 | (v|w2)&^t) >> 63
	w2 = t + c
	c = u + (t&c|(t|c)&^w2)>>63

	// i = 1, j = 2
	u, v = mul64(a2, b1)
	t = v + w3
	u += (v&w3 | (v|w3)&^t) >> 63
	w3 = t + c
	c = u + (t&c|(t|c)&^w3)>>63

	// i = 1, j = 3
	u, v = mul64(a3, b1)
	t = v + w4
	u += (v&w4 | (v|w4)&^t) >> 63
	w4 = t + c
	w5 = u + (t&c|(t|c)&^w4)>>63

	//
	// i = 2, j = 0
	c, v = mul64(a0, b2)
	t = v + w2
	c += (v&w2 | (v|w2)&^t) >> 63
	w2 = t

	// i = 2, j = 1
	u, v = mul64(a1, b2)
	t = v + w3
	u += (v&w3 | (v|w3)&^t) >> 63
	w3 = t + c
	c = u + (t&c|(t|c)&^w3)>>63

	// i = 2, j = 2
	u, v = mul64(a2, b2)
	t = v + w4
	u += (v&w4 | (v|w4)&^t) >> 63
	w4 = t + c
	c = u + (t&c|(t|c)&^w4)>>63

	// i = 2, j = 3
	u, v = mul64(a3, b2)
	t = v + w5
	u += (v&w5 | (v|w5)&^t) >> 63
	w5 = t + c
	w6 = u + (t&c|(t|c)&^w5)>>63

	//
	// i = 3, j = 0
	c, v = mul64(a0, b3)
	t = v + w3
	c += (v&w3 | (v|w3)&^t) >> 63
	w3 = t

	// i = 3, j = 1
	u, v = mul64(a1, b3)
	t = v + w4
	u += (v&w4 | (v|w4)&^t) >> 63
	w4 = t + c
	c = u + (t&c|(t|c)&^w4)>>63

	// i = 3, j = 2
	u, v = mul64(a2, b3)
	t = v + w5
	u += (v&w5 | (v|w5)&^t) >> 63
	w5 = t + c
	c = u + (t&c|(t|c)&^w5)>>63

	// i = 3, j = 3
	u, v = mul64(a3, b3)
	t = v + w6
	u += (v&w6 | (v|w6)&^t) >> 63
	w6 = t + c
	w7 = u + (t&c|(t|c)&^w6)>>63

	w[0] = w0
	w[1] = w1
	w[2] = w2
	w[3] = w3
	w[4] = w4
	w[5] = w5
	w[6] = w6
	w[7] = w7
}

// Handbook of Applied Cryptography
// Hankerson, Menezes, Vanstone
// 14.16 Algorithm Multiple-precision squaring
func square256(w *[8]uint64, a *FieldElement) {

	var w0, w1, w2, w3, w4, w5, w6, w7 uint64
	var u, v, c, vv, uu, z1, z2, z3, e uint64
	var a0 = a[0]
	var a1 = a[1]
	var a2 = a[2]
	var a3 = a[3]

	// i = 0
	c, w0 = square64(a0)

	// i = 0, j = 1
	u, v = mul64(a0, a1)
	z1 = u >> 63 // z1 for w2
	u = u<<1 + v>>63
	v = v << 1
	w1 = v + c
	e = (v&c | (v|c)&^w1) >> 63
	uu = u + e
	z1 += (u&e | (u|e)&^uu) >> 63
	c = uu

	// i = 0, j = 2
	u, v = mul64(a0, a2)
	z2 = u >> 63 // z2 for w3
	u = u<<1 + v>>63
	v = v << 1
	w2 = v + c
	e = z1 + (v&c|(v|c)&^w2)>>63
	uu = u + e
	z2 += (u&e | (u|e)&^uu) >> 63
	c = uu

	// i = 0, j = 3
	u, v = mul64(a0, a3)
	z1 = u >> 63 // z1 for w4
	u = u<<1 + v>>63
	v = v << 1
	w3 = v + c
	e = z2 + (v&c|(v|c)&^w3)>>63
	w4 = u + e
	z1 += (u&e | (u|e)&^w4) >> 63

	// i = 1
	c, v = square64(a1)
	vv = v + w2
	c += (v&w2 | (v|w2)&^vv) >> 63
	w2 = vv

	// i = 1, j = 2
	u, v = mul64(a1, a2)
	z2 = u >> 63 // z2 for w4
	u = u<<1 + v>>63
	v = v << 1
	vv = v + w3
	e = (v&w3 | (v|w3)&^vv) >> 63
	uu = u + e
	z2 += (u&e | (u|e)&^uu) >> 63
	w3 = vv + c
	e = (vv&c | (vv|c)&^w3) >> 63
	c = uu + e
	z2 += (uu&e | (uu|e)&^c) >> 63

	// i = 1, j = 3
	u, v = mul64(a1, a3)
	z3 = u >> 63 // z3 for w5
	u = u<<1 + v>>63
	v = v << 1
	vv = v + w4
	e = z1 + z2 + (v&w4|(v|w4)&^vv)>>63
	uu = u + e
	z3 += (u&e | (u|e)&^uu) >> 63
	w4 = vv + c
	e = (vv&c | (vv|c)&^w4) >> 63
	w5 = uu + e
	z3 += (uu&e | (uu|e)&^w5) >> 63

	// i = 2
	c, v = square64(a2)
	vv = v + w4
	c += (v&w4 | (v|w4)&^vv) >> 63
	w4 = vv

	// i = 2, j = 3
	u, v = mul64(a2, a3)
	z1 = u >> 63 // z1 for w6
	u = u<<1 + v>>63
	v = v << 1
	vv = v + w5
	e = z3 + (v&w5|(v|w5)&^vv)>>63
	uu = u + e
	z1 += (u&e | (u|e)&^uu) >> 63
	v = vv
	u = uu
	w5 = vv + c
	e = (vv&c | (vv|c)&^w5) >> 63
	w6 = uu + e
	z1 += (uu&e | (uu|e)&^w6) >> 63

	// i = 3
	c, v = square64(a3)
	vv = v + w6
	w7 = c + z1 + (v&w6|(v|w6)&^vv)>>63
	w6 = vv

	w[0] = w0
	w[1] = w1
	w[2] = w2
	w[3] = w3
	w[4] = w4
	w[5] = w5
	w[6] = w6
	w[7] = w7
}

// Reduces T as T (R^-1) modp
// Handbook of Applied Cryptography
// Hankerson, Menezes, Vanstone
// Algorithm 14.32 Montgomery reduction
func mont(c *FieldElement, w *[8]uint64) {
	w0 := w[0]
	w1 := w[1]
	w2 := w[2]
	w3 := w[3]
	w4 := w[4]
	w5 := w[5]
	w6 := w[6]
	w7 := w[7]
	p0 := modulus[0]
	p1 := modulus[1]
	p2 := modulus[2]
	p3 := modulus[3]
	var e1, e2, el, res uint64
	var t1, t2, u uint64

	// i = 0
	u = w0 * inp
	//
	e1, res = mul64(u, p0)
	t1 = res + w0
	e1 += (res&w0 | (res|w0)&^t1) >> 63
	w0 = t1
	//
	e2, res = mul64(u, p1)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w1
	e2 += (t1&w1 | (t1|w1)&^t2) >> 63
	w1 = t2
	//
	e1, res = mul64(u, p2)
	t1 = res + e2
	e1 += (res&e2 | (res|e2)&^t1) >> 63
	t2 = t1 + w2
	e1 += (t1&w2 | (t1|w2)&^t2) >> 63
	w2 = t2
	//
	e2, res = mul64(u, p3)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w3
	e2 += (t1&w3 | (t1|w3)&^t2) >> 63
	w3 = t2
	//
	t1 = w4 + el
	e1 = (w4&el | (w4|el)&^t1) >> 63
	t2 = t1 + e2
	e1 += (t1&e2 | (t1|e2)&^t2) >> 63
	w4 = t2
	el = e1

	// i = 1
	u = w1 * inp
	//
	e1, res = mul64(u, p0)
	t1 = res + w1
	e1 += (res&w1 | (res|w1)&^t1) >> 63
	w1 = t1
	//
	e2, res = mul64(u, p1)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w2
	e2 += (t1&w2 | (t1|w2)&^t2) >> 63
	w2 = t2
	//
	e1, res = mul64(u, p2)
	t1 = res + e2
	e1 += (res&e2 | (res|e2)&^t1) >> 63
	t2 = t1 + w3
	e1 += (t1&w3 | (t1|w3)&^t2) >> 63
	w3 = t2
	//
	e2, res = mul64(u, p3)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w4
	e2 += (t1&w4 | (t1|w4)&^t2) >> 63
	w4 = t2
	//
	t1 = w5 + el
	e1 = (w5&el | (w5|el)&^t1) >> 63
	t2 = t1 + e2
	e1 += (t1&e2 | (t1|e2)&^t2) >> 63
	w5 = t2
	el = e1

	// i = 2
	u = w2 * inp
	//
	e1, res = mul64(u, p0)
	t1 = res + w2
	e1 += (res&w2 | (res|w2)&^t1) >> 63
	w2 = t1
	//
	e2, res = mul64(u, p1)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w3
	e2 += (t1&w3 | (t1|w3)&^t2) >> 63
	w3 = t2
	//
	e1, res = mul64(u, p2)
	t1 = res + e2
	e1 += (res&e2 | (res|e2)&^t1) >> 63
	t2 = t1 + w4
	e1 += (t1&w4 | (t1|w4)&^t2) >> 63
	w4 = t2
	//
	e2, res = mul64(u, p3)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w5
	e2 += (t1&w5 | (t1|w5)&^t2) >> 63
	w5 = t2
	//
	t1 = w6 + el
	e1 = (w6&el | (w6|el)&^t1) >> 63
	t2 = t1 + e2
	e1 += (t1&e2 | (t1|e2)&^t2) >> 63
	w6 = t2
	el = e1

	// i = 3
	u = w3 * inp
	//
	e1, res = mul64(u, p0)
	t1 = res + w3
	e1 += (res&w3 | (res|w3)&^t1) >> 63
	w3 = t1
	//
	e2, res = mul64(u, p1)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w4
	e2 += (t1&w4 | (t1|w4)&^t2) >> 63
	w4 = t2
	//
	e1, res = mul64(u, p2)
	t1 = res + e2
	e1 += (res&e2 | (res|e2)&^t1) >> 63
	t2 = t1 + w5
	e1 += (t1&w5 | (t1|w5)&^t2) >> 63
	w5 = t2
	//
	e2, res = mul64(u, p3)
	t1 = res + e1
	e2 += (res&e1 | (res|e1)&^t1) >> 63
	t2 = t1 + w6
	e2 += (t1&w6 | (t1|w6)&^t2) >> 63
	w6 = t2
	//
	t1 = w7 + el
	e1 = (w7&el | (w7|el)&^t1) >> 63
	t2 = t1 + e2
	e1 += (t1&e2 | (t1|e2)&^t2) >> 63
	w7 = t2

	e1--
	c[0] = w4 - ((p0) & ^e1)
	e2 = (^w4&p0 | (^w4|p0)&c[0]) >> 63
	c[1] = w5 - ((p1 + e2) & ^e1)
	e2 = (^w5&p1 | (^w5|p1)&c[1]) >> 63
	c[2] = w6 - ((p2 + e2) & ^e1)
	e2 = (^w6&p2 | (^w6|p2)&c[2]) >> 63
	c[3] = w7 - ((p3 + e2) & ^e1)

	sub(c, c, modulus)
}
