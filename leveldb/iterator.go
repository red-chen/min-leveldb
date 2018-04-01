package leveldb

type Iterator interface {
	Len() int
	Next() (key []byte, value []byte, ok bool)
}

type MapIterator struct {
	index  int
	length int
	keys   []string
	data   *map[string][]byte
}

func NewMapIterator(data *map[string][]byte) *MapIterator {
	self := MapIterator{}
	self.data = data

	for k, _ := range *self.data {
		self.keys = append(self.keys, k)
	}

	self.length = len(self.keys)
	return &self
}

func (self *MapIterator) Next() (key []byte, value []byte, ok bool) {
	if self.index >= self.length {
		ok = false
		return nil, nil, ok
	}
	key = []byte(self.keys[self.index])
	value, ok = (*self.data)[string(key)], true
	self.index++
	return key, value, ok
}

func (self *MapIterator) Len() int {
	return self.length
}
