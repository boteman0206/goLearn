package mongodb学习

/**

一：MongoDB 数据库引用
	MongoDB 的关系表示多个文档之间在逻辑上的相互联系。 文档间可以通过嵌入和引用来建立联系。
		1: 嵌入方式就是直接在用户表中保存用户的地址结构数据或者其他数据
		2：引用方式就是在user用户中保存用户的地址的文档的objectid值（手动引用），查询的时候查询两次即可

	引用方式两种：
	1：手动引用（Manual References），就是保存id而已，和mysql的关联外键一样
	2： DBRefs
		DBRef的形式 ：{ $ref : , $id : , $db :  }
		三个字段表示的意义为：
			$ref：集合名称
			$id：引用的id
			$db:数据库名称，可选参数

		示例：以下实例中用户数据文档使用了 DBRef, 字段 address：
              {"_id":ObjectId("53402597d852426020000002"),
			   "address": {
			  		 "$ref": "address_home",
			  		 "$id": ObjectId("534009e4d852427820000002"),
                     "$db": "runoob"},
			   "contact": "987654321",
			   "dob": "01-01-1991",
			   "name": "Tom Benzamin"}
			address DBRef 字段指定了引用的地址文档是在 runoob 数据库下的 address_home 集合，id 为 534009e4d852427820000002。
			以下代码中，我们通过指定 $ref 参数（address_home 集合）来查找集合中指定id的用户地址信息：
			>var user = db.users.findOne({"name":"Tom Benzamin"})
			>var dbRef = user.address
			>db[dbRef.$ref].findOne({"_id":(dbRef.$id)})
			在 MongoDB4.0 版本是这样写; >db[dbRef.$ref].findOne({"_id":ObjectId(dbRef.$id)})


二; MongoDB 覆盖索引查询
	1: 所有的查询字段是索引的一部分
	2: 所有的查询返回字段在同一个索引中
	由于所有出现在查询中的字段是索引的一部分， MongoDB 无需在整个数据文档中检索匹配查询条件和返回使用相同索引的查询结果。因为索引存在于RAM中，从索引中获取数据比通过扫描文档读取数据要快得多
	示例：
		我们在 users 集合中创建联合索引，字段为 gender 和 user_name :
		>db.users.createIndex({gender:1,user_name:1})
		>db.users.find({gender:"M"},{user_name:1,_id:0})
		也就是说，对于上述查询，MongoDB的不会去数据库文件中查找。相反，它会从索引中提取数据，这是非常快速的数据查询。
		如果是以下的查询，不能使用覆盖索引查询： 所有索引字段是一个数组，所有索引字段是一个子文档


三： MongoDB 查询分析
	MongoDB 查询分析可以确保我们所建立的索引是否有效，是查询语句性能分析的重要工具。
	MongoDB 查询分析常用函数有：explain() 和 hint()。
	1; 使用 explain()
		explain 操作提供了查询信息，使用索引及查询统计等。有利于我们对索引的优化。
		db.users.find({gender:"M"},{user_name:1,_id:0}).explain()
		参数解释：
			1：indexOnly: 字段为 true ，表示我们使用了索引。
			2：cursor：因为这个查询使用了索引，MongoDB 中索引存储在B树结构中，所以这是也使用了 BtreeCursor 类型的游标。如果没有使用索引，游标的类型是 BasicCursor。这个键还会给出你所使用的索引的名称，你通过这个名称可以查看当前数据库下的system.indexes集合（系统自动创建，由于存储索引信息，这个稍微会提到）来得到索引的详细信息。
			3：n：当前查询返回的文档数量。
			4：nscanned/nscannedObjects：表明当前这次查询一共扫描了集合中多少个文档，我们的目的是，让这个数值和返回文档的数量越接近越好。
			5：millis：当前查询所需时间，毫秒数。
			6：indexBounds：当前查询具体使用的索引。
	2：使用 hint() // 强制走索引
		虽然MongoDB查询优化器一般工作的很不错，但是也可以使用 hint 来强制 MongoDB 使用一个指定的索引。
		这种方法某些情形下会提升性能。 一个有索引的 collection 并且执行一个多字段的查询(一些字段已经索引了)。
		db.users.find({gender:"M"},{user_name:1,_id:0}).hint({gender:1,user_name:1})
		db.users.find({gender:"M"},{user_name:1,_id:0}).hint({gender:1,user_name:1}).explain()


四： MongoDB 原子操作


*/
