package mongodb学习

/**
https://haicoder.net/mongodb/mongodb-find-array-subfield.html
https://www.runoob.com/mongodb/mongodb-capped-collections.html 参考

一：mongodb的概念参考图 mongodb的概念图.png
	1：数据库
		一个mongodb中可以建立多个数据库。MongoDB的默认数据库为"db"，该数据库存储在data目录中。 MongoDB的单个实例可以容纳多个独立的数据库，每一个都有自己的集合和权限，不同的数据库也放置在不同的文件中。
			1：show dbs 命令可以显示所有数据的列表。
			2：use xxx 命令可以连接到指定的数据库
		有一些数据库名是保留的，可以直接访问这些有特殊作用的数据库。
			1：admin： 从权限的角度来看，这是"root"数据库。要是将一个用户添加到这个数据库，这个用户自动继承所有数据库的权限。一些特定的服务器端命令也只能从这个数据库运行，比如列出所有的数据库或者关闭服务器。
			2：local: 这个数据永远不会被复制，可以用来存储限于本地单台服务器的任意集合
			3：config: 当Mongo用于分片设置时，config数据库在内部使用，用于保存分片的相关信息。

	2: 文档(Document)
		文档是一组键值(key-value)对(即 BSON)。MongoDB 的文档不需要设置相同的字段，并且相同的字段不需要相同的数据类型，这与关系型数据库有很大的区别，也是 MongoDB 非常突出的特点。
		需要注意的是：
			1：文档中的键/值对是有序的。
			2：文档中的值不仅可以是在双引号里面的字符串，还可以是其他几种数据类型（甚至可以是整个嵌入的文档)。
			3：MongoDB区分类型和大小写。
			4：MongoDB的文档不能有重复的键。
			5：文档的键是字符串。除了少数例外情况，键可以使用任意UTF-8字符。
		文档键命名规范：
			1：键不能含有\0 (空字符)。这个字符用来表示键的结尾。
			2：.和$有特别的意义，只有在特定环境下才能使用。
			3：以下划线"_"开头的键是保留的(不是严格要求的)
	3：集合
		集合就是 MongoDB 文档组，类似于 RDBMS （关系数据库管理系统：Relational Database Management System)中的表格。
		集合存在于数据库中，集合没有固定的结构，这意味着你在对集合可以插入不同格式和类型的数据，但通常情况下我们插入集合的数据都会有一定的关联性。
		合法的集合名
			集合名不能是空字符串""。
			集合名不能含有\0字符（空字符)，这个字符表示集合名的结尾。
			集合名不能以"system."开头，这是为系统集合保留的前缀。
			用户创建的集合名字不能含有保留字符。有些驱动程序的确支持在集合名里面包含，这是因为某些系统生成的集合中包含该字符。除非你要访问这种系统创建的集合，否则千万不要在名字里出现$。

	4：capped collections
		Capped collections 就是固定大小的collection。 它有很高的性能以及队列过期的特性(过期按照插入的顺序). 有点和 "RRD" 概念类似。
		Capped collections 是高性能自动的维护对象的插入顺序。它非常适合类似记录日志的功能和标准的 collection 不同，你必须要显式的创建一个capped collection，指定一个 collection 的大小，单位是字节。collection 的数据存储空间值提前分配的
		Capped collections 可以按照文档的插入顺序保存到集合中，而且这些文档在磁盘上存放位置也是按照插入顺序来保存的，所以当我们更新Capped collections 中文档的时候，更新后的文档不可以超过之前文档的大小，这样话就可以确保所有文档在磁盘上的位置一直保持不变。
		由于 Capped collection 是按照文档的插入顺序而不是使用索引确定插入位置，这样的话可以提高增添数据的效率。MongoDB 的操作日志文件 oplog.rs 就是利用 Capped Collection 来实现的。
		db.createCollection("mycoll", {capped:true, size:100000})
		注意点：
			1：在 capped collection 中，你能添加新的对象。
			2：能进行更新，然而，对象不会增加存储空间。如果增加，更新就会失败 。
			3：使用 Capped Collection 不能删除一个文档，可以使用 drop() 方法删除 collection 所有的行。
			4：删除之后，你必须显式的重新创建这个 collection。
			5： 在32bit机器中，capped collection 最大存储为 1e9( 1X10^9)个字节。

	5：元数据
		数据库的信息是存储在集合中。它们使用了系统的命名空间：  dbname.system.*
		在MongoDB数据库中名字空间 <dbname>.system.* 是包含多种系统信息的特殊集合(Collection)，如下:
			1：dbname.system.namespaces	列出所有名字空间。
			2：dbname.system.indexes	列出所有索引。
			3：dbname.system.profile	包含数据库概要(profile)信息。
			4：dbname.system.users	列出所有可访问数据库的用户。
			5：dbname.local.sources	包含复制对端（slave）的服务器信息和状态。

二： MongoDB 数据类型 参考入 数据类型.png
	1:ObjectId   ObjectId 类似唯一主键，可以很快的去生成和排序，包含 12 bytes，含义是：
		前 4 个字节表示创建 unix 时间戳,格林尼治时间 UTC 时间，比北京时间晚了 8 个小时
		接下来的 3 个字节是机器标识码
		紧接的两个字节由进程 id 组成 PID
		最后三个字节是随机数
		var newObject = ObjectId()
		获取文档的创建时间: newObject.getTimestamp()
		ObjectId 转为字符串 newObject.str
	2: 字符串  BSON 字符串都是 UTF-8 编码。
	3: 时间戳
		BSON 有一个特殊的时间戳类型用于 MongoDB 内部使用，与普通的 日期 类型不相关。 时间戳值是一个 64 位的值。其中：
			前32位是一个 time_t 值（与Unix新纪元相差的秒数）
			后32位是在某秒中操作的一个递增的序数
		BSON 时间戳类型主要用于 MongoDB 内部使用。在大多数情况下的应用开发中，你可以使用 BSON 日期类型。
	4:日期
		表示当前距离 Unix新纪元（1970年1月1日）的毫秒数。日期类型是有符号的, 负数表示 1970 年之前的日期。



三：MongoDB - 连接
	mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
		mongodb:// 这是固定的格式，必须要指定。
		username:password@ 可选项，如果设置，在连接数据库服务器之后，驱动都会尝试登录这个数据库
		host1 必须的指定至少一个host, host1 是这个URI唯一要填写的。它指定了要连接服务器的地址。如果要连接复制集，请指定多个主机地址。
		portX 可选的指定端口，如果不填，默认为27017
		/database 如果指定username:password@，连接并验证登录指定数据库。若不指定，默认打开 test 数据库。
		?options 是连接选项。如果不使用/database，则前面需要加上/。所有连接选项都是键值对name=value，键值对之间通过&或;（分号）隔开
			具体的option选项参考图  option选项.png


四： MongoDB 索引
	索引通常能够极大的提高查询的效率，如果没有索引，MongoDB在读取数据时必须扫描集合中的每个文件并选取那些符合查询条件的记录。
	这种扫描全集合的查询效率是非常低的，特别在处理大量的数据时，查询可以要花费几十秒甚至几分钟，这对网站的性能是非常致命的。
	索引是特殊的数据结构，索引存储在一个易于遍历读取的数据集合中，索引是对数据库表中一列或多列的值进行排序的一种结构
	1：createIndex() 方法：
		注意在 3.0.0 版本前创建索引方法为 db.collection.ensureIndex()，之后的版本使用了 db.collection.createIndex() 方法，ensureIndex() 还能用，但只是 createIndex() 的别名。
		1：db.collection.createIndex(keys, options)： 语法中 Key 值为你要创建的索引字段，1 为指定按升序创建索引，如果你想按降序来创建索引指定为 -1 即可。
		2：createIndex() 方法中你也可以设置使用多个字段创建索引（关系型数据库中称作复合索引）
			db.col.createIndex({"title":1,"description":-1})
	参数参考createIndex.png图片
		在后台创建索引：db.values.createIndex({open: 1, close: 1}, {background: true})

	2：查看集合索引： db.col.getIndexes()
	3：查看集合索引大小 db.col.totalIndexSize()
	4：删除集合指定索引  db.col.dropIndex("索引名称")


*/
