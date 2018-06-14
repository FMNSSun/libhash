package libhash

import "io"
import "os"
import "hash"
import "crypto/sha256"
import "crypto/md5"
import "errors"
import "hash/adler32"
import "hash/crc32"
import "hash/crc64"
import "hash/fnv"
import "crypto/sha1"
import "crypto/sha512"
import "golang.org/x/crypto/blake2b"
import "golang.org/x/crypto/md4"
import "golang.org/x/crypto/sha3"
import "golang.org/x/crypto/blake2s"
import "encoding/hex"
import "github.com/FMNSSun/libhash/sum"
import "github.com/FMNSSun/libhash/bsdsum"
import "github.com/FMNSSun/libhash/sysvsum"
import "github.com/FMNSSun/libhash/fletcher"

func IsHashSupported(hash string) bool {
	return getHash(hash) != nil
}

func getHash(hash string) hash.Hash {
	switch hash {
	case "sha256":
		return sha256.New()

	case "sha224":
		return sha256.New224()

	case "md5":
		return md5.New()

	case "adler32":
		return adler32.New()

	case "crc32", "crc32/ieee":
		return crc32.New(crc32.MakeTable(crc32.IEEE))

	case "crc32/castagnoli":
		return crc32.New(crc32.MakeTable(crc32.Castagnoli))

	case "crc32/koopman":
		return crc32.New(crc32.MakeTable(crc32.Koopman))

	case "crc64", "crc64/iso":
		return crc64.New(crc64.MakeTable(crc64.ISO))

	case "crc64/ecma":
		return crc64.New(crc64.MakeTable(crc64.ECMA))

	case "fnv/128":
		return fnv.New128()

	case "fnv/128a":
		return fnv.New128a()

	case "fnv/64":
		return fnv.New64()

	case "fnv/64a":
		return fnv.New64a()

	case "fnv/32":
		return fnv.New32()

	case "fnv/32a":
		return fnv.New32a()

	case "sha1":
		return sha1.New()

	case "sha512":
		return sha512.New()

	case "sha384":
		return sha512.New384()

	case "sha512/224", "sha512-224":
		return sha512.New512_224()

	case "sha512/256", "sha512-256":
		return sha512.New512_256()

	case "blake2b/256", "blake2b-256":
		h, err := blake2b.New256(nil)

		if err != nil {
			return nil
		}

		return h

	case "blake2b/384", "blake2b-384":
		h, err := blake2b.New384(nil)

		if err != nil {
			return nil
		}

		return h

	case "blake2b/512", "blake2b-512":
		h, err := blake2b.New512(nil)

		if err != nil {
			return nil
		}

		return h

	case "md4":
		return md4.New()

	case "sha3/224", "sha3-224":
		return sha3.New224()

	case "sha3/256", "sha3-256":
		return sha3.New256()

	case "sha3/384", "sha3-384":
		return sha3.New384()

	case "sha3/512", "sha3-512":
		return sha3.New512()

	case "blake2s/256", "blake2s-256":
		h, err := blake2s.New256(nil)

		if err != nil {
			return nil
		}

		return h

	case "sum/8", "sum-8":
		return sum.New8()

	case "sum/16", "sum-16":
		return sum.New16()

	case "sum/32", "sum-32":
		return sum.New32()

	case "sum/64", "sum-64":
		return sum.New64()

	case "bsdsum/8", "bsdsum-8":
		return bsdsum.New8()

	case "bsdsum/16", "bsdsum-16":
		return bsdsum.New16()

	case "bsdsum/32", "bsdsum-32":
		return bsdsum.New32()

	case "bsdsum/64", "bsdsum-64":
		return bsdsum.New64()

	case "sysv":
		return sysvsum.New()

	case "fletcher/16", "fletcher-16":
		return fletcher.New16()
	}

	return nil
}

func AsHex(data []byte) string {
	return hex.EncodeToString(data)
}

func Hash(hash string, data []byte) string {
	return AsHex(HashRaw(hash, data))
}

func HashRaw(hash string, data []byte) []byte {
	h := getHash(hash)

	if h == nil {
		return nil
	}

	h.Write(data)
	return h.Sum(nil)
}

func HashRawStr(hash string, data string) []byte {
	return HashRaw(hash, []byte(data))
}

func HashStr(hash string, data string) string {
	return AsHex(HashRawStr(hash, data))
}

var ErrUnknownHash error = errors.New("Unknown hash!")

func HashReader(hash string, r io.Reader) ([]byte, error) {
	h := getHash(hash)

	if h == nil {
		return nil, ErrUnknownHash
	}

	_, err := io.Copy(h, r)

	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func HashFile(hash string, path string) (string, error) {
	h, err := HashRawFile(hash, path)

	if err != nil {
		return "", err
	}

	return AsHex(h), nil
}

func HashRawFile(hash string, path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	return HashReader(hash, f)
}
