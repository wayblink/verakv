# LSM

# 为什么是LSM？

LSM-tree(Log-Structured Merged-tree) 现如今已经被广泛应用在了各个NoSQL存储系统中，
如BigTable, Dynamo, HBase, Cassandra, LevelDB, RocksDB 和 AsterixDB。
相比于传统的in-place updates(就地更新) 索引结构，LSM-tree 第一次写入都缓存到内存中，
并通过后台的flush来顺序追加到磁盘中，也就是out-of-palce updates。
LSM-tree这样的实现方式有非常多的优点，包括写性能的提升、简单的并发控制和异常恢复等。
这些优点使LSM树可以服务于多种场景，如Facebook开发的基于LSM-tree的存储引擎rocksdb，
被广泛应用在实时数据处理/图数据处理/流式数据处理以及OLTP(on line transaction processing)等多种workload。

*key point*

以日志存储数据

