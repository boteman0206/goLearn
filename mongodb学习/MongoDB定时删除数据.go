package mongodb学习

/**

一： mongo删除范围数据
	mongo删除一个范围的数据有点麻烦的，不能直接用remove方法，只能先查出来满足条件的文档，再逐个删除。如果需要删除某个范围，特别是时间范围的数据，有两种解决办法：
	1：使用mongo TTL(Time To Live)。
	2：写个脚本，用crontab写个定时任务，定时连接mongo查询满足删除条件的数据，逐条删除。



二：TTL 索引
	TTL 索引是 MongoDB 中一种特殊的单字段索引，其中字段类型必须是 ISODate 类型或者包含有 ISODate类型的数组。
	1： 创建 TTL 索引和创建普通索引的方法一样，只是多加一个属性而已
		db.collection名称.createIndex( { "ISODate类型的字段": 1 }, { expireAfterSeconds: 过期时间，单位秒 }, { background: true }  )
			如果索引字段是数组，并且字段中有多个日期值，则 MongoDB 使用数组中即最早的日期值来计算到期阈值。
			如果文档中的索引字段不是 ISODate 类型或者包含有 ISODate类型的数组，则该文档不会过期。
			如果文档不包含索引字段(错误指定索引字段)，则该文档不会过期。
	2：如果将 expireAfterSeconds 值设为 0，则过期时间由索引字段的时间来决定。通过这个特性你可以在你的集合中加入一个ISODate类型的字段，插入过期时间并指定它为 TTL 索引字段，来间接实现在指定时间过期。
	3： 和其他索引一样，TTL 索引也支持查询优化
	4： 单字段索引，混合索引不支持。

三： TTL 索引机制
	当你给集合中某一个字段建立 TTL 索引后，后台会启一个线程，不断查询（默认 60s 一次）索引的值来判断文档是否过期，如果过期，则删除，
	但是删除动作可能不会立即执行，需要依据 mongo 实例的负载情况，如果负载很高，可能会稍微延后一段时间再删除。



四： 示例
	// 创建一个 TTL 索引
	db.eventlog.createIndex( { "lastModifiedDate": 1 }, { expireAfterSeconds: 600 } ) // 设置index的过期时间
	db.eventlog.insert({"lastModifiedDate":new Date(),"logevent":2, "logmessage":"hello world"}); // 插入一条数据
	db.eventlog.find().pretty(); 10分钟之后失效
*/
