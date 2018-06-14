# libhash

A library making hashing more conventient. 

```go
h, err := libhash.HashFile("sha1", "/etc/passwd")

if err != nil {
	fmt.Println(h)
}

fmt.Println(libhash.HashStr("md5", "Hello, world!"))
```

## Supported Hash Algorithms

 * sha256
 * sha224
 * md5
 * adler32
 * crc32, crc32/ieee
 * crc32/castagnoli
 * crc32/koopman
 * crc64, crc64/iso
 * crc64/ecma
 * fnv/128
 * fnv/128a
 * fnv/64
 * fnv/64a
 * fnv/32
 * fnv/32a
 * sha1
 * sha512
 * sha384
 * sha512/224, sha512-224
 * sha512/256, sha512-256
 * blake2b/256, blake2b-256
 * blake2b/384, blake2b-384
 * blake2b/512, blake2b-512
 * md4
 * sha3/224, sha3-224
 * sha3/256, sha3-256
 * sha3/384, sha3-384
 * sha3/512, sha3-512
 * blake2s/256, blake2s-256
 * sum/8, sum-8
 * sum/16, sum-16
 * sum/32, sum-32
 * sum/64, sum-64
 * bsdsum/8, bsdsum-8
 * bsdsum/16, bsdsum-16
 * bsdsum/32, bsdsum-32
 * bsdsum/64, bsdsum-64
 * sysv
 * fletcher/16, fletcher-16, fletcher16
 * xor8, xor/8, xor-8
