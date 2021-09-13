package redis相关

/**
1:Redis 发布订阅
	1.1：订阅消息： SUBSCRIBE redisChat

	1.2：发布消息：重开redis-cli ： PUBLISH redisChat "Redis is a great caching technique"

	1.3：命令详解：
		1.3.1：	PSUBSCRIBE pattern [pattern ...] 订阅一个或多个符合给定模式的频道。
		1.3.2： PUBSUB subcommand [argument [argument ...]] 查看订阅与发布系统状态。 ：PUBSUB  CHANNELS  查看发布的管道信息
		1.3.3：	SUBSCRIBE channel [channel ...] 订阅给定的一个或多个频道的信息。

2：redis的事务
	2.1: 关系型数据中的事务都是原子性的，而redis 的事务是非原子性的
	2.2: Redis事务从你输入MULTI命令开始，这个命令总是答复OK。从这个时候开始用户可以提出多种命令，Redis将会把他们入队，而不是执行这些命令。所有的命令只有在EXEC调用之后才会执行。
	2.3: WATCH命令为Redis事务提供了CAS行为
	2.4: redis为什么不支持事务？
		2.4.1：只有当被调用的Redis命令有语法错误时，这条命令才会执行失败（在将这个命令放入事务队列期间，Redis能够发现此类问题），
				或者对某个键执行不符合其数据类型的操作：实际上，这就意味着只有程序错误才会导致Redis命令执行失败，这种错误很有可能在程序开发期间发现，一般很少在生产环境发现。
		2.4.1：Redis已经在系统内部进行功能简化，这样可以确保更快的运行速度，因为Redis不需要事务回滚的能力

3: redis的lua脚本
	Redis 脚本使用 Lua 解释器来执行脚本。 Reids 2.6 版本通过内嵌支持 Lua 环境。执行脚本的常用命令为 EVAL。

*/
