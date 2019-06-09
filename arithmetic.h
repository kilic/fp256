#define mul(rc, ra, rb) \
  XORQ R10, R10 \
	XORQ R11, R11 \
	XORQ R12, R12 \
	XORQ R13, R13 \
	XORQ R14, R14 \
	XORQ R15, R15 \
	MOVQ 0+rb, CX \
	MOVQ 0+ra, AX \
	MULQ CX \
	MOVQ AX, 0+rc \
	MOVQ DX, R9 \
	\ // a = 1, b = 0
	MOVQ 8+ra, AX \
	MULQ CX \
	ADDQ AX, R9 \
	ADCQ DX, R10 \
	\ // a = 2, b = 0
	MOVQ 16+ra, AX \
	MULQ CX \
	ADDQ AX, R10 \
	ADCQ DX, R11 \
	\ // a = 3, b = 0
	MOVQ 24+ra, AX \
	MULQ CX \
	ADDQ AX, R11 \
	ADCQ DX, R12 \
	\ //
  MOVQ 8+rb, CX \
	\ // a = 0, b = 1
  MOVQ 0+ra, AX \
	MULQ CX \
	ADDQ AX, R9 \
	ADCQ DX, R10 \
	ADCQ $0, R11 \
	ADCQ $0, R12 \
	\ // a = 1, b = 1
  MOVQ 8+ra, AX \
	MULQ CX \
	ADDQ AX, R10 \
	ADCQ DX, R11 \
	ADCQ $0, R12 \
	ADCQ $0, R13 \
	\ // a = 2, b = 1
  MOVQ 16+ra, AX \
	MULQ CX \
	ADDQ AX, R11 \
	ADCQ DX, R12 \
	ADCQ $0, R13 \
	ADCQ $0, R14 \
	\ // a = 3, b = 1
  MOVQ 24+ra, AX \
	MULQ CX \
	ADDQ AX, R12 \
	ADCQ DX, R13 \
	ADCQ $0, R14 \
	ADCQ $0, R15 \
	\ //
	MOVQ 16+rb, CX \
	\ // a = 0, b = 2
  MOVQ 0+ra, AX \
	MULQ CX \
	ADDQ AX, R10 \
	ADCQ DX, R11 \
	ADCQ $0, R12 \
	ADCQ $0, R13 \
	\ // a = 1, b = 2
  MOVQ 8+ra, AX \
	MULQ CX \
	ADDQ AX, R11 \
	ADCQ DX, R12 \
	ADCQ $0, R13 \
	ADCQ $0, R14 \
	\ // a = 2, b = 2
  MOVQ 16+ra, AX \
	MULQ CX \
	ADDQ AX, R12 \
	ADCQ DX, R13 \
	ADCQ $0, R14 \
	ADCQ $0, R15 \
	\ // a = 3, b = 2
  MOVQ 24+ra, AX \
	MULQ CX \
	ADDQ AX, R13 \
	ADCQ DX, R14 \
	ADCQ $0, R15 \
	\ //
	MOVQ 24+rb, CX \
	\ //a = 0, b = 3
  MOVQ 0+ra, AX \
	MULQ CX \
  ADDQ AX, R11 \
	ADCQ DX, R12 \
	ADCQ $0, R13 \
	ADCQ $0, R14 \
	\ // a = 1, b = 3
  MOVQ 8+ra, AX \
	MULQ CX \
	ADDQ AX, R12 \
	ADCQ DX, R13 \
	ADCQ $0, R14 \
	ADCQ $0, R15 \
	\ // a = 2, b = 3
  MOVQ 16+ra, AX \
	MULQ CX \
	ADDQ AX, R13 \
	ADCQ DX, R14 \
	ADCQ $0, R15 \
	\ // a = 3, b = 3
  MOVQ 24+ra, AX \
	MULQ CX \
	ADDQ AX, R14 \
	ADCQ DX, R15 \
  \ //
	MOVQ R9, 8+rc \
	MOVQ R10, 16+rc \
	MOVQ R11, 24+rc \
	MOVQ R12, 32+rc \
	MOVQ R13, 40+rc \
	MOVQ R14, 48+rc \
	MOVQ R15, 56+rc \

#define square(rw, ra) \
  MOVQ 0+ra, R13 \
  MOVQ 8+ra, R14 \
  MOVQ 16+ra, R15 \
  MOVQ 24+ra, BX \
  \
  MOVQ R13, AX \
  MULQ R13 \
  MOVQ AX, 0+rw \
  MOVQ DX, CX \
  \
  XORQ R8, R8 \
  MOVQ R13, AX \
  MULQ R14 \
  ADDQ AX, AX \
  ADCQ DX, DX \
  ADCQ $0, R8 \
  ADDQ CX, AX \
  ADCQ $0, DX \
  MOVQ AX, 8+rw \
  MOVQ DX, CX \
  \
  XORQ R9, R9 \
  MOVQ R13, AX \
  MULQ R15 \
  ADDQ AX, AX \
  ADCQ DX, DX \
  ADCQ $0, R9 \
  ADDQ CX, AX \
  ADCQ R8, DX \
  ADCQ $0, R9 \
  MOVQ AX, R10 \
  MOVQ DX, CX \
  \
  XORQ R8, R8 \
  MOVQ R13, AX \
  MULQ BX \
  ADDQ AX, AX \
  ADCQ DX, DX \
  ADCQ $0, R8 \
  ADDQ CX, AX \
  ADCQ R9, DX \
  ADCQ $0, R8 \
  MOVQ AX, R11 \
  MOVQ DX, R12 \
  \
  XORQ CX, CX \
  MOVQ R14, AX \
  MULQ R14 \
  ADDQ AX, R10 \
  ADCQ DX, CX \
  MOVQ R10, 16+rw \
  \
  XORQ R9, R9 \
  MOVQ R14, AX \
  MULQ R15 \
  ADDQ AX, AX \
  ADCQ DX, DX \
  ADCQ $0, R9 \
  ADDQ CX, AX \
  ADCQ $0, DX \
  ADDQ AX, R11 \
  ADCQ $0, DX \
  ADCQ $0, R9 \
  MOVQ DX, CX \
  MOVQ R11, 24+rw \
  \
  XORQ R10, R10 \
  MOVQ R14, AX \
  MULQ BX \
  ADDQ AX, AX \
  ADCQ DX, DX \
  ADCQ $0, R10 \
  ADDQ CX, AX \
  ADCQ R8, DX \
  ADCQ $0, R10 \
  ADDQ AX, R12 \
  ADCQ R9, DX \
  ADCQ $0, R10 \
  MOVQ DX, R13 \
  \
  XORQ CX, CX \
  MOVQ R15, AX \
  MULQ R15 \
  ADDQ AX, R12 \
  ADCQ DX, CX \
  MOVQ R12, 32+rw \
  \
  XORQ R8, R8 \
  MOVQ R15, AX \
  MULQ BX \
  ADDQ AX, AX \
  ADCQ DX, DX \
  ADCQ $0, R8 \
  ADDQ CX, AX \
  ADCQ R10, DX \
  ADCQ $0, R8   \
  ADDQ AX, R13 \
  ADCQ $0, DX \
  ADCQ $0, R8 \
  MOVQ DX, R14 \
  MOVQ R13, 40+rw \
  \
  MOVQ BX, AX \
  MULQ BX \
  ADDQ AX, R14 \
  ADCQ R8, DX \
  MOVQ R14, 48+rw \
  MOVQ DX, 56+rw \

#define mont(rc, rw) \
  \ // i = 0
  MOVQ 0+rw, R8 \ // w0
  MOVQ 8+rw, R9 \ // w1
  MOVQ 16+rw, R10 \ // w2
  MOVQ 24+rw, R11 \ // w3
  MOVQ 32+rw, R12 \ // w4
  \ // u @ CX = w0 * inp
  MOVQ R8, AX \
  MULQ ·inp(SB) \
  MOVQ AX, CX \
  \ // w0
  XORQ R14, R14 \
  MOVQ ·modulus+0(SB), AX \
  MULQ CX \
  ADDQ AX, R8 \
  ADCQ DX, R14 \
  \ // w1
  XORQ R13, R13 \
  MOVQ ·modulus+8(SB), AX \
  MULQ CX \
  ADDQ AX, R9 \
  ADCQ DX, R13 \
  ADDQ R14, R9 \
  ADCQ $0, R13 \
  \ // w2
  XORQ R14, R14 \
  MOVQ ·modulus+16(SB), AX \
  MULQ CX \
  ADDQ AX, R10 \
  ADCQ DX, R14 \
  ADDQ R13, R10 \
  ADCQ $0, R14 \ 
  \ // w3
  XORQ R13, R13 \
  MOVQ ·modulus+24(SB), AX \
  MULQ CX \
  ADDQ AX, R11 \
  ADCQ DX, R13 \
  ADDQ R14, R11 \
  ADCQ $0, R13 \
  \ // w4
  XORQ R15, R15 \
  ADDQ R13, R12 \
  ADCQ $0, R15 \
  \ //
  \ // i = 1
  \ // R9 // w1
  \ // R10 // w2
  \ // R11 // w3
  \ // R12 // w4
  MOVQ 40+rw, R8 \ // w5
  \ // u @ CX = w1 * inp
  MOVQ R9, AX \
  MULQ ·inp(SB) \
  MOVQ AX, CX \
  \ // w1
  XORQ R14, R14 \
  MOVQ ·modulus+0(SB), AX \
  MULQ CX \
  ADDQ AX, R9 \
  ADCQ DX, R14 \ 
  \ // w2
  XORQ R13, R13 \
  MOVQ ·modulus+8(SB), AX \
  MULQ CX \
  ADDQ AX, R10 \
  ADCQ DX, R13 \
  ADDQ R14, R10 \
  ADCQ $0, R13 \
  \ // w3
  XORQ R14, R14 \
  MOVQ ·modulus+16(SB), AX \
  MULQ CX \
  ADDQ AX, R11 \
  ADCQ DX, R14 \
  ADDQ R13, R11 \
  ADCQ $0, R14 \
  \ // w4
  XORQ R13, R13 \
  MOVQ ·modulus+24(SB), AX \
  MULQ CX \
  ADDQ AX, R12 \
  ADCQ DX, R13 \
  ADDQ R14, R12 \
  ADCQ $0, R13 \
  \ // w5
  ADDQ R13, R15 \
  ADCQ R15, R8 \
  MOVQ $0, R15 \
  ADCQ $0, R15 \
  \ // i = 2
  \ // R10 // w2
  \ // R11 // w3
  \ // R12 // w4
  \ // R8 // w5
  MOVQ 48+rw, R9 \ // w6
  \ // u @ CX = w2 * inp
  MOVQ R10, AX \
  MULQ ·inp(SB) \
  MOVQ AX, CX \
  \ // w2
  XORQ R14, R14 \
  MOVQ ·modulus+0(SB), AX \
  MULQ CX \
  ADDQ AX, R10 \
  ADCQ DX, R14 \
  \ // w3
  XORQ R13, R13 \
  MOVQ ·modulus+8(SB), AX \
  MULQ CX \
  ADDQ AX, R11 \
  ADCQ DX, R13 \
  ADDQ R14, R11 \
  ADCQ $0, R13 \
  \ // w4
  XORQ R14, R14 \
  MOVQ ·modulus+16(SB), AX \
  MULQ CX \
  ADDQ AX, R12 \
  ADCQ DX, R14 \
  ADDQ R13, R12 \
  ADCQ $0, R14 \
  \ // w5
  XORQ R13, R13 \
  MOVQ ·modulus+24(SB), AX \
  MULQ CX \
  ADDQ AX, R8 \
  ADCQ DX, R13 \
  ADDQ R14, R8 \
  ADCQ $0, R13 \
  \ // w6
  ADDQ R13, R15 \
  ADCQ R15, R9 \
  MOVQ $0, R15 \
  ADCQ $0, R15 \
  \ // i = 3
  \ // R11 // w3
  \ // R12 // w4
  \ // R8 // w5
  \ // R9 // w6
  MOVQ 56+rw, R10 \ // w7
  \ // u @ CX = w3 * inp
  MOVQ R11, AX \
  MULQ ·inp(SB) \
  MOVQ AX, CX \
  \ // w3
  XORQ R14, R14 \
  MOVQ ·modulus+0(SB), AX \
  MULQ CX \
  ADDQ AX, R11 \
  ADCQ DX, R14 \
  \ // w4
  XORQ R13, R13 \
  MOVQ ·modulus+8(SB), AX \
  MULQ CX \
  ADDQ AX, R12 \
  ADCQ DX, R13 \
  ADDQ R14, R12 \
  ADCQ $0, R13 \
  \ // w5
  XORQ R14, R14 \
  MOVQ ·modulus+16(SB), AX \
  MULQ CX \
  ADDQ AX, R8 \
  ADCQ DX, R14 \
  ADDQ R13, R8 \
  ADCQ $0, R14 \
  \ // w6
  XORQ R13, R13 \
  MOVQ ·modulus+24(SB), AX \
  MULQ CX \
  ADDQ AX, R9 \
  ADCQ DX, R13 \
  ADDQ R14, R9 \
  ADCQ $0, R13 \
  \ // w7
  ADDQ R13, R15 \
  ADCQ R15, R10 \
  MOVQ $0, R15 \
  ADCQ $0, R15 \
  \
  MOVQ R12, R11 \
  MOVQ R8, R13 \
  MOVQ R9, R14 \
  MOVQ R10, AX \
  \
  SUBQ ·modulus+0(SB), R12 \
  SBBQ ·modulus+8(SB), R8 \
  SBBQ ·modulus+16(SB), R9 \
  SBBQ ·modulus+24(SB), R10 \
  SBBQ $0, R15 \
  \
  CMOVQCC R12, R11 \
  CMOVQCC R8, R13 \
  CMOVQCC R9, R14 \
  CMOVQCC R10, AX \
  \
  MOVQ R11, 0+rc \
  MOVQ R13, 8+rc \
  MOVQ R14, 16+rc \
  MOVQ AX, 24+rc
  