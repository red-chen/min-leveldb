# 从零开始构建LevelDB - 使用最简单的方式实现数据的读取


### 下载代码

```
git checkout https://github.com/red-chen/min-leveldb.git
cd min-leveldb; git checkout v1.3
```


### 问题
- 这个数据库只是实现内存的存取，那么重启之后，数据就没有了。那数据库怎么实现持久化呢？