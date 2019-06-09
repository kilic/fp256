// +build amd64, !pure_go

package fp

// go:noescape
func add(c, a, b *FieldElement)

// go:noescape
func addn(a, b *FieldElement) uint64

// go:noescape
func double(c, a *FieldElement)

// go:noescape
func neg(c, a *FieldElement)

// go:noescape
func sub(c, a, b *FieldElement)

// go:noescape
func subn(a, b *FieldElement) uint64

// go:noescape
func mul(w *[8]uint64, a, b *FieldElement)

// go:noescape
func mont(c *FieldElement, w *[8]uint64)

// go:noescape
func montmul(c, a, b *FieldElement)

// go:noescape
func square(w *[8]uint64, a *FieldElement)

// go:noescape
func montsquare(c, a *FieldElement)
