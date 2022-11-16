package 图解redis系列

/*
str相关
INCR number  将 key 中储存的数字值增一
INCRBY number 10  将key中存储的数字值加 10
DECR number  将 key 中储存的数字值减一
DECRBY number 10  将key中存储的数字值键 10
STRLEN name  返回 key 所储存的字符串值的长度
EXISTS name  # 判断某个 key 是否存在
EXPIRE name  60  设置 key 在 60 秒后过期（该方法是针对已经存在的key设置过期时间）
TTL name 查看数据还有多久过期

SET key  value EX 60  设置 key-value 类型的值，并设置该key的过期时间为 60 秒
SETEX key  60 value

SETNX key value  不存在就插入：存在的话不执行任何操作


list相关：
LLEN LIST01 获取list01的长度
LPUSH key value [value ...]  # 将一个或多个值value插入到key列表的表头(最左边)，最后的值在最前面
RPUSH key value [value ...] # 将一个或多个值value插入到key列表的表尾(最右边)
LPOP key  # 移除并返回key列表的头元素
RPOP key  # 移除并返回key列表的尾元素
LRANGE key start stop    # 返回列表key中指定区间内的元素，区间以偏移量start和stop指定，从0开始
BLPOP key [key ...] timeout   # 从key列表表头弹出一个元素，没有就阻塞timeout秒，如果timeout=0则一直阻塞
BRPOP key [key ...] timeout   # 从key列表表尾弹出一个元素，没有就阻塞timeout秒，如果timeout=0则一直阻塞
所以，即使没有新消息写入List，消费者也要不停地调用 RPOP 命令，这就会导致消费者程序的 CPU 一直消耗在执行 RPOP 命令上，带来不必要的性能损失。
为了解决这个问题，Redis提供了 BRPOP 命令。BRPOP命令也称为阻塞式读取，客户端在没有读到队列数据时，自动阻塞，直到有新的数据写入队列，再开始读取新数据。和消费者程序自己不停地调用RPOP命令相比，这种方式能节省CPU开销。

BRPOPLPUSH LIST1 new_list TIMEOUT  // 会将list1中的弹出元素保存在new_list中

为了留存消息，List 类型提供了 BRPOPLPUSH 命令，这个命令的作用是让消费者程序从一个 List 中读取消息，同时，Redis 会把这个消息再插入到另一个 List（可以叫作备份 List）留存。
好了，到这里可以知道基于 List 类型的消息队列，满足消息队列的三大需求（消息保序、处理重复的消息和保证消息可靠性）。
	消息保序：使用 LPUSH + RPOP；
	阻塞读取：使用 BRPOP；
	重复消息处理：生产者自行实现全局唯一 ID；
	消息的可靠性：使用 BRPOPLPUSH
List 作为消息队列有什么缺陷？  List 不支持多个消费者消费同一条消息，因为一旦消费者拉取一条消息后，这条消息就从 List 中删除了，无法被其它消费者再次消费。
解决： 这就要说起 Redis 从 5.0 版本开始提供的 Stream 数据类型了，Stream 同样能够满足消息队列的三大需求，而且它还支持「消费组」形式的消息读取



Hash
HSET key field value  # 存储一个哈希表key的键值
HGET key field  # 获取哈希表key对应的field键值
HMSET key field value [field value...]   # 在一个哈希表key中存储多个键值对
HMGET key field [field ...]   # 批量获取哈希表key中多个field键值
HDEL key field [field ...]  # 删除哈希表key中的field键值
HLEN key  # 返回哈希表key中field的数量
HINCRBY key field n   # 为哈希表key中field键的值加上增量n
HGETALL key  # 返回哈希表key中所有的键值


Set
Set 类型是一个无序并唯一的键值集合，它的存储顺序不会按照插入的先后顺序进行存储。
一个集合最多可以存储 2^32-1 个元素。概念和数学中个的集合基本类似，可以交集，并集，差集等等，所以 Set 类型除了支持集合内的增删改查，同时还支持多个集合取交集、并集、差集。
Set 类型和 List 类型的区别如下：
	1：List 可以存储重复元素，Set 只能存储非重复元素；
	2：List 是按照元素的先后顺序存储元素的，而 Set 则是无序方式存储元素的。
SADD key member [member ...] # 往集合key中存入元素，元素存在则忽略，若key不存在则新建
SREM key member [member ...]  # 从集合key中删除元素
SMEMBERS key   获取集合key中所有元素
SCARD key  # 获取集合key中的元素个数
SISMEMBER key member 判断member元素是否存在于集合key中
SRANDMEMBER key [count]  # 从集合key中随机选出count个元素，元素不从key中删除
SPOP key [count] # 从集合key中随机选出count个元素，元素从key中删除
Set运算操作：
SINTER key [key ...]  # 交集运算
SINTERSTORE destination key [key ...]  # 将交集结果存入新集合destination中
SUNION key [key ...]  # 并集运算
SUNIONSTORE destination key [key ...]   将并集结果存入新集合destination中
SDIFF key [key ...]   # 差集运算
SDIFFSTORE destination key [key ...]   # 将差集结果存入新集合destination中
但是要提醒你一下，这里有一个潜在的风险。Set 的差集、并集和交集的计算复杂度较高，在数据量较大的情况下，如果直接执行这些计算，会导致 Redis 实例阻塞。
使用：
	点赞
	共同关注
	抽奖活动


Zset
Zset 类型（有序集合类型）相比于 Set 类型多了一个排序属性 score（分值），对于有序集合 ZSet 来说，每个存储元素相当于有两个值组成的，一个是有序结合的元素值，一个是排序值。
ZADD key score member [[score member]...]  # 往有序集合key中加入带分值元素
ZREM key member [member...]  # 往有序集合key中删除元素
ZSCORE key member   # 返回有序集合key中元素member的分值
ZCARD key  # 返回有序集合key中元素个数
ZINCRBY key increment member # 为有序集合key中元素member的分值加上increment
ZRANGE key start stop [WITHSCORES]  # 正序获取有序集合key从start下标到stop下标的元素 带上WITHSCORES参数可以把分数也显示出来
ZREVRANGE key start stop [WITHSCORES]  # 倒序获取有序集合key从start下标到stop下标的元素
ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]  # 返回有序集合中指定分数区间内的成员，分数由低到高排序。
ZRANGEBYLEX key min max [LIMIT offset count]  # 返回指定成员区间内的成员，按字典正序排列, 分数必须相同。
ZREVRANGEBYLEX key max min [LIMIT offset count] # 返回指定成员区间内的成员，按字典倒序排列, 分数必须相同
Zset 运算操作（相比于 Set 类型，ZSet 类型没有支持差集运算）：
# 并集计算(相同元素分值相加)，numberkeys一共多少个key，WEIGHTS每个key对应的分值乘积
ZUNIONSTORE destkey numberkeys key [key...]
# 交集计算(相同元素分值相加)，numberkeys一共多少个key，WEIGHTS每个key对应的分值乘积
ZINTERSTORE destkey numberkeys key [key...]

应用场景： 在面对需要展示最新列表、排行榜等场景时，如果数据更新频繁或者需要分页显示，可以优先考虑使用 Sorted Set。




BitMap
Bitmap，即位图，是一串连续的二进制数组（0和1），可以通过偏移量（offset）定位元素。BitMap通过最小的单位bit来进行0|1的设置，表示某个元素的值或者状态，时间复杂度为O(1)。
由于 bit 是计算机中最小的单位，使用它进行储存将非常节省空间，特别适合一些数据量大且使用二值统计的场景。  Bitmap 本身是用 String 类型作为底层数据结构实现的一种统计二值状态的数据类型。
SETBIT key offset value  # 设置值，其中value只能是 0 和 1
GETBIT key offset   # 获取值
BITCOUNT key start end  # 获取指定范围内值为 1 的个数  # start 和 end 以字节为单位
BITPOS [key] [value] # 返回指定key中第一次出现指定value(0/1)的位置

应用场景
	Bitmap 类型非常适合二值状态统计的场景，这里的二值状态就是指集合元素的取值就只有 0 和 1 两种，在记录海量数据时，Bitmap 能够有效地节省内存空间。
	签到统计
	判断用户登陆态




HyperLogLog
简单来说 HyperLogLog 提供不精确的去重计数。
PFADD key element [element ...]  # 添加指定元素到 HyperLogLog 中
PFCOUNT key [key ...]  # 返回给定 HyperLogLog 的基数估算值。
PFMERGE destkey sourcekey [sourcekey ...]   # 将多个 HyperLogLog 合并为一个 HyperLogLog

百万级网页 UV 计数
PFADD page1:uv user1 user2 user3 user4 user5
这也就意味着，你使用 HyperLogLog 统计的 UV 是 100 万，但实际的 UV 可能是 101 万。虽然误差率不算大，但是，如果你需要精确统计结果的话，最好还是继续用 Set 或 Hash 类型。



GEO
Redis GEO 是 Redis 3.2 版本新增的数据类型，主要用于存储地理位置信息，并对存储的信息进行操作。
GEO 本身并没有设计新的底层数据结构，而是直接使用了 Sorted Set 集合类型。
GEOADD key longitude latitude member [longitude latitude member ...]   # 存储指定的地理空间位置，可以将一个或多个经度(longitude)、纬度(latitude)、位置名称(member)添加到指定的 key 中。
GEOPOS key member [member ...]  # 从给定的 key 里返回所有指定名称(member)的位置（经度和纬度），不存在的返回 nil。
GEODIST key member1 member2 [m|km|ft|mi]  # 返回两个给定位置之间的距离。
GEORADIUS key longitude latitude radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]  # 根据用户给定的经纬度坐标来获取指定范围内的地理位置集合。


应用场景：滴滴叫车
执行下面的这个命令，就可以把 ID 号为 33 的车辆的当前经纬度位置存入 GEO 集合中
GEOADD cars:locations 116.034579 39.030452 33
例如，LBS 应用执行下面的命令时，Redis 会根据输入的用户的经纬度信息（116.054579，39.030452 ），查找以这个经纬度为中心的 5 公里内的车辆信息，并返回给 LBS 应用。
GEORADIUS cars:locations 116.054579 39.030452 5 km ASC COUNT 10



Stream










*/
