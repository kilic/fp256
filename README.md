This library contains arithmetic operations for prime fields upto 256 bit. Operations are optimized for AMD64 architecture.

### Usage

#### Field

New field can be created from standart big.Int prime number. Montgomery constants are precomputed during construction of new field.

```go
pStr = "0x73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001"
pBig := new(big.Int).SetString(pStr[:2],16)
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

Prime field operations benchmarked on _2,7 GHz i5_.

```
BenchmarkAddition                      5.90 ns/op
BenchmarkMontgomeryMultiplication      37.1 ns/op
BenchmarkMontgomerySquaring            32.9 ns/op
BenchmarkInvertion                     2829 ns/op
```

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



