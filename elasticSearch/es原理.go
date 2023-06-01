package elasticSearch

/**


文档存储位置：
	所有的索引和文档数据是存储在本地的磁盘中，具体的路径可在 ES 的配置文件../config/elasticsearch.yml中配置，如下：

优化技巧：
	Elasticsearch 重度使用磁盘，你的磁盘能处理的吞吐量越大，你的节点就越稳定。这里有一些优化磁盘 I/O 的技巧：
		使用 SSD。就像其他地方提过的， 他们比机械磁盘优秀多了。
		使用 RAID 0。条带化 RAID 会提高磁盘 I/O，代价显然就是当一块硬盘故障时整个就故障了。不要使用镜像或者奇偶校验 RAID 因为副本已经提供了这个功能。
		另外，使用多块硬盘，并允许 Elasticsearch 通过多个 path.data 目录配置把数据条带化分配到它们上面。
		不要使用远程挂载的存储，比如 NFS 或者 SMB/CIFS。这个引入的延迟对性能来说完全是背道而驰的。
		如果你用的是 EC2，当心 EBS。即便是基于 SSD 的 EBS，通常也比本地实例的存储要慢。


倒排列表：
	Elasticsearch 基于 Lucene，一个 Lucene 索引 我们在 Elasticsearch 称作 分片 ， 并且引入了 按段搜索 的概念。

倒排索引
	倒排索引被写入磁盘后是不可改变 的：它永远不会修改。






*/
