This library covers arithmetic operations for prime fields upto 256 bits in Go language. Field is implemented as 4 x 64 bit fashion. Currently, addition, subtraction, multiplication, squaring and inversion operations are supported. Multiplication, squaring and invertion operations are done in Montgomery space.

### Usage

#### Field

New field can be created from standart big.Int prime number. Montgomery constants are precomputed during construction of new field.

```go
pBig := new(big.Int).SetString("0x73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001",16)
field := NewField(pBig)
```

#### Field Element

New field element can be created using bytes input. Given 32 bytes input new field element is transformed into Montgomery domain.

```go
// input bytes with big endian order
feBytes := []byte{12, 14, 250, ... }
fe := field.NewElement(fe)

// random element
fe2 := new(FieldElement)
field.RandElement(fe2, rand.Reader)
```

### Benchmarks

Prime field operations benchmarked on _2,7 GHz Intel Core i5_.

```
BenchmarkFieldAddition                    14.1 ns/op
BenchmarkFieldSubtraction                 13.1 ns/op
BenchmarkFieldMontgomeryReduction         98.1 ns/op
BenchmarkFieldMontgomeryMultiplication    170 ns/op
BenchmarkFieldMontgomerySquaring          162 ns/op
BenchmarkFieldInverse                     5202 ns/op
```

### Todo

* Exponentiation
* Square Root

### References

#### Books and papers

* [Handbook of Applied Cryptography](http://cacr.uwaterloo.ca/hac/)
* [Guide to Elliptic Curve Cryptography](https://www.springer.com/gp/book/9780387952734)
* [Efficient Software-Implementation of Finite Fields with Applications to Cryptography](https://www.researchgate.net/publication/225962646_Efficient_Software-Implementation_of_Finite_Fields_with_Applications_to_Cryptography)
* [The Montgomery Modular Inverse - Revisited](https://ieeexplore.ieee.org/abstract/document/863048)

#### Related or reference libraries

* [cloudflare/bn256](https://github.com/ethereum/go-ethereum/tree/master/crypto/bn256)
* [zkcrypto/jubjub](https://github.com/zkcrypto/jubjub/blob/master/src/fq.rs)
* [matter-labs/eip1829](https://github.com/matter-labs/eip1829/blob/master/src/field.rs)



