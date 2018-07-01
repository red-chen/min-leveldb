package leveldb

import (
	"bytes"
	"fmt"
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_leveldb(t *testing.T) {
	var err error
	var v []byte
	db, err := New("/tmp/leveldb", nil)
	check(err)

	err = db.Put([]byte("k1"), []byte("v1"))
	check(err)

	err = db.Put([]byte("k4"), []byte("v4"))
	check(err)

	err = db.Put([]byte("k2"), []byte("v2"))
	check(err)

	err = db.Put([]byte("k1"), []byte("v11"))
	check(err)

	err = db.Put([]byte("k8"), []byte("v8"))
	check(err)

	v, err = db.Get([]byte("k8"))

	if !bytes.Equal(v, []byte("v8")) {
		t.Fatal()
	}

	v, err = db.Get([]byte("k1"))

	if !bytes.Equal(v, []byte("v11")) {
		t.Fatal()
	}

	err = db.Del([]byte("k1"))
	check(err)

	_, err = db.Get([]byte("k1"))

	if err == nil {
		t.Fatal()
	}

	iter := db.Iterator()

	for iter.Next() {
		fmt.Printf("%s - %s \n", string(iter.Key()), string(iter.Value()))
	}
}
