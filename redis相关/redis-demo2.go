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
	3.1:redis Eval 命令基本语法如下： redis 127.0.0.1:6379> EVAL script numkeys key [key ...] arg [arg ...]
		参数说明：
			script： 参数是一段 Lua 5.1 脚本程序。脚本不必(也不应该)定义为一个 Lua 函数。
			numkeys： 用于指定键名参数的个数。
			key [key ...]： 从 EVAL 的第三个参数开始算起，表示在脚本中所用到的那些 Redis 键(key)，这些键名参数可以在 Lua 中通过全局变量 KEYS 数组，用 1 为基址的形式访问( KEYS[1] ， KEYS[2] ，以此类推)。
			arg [arg ...]： 附加参数，在 Lua 中通过全局变量 ARGV 数组访问，访问的形式和 KEYS 变量类似( ARGV[1] 、 ARGV[2] ，诸如此类)。
		 示例：eval "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}" 2 key1 key2 first second
	3.2： SCRIPT LOAD "return 'hello moto'"  --加载脚本到redis
	3.3： SCRIPT EXISTS sha1  -- 检查脚本是否存在
	3.4： redis Evalsha 也是执行脚本：EVALSHA sha1 numkeys key [key ...] arg [arg ...]  示例：EVALSHA "232fd51614574cf0867b83d384a5e898cfd24e5a" 0 -- 返回上面脚本的执行hello moto

4: redis的save命令
	Redis SAVE 命令用于创建当前数据库的备份。 redis 127.0.0.1:6379> SAVE  该命令将在 redis 安装目录中创建dump.rdb文件。
	Bgsave： 创建 redis 备份文件也可以使用命令 BGSAVE，该命令在后台执行。
5：恢复数据
	如果需要恢复数据，只需将备份文件 (dump.rdb) 移动到 redis 安装目录并启动服务即可。获取 redis 目录可以使用 CONFIG 命令，如下 CONFIG GET dir

6：设置redis的密码：
	我们可以通过 redis 的配置文件设置密码参数，这样客户端连接到 redis 服务就需要密码验证，这样可以让你的 redis 服务更安全。
	CONFIG get requirepass：默认情况下 requirepass 参数是空的，这就意味着你无需通过密码验证就可以连接到 redis 服务。
    CONFIG set requirepass "w3cschool.cc"   设置密码后，客户端连接 redis 服务就需要密码验证，否则无法执行命令。
	AUTH password： AUTH 命令基本语法格式如下：
*/
