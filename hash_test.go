package libhash

import "testing"

func checkHash(t *testing.T, hashf, hashs, hashe string) {
	h := HashStr(hashf, hashs)
	if h != hashe {
		t.Errorf("%q: Expected %q but got %q", hashf, hashe, h)
	}
}

func checkHashes(t *testing.T, hashf string, kv map[string]string) {
	for k,v := range kv {
		checkHash(t, hashf, k, v)
	}
}

func run(t *testing.T, kkv map[string]map[string]string) {
	for k, kv := range kkv {
		checkHashes(t, k, kv)
	}
}

const L_STR = "71144850f4fb4cc55fc0ee6935badddf3fcb152acc59cfc7f45ff83b85f5a0bc"
const L_AE = "abcde"
const L_AH = "abcdefgh"
const L_HW = "Hello, world!"

func Test(t *testing.T) {
	table := map[string]map[string]string {
		"fletcher/16" : map[string]string {
			L_AE : "c8f0",
			L_AH : "0627",
		},

		"sha1" : map[string]string {
			L_STR : "650d340ea539598d225fdf54b2c9e850a04476cc",
			L_HW : "943a702d06f34599aee1f8da8ef9f7296031d699",
			L_AE : "03de6c570bfe24bfc328ccd7ca46b76eadaf4334",
			L_AH : "425af12a0743502b322e93a015bcf868e324d56a",
		},

		"md5" : map[string]string {
			L_STR : "4e81fbe27a8c79e59e1f6c1a5624cd58",
			L_HW : "6cd3556deb0da54bca060b4c39479839",
			L_AE : "ab56b4d92b40713acc5af89985d4b786",
			L_AH : "e8dc4081b13434b45189a720b77b6818", 
		},

		"xor8" : map[string]string {
			L_AE : "61",
			L_AH : "08",
		},
	}

	run(t, table)
}
