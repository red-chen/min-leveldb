package leveldb

import "fmt"

type Options struct {
}

type DB struct {
	path    string
	options *Options
	data    map[string][]byte
}

func New(path string, options *Options) (*DB, error) {
	self := DB{
		path:    path,
		options: options,
		data:    make(map[string][]byte),
	}
	return &self, nil
}

func (self *DB) Close() error {
	return nil
}

func (self *DB) Put(key []byte, value []byte) error {
	self.data[string(key)] = value
	return nil
}

func (self *DB) Get(key []byte) ([]byte, error) {
	if v, ok := self.data[string(key)]; ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("NotFound")
	}
}

func (self *DB) Del(key []byte) error {
	if _, ok := self.data[string(key)]; ok {
		delete(self.data, string(key))
		return nil
	} else {
		return fmt.Errorf("NotFound")
	}
}

func (self *DB) Iterator() Iterator {
	return NewDefaultIterator(self.data)
}
