package redis相关

/**
	https://www.redis.net.cn/tutorial/3504.html
	1：类似的如果安装的时候找不到安装的路径可以使用命令 CONFIG GET * 获取所有的redis配置文件

	2：Redis支持五种数据类型：string（字符串），hash（哈希），list（列表），set（集合）及zset(sorted set：有序集合)。
		2.1： string 注意：一个键最大能存储512MB。常用命令 get set



		2.2: Hash（哈希） Redis hash是一个string类型的field和value的映射表，hash特别适合用于存储对象。
			HMSET runoobkey name "redis tutorial" description "redis basic commands for caching" likes 20 visitors 23000 //同时设置多个 field-value (域-值)对设置到哈希表 key 中。
			HSET key field value // 设置一个值
			HGETALL runoobkey  // 获取所有field
			hget key field // 获取一个field

		2.3： List（列表）Redis 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素导列表的头部（左边）或者尾部（右边）。
			lpush redis.net.cn redis
  		    lrange redis.net.cn 0 10

		2.4： Set（集合） Redis的Set是string类型的无序集合。集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是O(1)。 集合内元素的唯一性，第二次插入的元素将被忽略。
			sadd key member  ：sadd  redis.net redis
			smembers redis.net.cn  扫描key

		2.5：zset(sorted set：有序集合)： Redis zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。zset的成员是唯一的,但分数(score)却可以重复。
			zadd key score member(添加元素到集合，元素在集合中存在则更新对应score) ： zadd redis.net.cn 0 redis,  zadd redis.net.cn 0 mongodb
            ZRANGEBYSCORE redis.net.cn 0 1000

		2.6： Redis 在 2.8.9 版本添加了 HyperLogLog 结构。

	3: 比较常用的mysql的命令： https://www.redis.net.cn/tutorial/3507.html
		3.1： Redis 键相关的基本命令：
			 3.1.1： TTL KEY_NAME：Redis TTL 命令以秒为单位返回 key 的剩余过期时间
				当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以毫秒为单位，返回 key 的剩余生存时间。
			 3.1.2： DEL key 该命令用于在 key 存在是删除 key。
             3.1.3： EXISTS key 检查给定 key 是否存在。 若 key 存在返回 1 ，否则返回 0 。
             3.1.4： EXPIRE key seconds 为给定 key 设置过期时间。
             3.1.5： TYPE key 返回 key 所储存的值的类型。
             3.1.6: SELECT index 切换到指定的数据库
             系统命令：
				1: 获取 redis 服务器的统计信息： info 可以显示redis的详细信息
				2: 获取redis的服务器时间：time 第一个字符串是当前时间(以 UNIX 时间戳格式表示)，而第二个字符串是当前这一秒钟已经逝去的微秒
				3: 返回当前数据库的key数量：dbsize
				4: CONFIG GET parameter 获取指定配置参数的值  示例：config get port  CONFIG GET dir
				5: FLUSHALL 删除所有数据库的所有key
				6: FLUSHDB 删除当前数据库的所有key

		3.2： 字符常用命令：
			3.2.1：GETRANGE key start end 返回 key 中字符串值的子字符 -- 字符串截取
			3.2.2：	SETEX key seconds value 将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)。
			3.2.3： 	SETNX key value 只有在 key 不存在时设置 key 的值。
			3.2.4： INCR key 将 key 中储存的数字值增一。
			3.2.5： INCRBY key increment 将 key 所储存的值加上给定的增量值（increment） 。
			3.2.6： DECR key 将 key 中储存的数字值减一。
			3,。2.7：DECRBY key decrement key 所储存的值减去给定的减量值（decrement） 。
		3.3： hash常用命令
			3.3.1： HDEL key field2 [field2] 删除一个或多个哈希表字段
			3.3.2： 	HEXISTS key field 查看哈希表 key 中，指定的字段是否存在。
			3.3.3：HSCAN key cursor [MATCH pattern] [COUNT count] 迭代哈希表中的键值对。
			3.3.4：HKEYS key 获取所有哈希表中的字段
			3.3.5：HVALS key 获取哈希表中所有值
		3.4：list的常用命令
			3.4.1：	LPUSH key value1 [value2] 将一个或多个值插入到列表头部
			3.4.2：LPOP key 移出并获取列表的第一个元素
			3.4.3：LSET key index value 通过索引设置列表元素的值
			3.4.4：BLPOP key1 [key2 ] timeout 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
		3.5：set常用命令
			3.5.1：SADD key member1 [member2] 向集合添加一个或多个成员
			3.5.2：SCARD key 获取集合的成员数
			3.5.3： 	SDIFF key1 [key2] 返回给定所有集合的差集
			3.5.4：SINTER key1 [key2] 返回给定所有集合的交集
			3.5.5：SUNION key1 [key2] 返回所有给定集合的并集
			3.5.6：SMEMBERS key 返回集合中的所有成员
			3.5.7：	SSCAN key cursor [MATCH pattern] [COUNT count] 迭代集合中的元素   scan 176 MATCH *11* COUNT 1000

		3.6： zset常用命令
			3.6.1：	ZADD key score1 member1 [score2 member2] 向有序集合添加一个或多个成员，或者更新已存在成员的分数
			3.6.2： ZCARD key 获取有序集合的成员数
			3.6.3： ZRANK key member 返回有序集合中指定成员的索引
			3.6.4： ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT] 通过分数返回有序集合指定区间内的成员
			3.6.5： ZSCAN key cursor [MATCH pattern] [COUNT count] 迭代有序集合中的元素（包括元素成员和元素分值）

		3.7：scan命令 scan 176 MATCH *11* COUNT 1000
			SCAN 命令、 SSCAN 命令、 HSCAN 命令和 ZSCAN 命令都返回一个包含两个元素的 multi-bulk 回复： 回复的第一个元素是字符串表示的无符号 64 位整数（游标）， 回复的第二个元素是另一个 multi-bulk 回复， 这个 multi-bulk 回复包含了本次被迭代的元素。
			SCAN 命令返回的每个元素都是一个数据库键。
			SSCAN 命令返回的每个元素都是一个集合成员。
			HSCAN 命令返回的每个元素都是一个键值对，一个键值对由一个键和一个值组成。
			ZSCAN 命令返回的每个元素都是一个有序集合元素，一个有序集合元素由一个成员（member）和一个分值（score）组成。
*/
