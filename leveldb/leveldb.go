package leveldb

import (
	"errors"
)

// 数据相关的选项
type Options struct {
	CreateIfMissing bool
}

type DB struct {
	Path    string
	Options *Options
	Data    map[string][]byte
}

// 创建一个DB实例
func New(path string, options *Options) (*DB, error) {
	self := DB{}
	self.Path = path
	self.Options = options
	self.Data = make(map[string][]byte)
	return &self, nil
}

// 关闭所有内部资源
func (self *DB) Close() error {
	self.Data = nil
	return nil
}

func (self *DB) Put(key []byte, value []byte) error {
	self.Data[string(key)] = value
	return nil
}

func (self *DB) Get(key []byte) ([]byte, error) {
	if v, ok := self.Data[string(key)]; ok {
		return v, nil
	}
	return nil, errors.New("NotFound")
}

func (self *DB) Del(key []byte) error {
	if _, ok := self.Data[string(key)]; ok {
		delete(self.Data, string(key))
		return nil
	}
	return errors.New("NotFound")
}

func (self *DB) Iterator() Iterator {
	iter := NewMapIterator(&self.Data)
	return Iterator(iter)
}
