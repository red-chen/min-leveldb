package leveldb

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLevelDB_Basic(t *testing.T) {
	var path string = "testdb"
	var db *DB
	var err error

	if db, err = New(
		path,
		&Options{CreateIfMissing: true}); err != nil {
		t.Error(err)
	}
	defer db.Close()

	if err = db.Put([]byte("name"), []byte("redchen")); err != nil {
		t.Error(err)
	}
	if err = db.Put([]byte("age"), []byte("30")); err != nil {
		t.Error(err)
	}

	var value []byte
	if value, err = db.Get([]byte("name")); err != nil {
		t.Error(err)
	}
	if !bytes.Equal([]byte("redchen"), value) {
		t.Fail()
	}

	if value, err = db.Get([]byte("age")); err != nil {
		t.Error(err)
	}
	if !bytes.Equal([]byte("30"), value) {
		t.Fail()
	}

	iter := db.Iterator()

	for {
		if k, v, ok := iter.Next(); ok {
			fmt.Printf("%s - %s \n", k, v)
		} else {
			break
		}
	}

	if 2 != iter.Len() {
		t.Error(fmt.Sprintf("Expect length 2, but %d", iter.Len()))
	}

	if value, err = db.Get([]byte("name1")); err != nil {
		if err.Error() != "NotFound" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
