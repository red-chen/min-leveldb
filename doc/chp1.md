# Min-leveldb 定义数据库接口，明确需求


```golang

type DB struct {
}

func New(path string, options *Options) (*DB, error) {
	return nil, nil
}

func (self *DB) Close() error {
	return nil
}

func (self *DB) Put(key []byte, value []byte) error {
	return nil
}

func (self *DB) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (self *DB) Del(key []byte) error {
	return nil
}

func (self *DB) Iterator() Iterator {
    return nil
}

```