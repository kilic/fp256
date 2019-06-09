package fp

import (
	"fmt"
)

type FieldElement [4]uint64

func (fe *FieldElement) String() string {
	return fmt.Sprintf("0x%16.16x%16.16x%16.16x%16.16x", fe[3], fe[2], fe[1], fe[0])
}

// Doesn't apply montgomery reduction
func (fe *FieldElement) Marshal(out []byte) {
	var a int
	for i := 0; i < 4; i++ {
		a = 31 - i*8
		out[a] = byte(fe[i])
		out[a-1] = byte(fe[i] >> 8)
		out[a-2] = byte(fe[i] >> 16)
		out[a-3] = byte(fe[i] >> 24)
		out[a-4] = byte(fe[i] >> 32)
		out[a-5] = byte(fe[i] >> 40)
		out[a-6] = byte(fe[i] >> 48)
		out[a-7] = byte(fe[i] >> 56)
	}
}

// Doesn't apply montgomery encoding
func (fe *FieldElement) Unmarshal(in []byte) *FieldElement {
	padded := make([]byte, 32)
	l := len(in)
	if l >= 32 {
		l = 32
	}
	copy(padded[32-l:], in[:])
	var a int
	for i := 0; i < 4; i++ {
		a = 31 - i*8
		fe[i] = uint64(padded[a]) | uint64(padded[a-1])<<8 |
			uint64(padded[a-2])<<16 | uint64(padded[a-3])<<24 |
			uint64(padded[a-4])<<32 | uint64(padded[a-5])<<40 |
			uint64(padded[a-6])<<48 | uint64(padded[a-7])<<56
	}
	return fe
}

func (fe *FieldElement) set(a *FieldElement) *FieldElement {
	fe[0] = a[0]
	fe[1] = a[1]
	fe[2] = a[2]
	fe[3] = a[3]
	return fe
}

func (fe *FieldElement) isEven() bool {
	const mask uint64 = 1
	return fe[0]&mask == 0
}

func (fe *FieldElement) isOne() bool {
	return 1 == fe[0] && 0 == fe[1] && 0 == fe[2] && 0 == fe[3]
}

func (fe *FieldElement) isZero() bool {
	return 0 == fe[0] && 0 == fe[1] && 0 == fe[2] && 0 == fe[3]
}

func (fe *FieldElement) eq(e *FieldElement) bool {
	return e[0] == fe[0] && e[1] == fe[1] && e[2] == fe[2] && e[3] == fe[3]
}

func (fe *FieldElement) cmp(fe2 *FieldElement) int64 {
	if fe[3] > fe2[3] {
		return 1
	} else if fe[3] < fe2[3] {
		return -1
	}
	if fe[2] > fe2[2] {
		return 1
	} else if fe[2] < fe2[2] {
		return -1
	}
	if fe[1] > fe2[1] {
		return 1
	} else if fe[1] < fe2[1] {
		return -1
	}
	if fe[0] > fe2[0] {
		return 1
	} else if fe[0] < fe2[0] {
		return -1
	}
	return 0
}

func (fe *FieldElement) rightShift(e uint64) {
	fe[0] = fe[0]>>1 | fe[1]<<63
	fe[1] = fe[1]>>1 | fe[2]<<63
	fe[2] = fe[2]>>1 | fe[3]<<63
	fe[3] = fe[3]>>1 | e<<63
}

func (fe *FieldElement) leftShift() uint64 {
	e := fe[3] >> 63
	fe[3] = fe[3]<<1 | fe[2]>>63
	fe[2] = fe[2]<<1 | fe[1]>>63
	fe[1] = fe[1]<<1 | fe[0]>>63
	fe[0] = fe[0] << 1
	return e
}
