# 从零开始构建LevelDB - 定义数据库接口，明确需求


参考LevelDB的接口，接口非常简单，https://github.com/google/leveldb/blob/master/include/leveldb/db.h，实现最基本数据读写

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