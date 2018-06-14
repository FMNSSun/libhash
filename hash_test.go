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

func Test(t *testing.T) {
	table := map[string]map[string]string {
		"fletcher/16" : map[string]string {
			"abcde" : "c8f0",
			"abcdefgh": "0627",
		},
	}

	run(t, table)
}
