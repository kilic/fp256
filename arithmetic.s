// +build amd64, !pure_go

#include "textflag.h"
#include "arithmetic.h"

TEXT ·double(SB), NOSPLIT, $0-16
  MOVQ a+8(FP), DI
	MOVQ 0(DI), R8
	MOVQ 8(DI), R9
	MOVQ 16(DI), R10
	MOVQ 24(DI), R11
  XORQ R12, R12
	ADDQ R8, R8
	ADCQ R9, R9
	ADCQ R10, R10
	ADCQ R11, R11
	ADCQ $0, R12
  MOVQ R8, R13
	MOVQ R9, R14
	MOVQ R10, R15
	MOVQ R11, AX
	MOVQ R12, BX
	SUBQ ·modulus+0(SB), R13
	SBBQ ·modulus+8(SB), R14
	SBBQ ·modulus+16(SB), R15
	SBBQ ·modulus+24(SB), AX
	SBBQ $0, BX
	CMOVQCC R13, R8
	CMOVQCC R14, R9
	CMOVQCC R15, R10
	CMOVQCC AX, R11
	MOVQ c+0(FP), DI
	MOVQ R8, 0(DI)
	MOVQ R9, 8(DI)
	MOVQ R10, 16(DI)
	MOVQ R11, 24(DI)
  RET

TEXT ·addn(SB), NOSPLIT, $0-24
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ 0(DI), R8
	MOVQ 8(DI), R9
	MOVQ 16(DI), R10
	MOVQ 24(DI), R11
	MOVQ $0, R12
	ADDQ  0(SI), R8
	ADCQ  8(SI), R9
	ADCQ 16(SI), R10
	ADCQ 24(SI), R11
	ADCQ $0, R12
	MOVQ R8, 0(DI)
	MOVQ R9, 8(DI)
	MOVQ R10, 16(DI)
	MOVQ R11, 24(DI)
	MOVQ R12, ret+16(FP)
	RET

TEXT ·add(SB), NOSPLIT, $0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	MOVQ 0(DI), R8
	MOVQ 8(DI), R9
	MOVQ 16(DI), R10
	MOVQ 24(DI), R11
	MOVQ $0, R12
	ADDQ  0(SI), R8
	ADCQ  8(SI), R9
	ADCQ 16(SI), R10
	ADCQ 24(SI), R11
	ADCQ $0, R12
	MOVQ R8, R13
	MOVQ R9, R14
	MOVQ R10, R15
	MOVQ R11, AX
	MOVQ R12, BX
	SUBQ ·modulus+0(SB), R13
	SBBQ ·modulus+8(SB), R14
	SBBQ ·modulus+16(SB), R15
	SBBQ ·modulus+24(SB), AX
	SBBQ $0, BX
	CMOVQCC R13, R8
	CMOVQCC R14, R9
	CMOVQCC R15, R10
	CMOVQCC AX, R11
  MOVQ c+0(FP), DI
	MOVQ R8, 0(DI)
	MOVQ R9, 8(DI)
	MOVQ R10, 16(DI)
	MOVQ R11, 24(DI)
	RET

TEXT ·neg(SB), NOSPLIT, $0-16
  MOVQ ·modulus+0(SB), R8
	MOVQ ·modulus+8(SB), R9
	MOVQ ·modulus+16(SB), R10
	MOVQ ·modulus+24(SB), R11
  MOVQ a+8(FP), DI
	SUBQ 0(DI), R8
	SBBQ 8(DI), R9
	SBBQ 16(DI), R10
	SBBQ 24(DI), R11
  MOVQ c+0(FP), DI
  MOVQ R8, 0(DI)
  MOVQ R9, 8(DI)
  MOVQ R10, 16(DI)
  MOVQ R11, 24(DI)
  RET

TEXT ·sub(SB), NOSPLIT, $0-24
  MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	MOVQ 0(DI), R8
	MOVQ 8(DI), R9
	MOVQ 16(DI), R10
	MOVQ 24(DI), R11
  MOVQ ·modulus+0(SB), R12
	MOVQ ·modulus+8(SB), R13
	MOVQ ·modulus+16(SB), R14
	MOVQ ·modulus+24(SB), R15
  MOVQ $0, AX
	SUBQ  0(SI), R8
	SBBQ  8(SI), R9
	SBBQ 16(SI), R10
	SBBQ 24(SI), R11
  CMOVQCC AX, R12
	CMOVQCC AX, R13
	CMOVQCC AX, R14
	CMOVQCC AX, R15
	ADDQ R12, R8
	ADCQ R13, R9
	ADCQ R14, R10
	ADCQ R15, R11
  MOVQ c+0(FP), DI
	MOVQ R8, 0(DI)
	MOVQ R9, 8(DI)
	MOVQ R10, 16(DI)
	MOVQ R11, 24(DI)
  RET


TEXT ·subn(SB), NOSPLIT, $0-24
  MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ 0(DI), R8
	MOVQ 8(DI), R9
	MOVQ 16(DI), R10
	MOVQ 24(DI), R11
	MOVQ $0, R12
	SUBQ  0(SI), R8
	SBBQ  8(SI), R9
	SBBQ 16(SI), R10
	SBBQ 24(SI), R11
	ADCQ $0, R12
	MOVQ R8, 0(DI)
	MOVQ R9, 8(DI)
	MOVQ R10, 16(DI)
	MOVQ R11, 24(DI)
	MOVQ R12, ret+16(FP)
  RET

TEXT ·mul(SB),NOSPLIT, $0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	MOVQ w+0(FP), BX
  mul(0(BX), 0(SI), 0(DI))
  RET

TEXT ·square(SB),NOSPLIT, $0-16
	MOVQ a+8(FP), DI
	MOVQ w+0(FP), SI
  square(0(SI), 0(DI))
  RET
  
TEXT ·mont(SB),NOSPLIT, $0-16
	MOVQ w+8(FP), DI
	MOVQ c+0(FP), SI
  mont(0(SI), 0(DI))
  RET

TEXT ·montmul(SB),NOSPLIT, $56-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
  mul(0(SP), 0(SI), 0(DI))
  MOVQ c+0(FP), DI
  mont(0(DI), 0(SP))
  RET

TEXT ·montsquare(SB),NOSPLIT, $56-16
	MOVQ a+8(FP), DI
  square(0(SP), 0(DI))
  MOVQ c+0(FP), DI
  mont(0(DI), 0(SP))
  RET

TEXT ·inverse(SB), NOSPLIT, $128-16
  // v: 0(SP), u: 32(SP), 
  // x1: 64(SP), x2: 96(SP)
  MOVQ ·modulus+0(SB), R8
  MOVQ R8, 0(SP)
  MOVQ ·modulus+8(SB), R8
  MOVQ R8, 8(SP)
  MOVQ ·modulus+16(SB), R8
  MOVQ R8, 16(SP)
  MOVQ ·modulus+24(SB), R8
  MOVQ R8, 24(SP)
  MOVQ a+8(FP), DI
  move(0(DI), 32(SP))
  clear(64(SP))
  MOVQ $1, 64(SP)
  clear(96(SP))
  XORQ R15, R15
  XORQ R14, R14
  XORQ R12, R12
loop:
  INCQ R15
  iszero(0(SP))
  JEQ final
  MOVQ 0(SP), AX
  TESTQ $1, AX
  JEQ v_is_even
  MOVQ 32(SP), AX
  TESTQ $1, AX
  JEQ u_is_even
  lt(0(SP), 32(SP))
  JLO u_gt_v
  subn(0(SP), 32(SP))
  rsh(0(SP))
  addn(96(SP), 64(SP))
  lsh(64(SP))
  JMP loop
v_is_even:
  rsh(0(SP))
  lsh(64(SP))
  JMP loop
u_is_even:
  rsh(32(SP))
  lsh(96(SP))
  JMP loop
u_gt_v:
  subn(32(SP), 0(SP))
  rsh(32(SP))
  addn(64(SP), 96(SP))
  lsh(96(SP))
  JMP loop
final:
  MOVQ R15, CX
  MOVQ 64(SP), R8
	MOVQ 72(SP), R9
	MOVQ 80(SP), R10
	MOVQ 88(SP), R11
  MOVQ ·modulus+0(SB), R12
  MOVQ ·modulus+8(SB), R13
  MOVQ ·modulus+16(SB), R14
  MOVQ ·modulus+24(SB), R15
  MOVQ $0, AX
  SUBQ R12, R8
	SBBQ R13, R9
	SBBQ R14, R10
	SBBQ R15, R11
  CMOVQCC AX, R12
	CMOVQCC AX, R13
	CMOVQCC AX, R14
	CMOVQCC AX, R15
	ADDQ R12, R8
	ADCQ R13, R9
	ADCQ R14, R10
	ADCQ R15, R11
loop2:
  XORQ R12, R12
	ADDQ R8, R8
	ADCQ R9, R9
	ADCQ R10, R10
	ADCQ R11, R11
	ADCQ $0, R12
  MOVQ R8, R13
	MOVQ R9, R14
	MOVQ R10, R15
	MOVQ R11, AX
	MOVQ R12, BX
  SUBQ ·modulus+0(SB), R13
	SBBQ ·modulus+8(SB), R14
	SBBQ ·modulus+16(SB), R15
	SBBQ ·modulus+24(SB), AX
	SBBQ $0, BX
	CMOVQCC R13, R8
	CMOVQCC R14, R9
	CMOVQCC R15, R10
	CMOVQCC AX, R11
  INCQ CX
  CMPQ CX, $513
  JNE loop2
  MOVQ c+0(FP), DI
  MOVQ R8, 0(DI)
  MOVQ R9, 8(DI)
  MOVQ R10, 16(DI)
  MOVQ R11, 24(DI)
  RET
  