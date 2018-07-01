package leveldb

import (
	"fmt"
	"testing"
)

func TestNewDefaultIterator(t *testing.T) {
	data := make(map[string][]byte)

	data["k1"] = []byte("v1")
	data["k5"] = []byte("v5")
	data["k3"] = []byte("v3")
	data["k4"] = []byte("v4")
	data["k9"] = []byte("v9")

	iter := NewDefaultIterator(data)

	if iter.length != 5 {
		t.Fatal()
	}

	for iter.Next() {
		fmt.Printf("%s : %s\n", iter.Key(), string(iter.Value()))
	}
}
