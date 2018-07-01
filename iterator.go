package leveldb

import "fmt"

type Iterator interface {
	Next() bool

	Key() []byte

	Value() []byte

	Close() error
}

type Pair struct {
	Key   []byte
	Value []byte
}

type DefaultIterator struct {
	data   []Pair
	index  int
	length int
}

func NewDefaultIterator(data map[string][]byte) *DefaultIterator {
	self := new(DefaultIterator)
	self.index = -1
	self.length = len(data)
	for k, v := range data {
		p := Pair{
			Key:   []byte(k),
			Value: v,
		}

		self.data = append(self.data, p)
	}
	return self
}

func (self *DefaultIterator) Next() bool {
	if self.index < self.length-1 {
		self.index++
		return true
	}
	return false
}

func (self *DefaultIterator) Key() []byte {
	if self.index == -1 || self.index >= self.length {
		panic(fmt.Errorf("IndexOutOfBoundsError"))
	}

	return self.data[self.index].Key
}

func (self *DefaultIterator) Value() []byte {
	if self.index >= self.length {
		panic(fmt.Errorf("IndexOutOfBoundsError"))
	}

	return self.data[self.index].Value
}

func (self *DefaultIterator) Close() error {
	return nil
}
