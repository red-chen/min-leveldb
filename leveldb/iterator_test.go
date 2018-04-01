package leveldb

import (
	"fmt"
	"testing"
)

func TestIterator_Basic(t *testing.T) {
	data := map[string][]byte{
		"k0": []byte("v0"),
		"k1": []byte("v1"),
		"k2": []byte("v2"),
		"k3": []byte("v3"),
		"k4": []byte("v4"),
		"k5": []byte("v5"),
	}

	iter := NewMapIterator(&data)

	for {
		if k, v, ok := iter.Next(); ok {
			fmt.Printf("%s - %s \n", k, v)
		} else {
			break
		}
	}

	if 6 != iter.Len() {
		t.Error(fmt.Sprintf("Expect length 6, but %d", iter.Len()))
	}
}
