package 图解redis系列

/**

一：AOF 日志
	AOF(Append Only File) 持久化功能，注意只会记录写操作命令，读操作命令是不会被记录的，因为没意义。
	在 Redis 中 AOF 持久化功能默认是不开启的，需要我们修改 redis.conf 配置文件中的以下参数
		appendonly   yes  // 标识是否开启aof持久化（默认 no，关闭）
		appendfilename "appendonly.aof"  // aof持久化文件的名称

二：RDB快照模式




三： 混合模式



*/
