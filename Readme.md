# 从零开始构建LevelDB

[![Build Status](https://travis-ci.org/red-chen/min-leveldb.png?branch=master)](https://travis-ci.org/red-chen/min-leveldb)

### 背景

对于云计算从业者，Google的几篇论文，比如GFS、MapReduce、BigTable一定耳熟能详，这几篇文章也为当前分布式大数据指明基本方向，
比如后来如著名的Apach Hadoop等。与此同时，Google也放出了LevelDB，实现思路和BigTable基本一致。
阅读LevelDB的源码是一个最好的学习途径，网上有非常多的文章介绍LevelDB，而且非常全面。
但是里面涉及的内容过多，对于初学者很难记住和消化。大部分文章只介绍实现原理和实现细节，缺乏整个演进路线的介绍。

LevelDB中并没有比较特别的新技术，只是为保证业务的需要，需要实现高可用、高并发，是一种工程角度的困难。因此，此项目计划从一个基本的数据存取，逐步演变成一个功能完整的LevelDB，
在演进的过程中，逐步加入LevelDB中所提到的技术。目标是展现出一个DB的诞生过程和LevelDB中的使用到的技术。PS:不追求LevelDB的性能和可用性。

本项目选用Go作为开发语言

### 目录
- [1. 定义数据库接口，明确需求](/doc/chp1.md)
- [2. 用内存的方式，实现读写](/doc/chp2.md)
