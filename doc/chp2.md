# 从零开始构建LevelDB - 用内存的方式，实现读写
> 代码用最简单的map，实现了Leveldb的核心方法，并提供了Iterator接口，可以遍历数据库中的KV元素。内容比较简单，不用单独在
> 通过文档描述了，大家可以参阅代码。

# 代码获取方式
```
git clone https://github.com/red-chen/min-leveldb.git
git checkout chp2
```

# 问题
1. 考虑一下怎么将数据持久化？

