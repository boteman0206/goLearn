package mongodb学习

/**
一： MongoDB 创建数据库
	语法：use test01  如果数据库不存在，则创建数据库，否则切换到指定数据库
		 show dbs  如果你想查看所有数据库，
		 db.test01.insertOne({"name":"jack", "age":19});  们刚创建的数据库 runoob 并不在数据库的列表中， 要显示它，我们需要向 runoob 数据库插入一些数据。
	注意点：
		MongoDB 中默认的数据库为 test，如果你没有创建新的数据库，集合将存放在 test 数据库中。
		在 MongoDB 中，集合只有在内容插入后才会创建! 就是说，创建集合(数据表)后要再插入一个文档(记录)，集合才会真正创建。

二：MongoDB 删除数据库
	语法：db.dropDatabase()  删除当前数据库，默认为 test，你可以使用 db 命令查看当前数据库名。

三： MongoDB 创建集合查看集合
	1：db.createCollection(name, options)
		name: 要创建的集合名称
		options: 可选参数, 指定有关内存大小及索引的选项  options 可以是如下参数：
				1：capped	布尔	（可选）如果为 true，则创建固定集合。固定集合是指有着固定大小的集合，当达到最大值时，它会自动覆盖最早的文档。当该值为 true 时，必须指定 size 参数。
				2：autoIndexId	布尔	3.2 之后不再支持该参数。（可选）如为 true，自动在 _id 字段创建索引。默认为 false。
				3：size	数值	（可选）为固定集合指定一个最大值，即字节数。如果 capped 为 true，也需要指定该字段。
				4：max	数值	（可选）指定固定集合中包含文档的最大数量。
		在插入文档时，MongoDB 首先检查固定集合的 size 字段，然后检查 max 字段。
		示例： db.createCollection("mycol", { capped : true, autoIndexId : true, size :6142800, max : 10000 } )
	2：show collections 或 show tables 命令：



四： 删除集合
	db.xxx.drop()：xxx表示集合的名称， 如果成功删除选定集合，则 drop() 方法返回 true，否则返回 false。


五：MongoDB 插入文档
	db.COLLECTION_NAME.insert(document)或 db.COLLECTION_NAME.save(document)
	save()：如果 _id 主键存在则更新数据，如果不存在就插入数据。该方法新版本中已废弃，可以使用 db.collection.insertOne() 或 db.collection.replaceOne() 来代替。
	insert(): 若插入的数据主键已经存在，则会抛 org.springframework.dao.DuplicateKeyException 异常，提示主键重复，不保存当前数据。

	3.2 版本之后新增了 db.collection.insertOne() 和 db.collection.insertMany()。
		1：db.collection.insertOne(
		   <document>,
		   {
			  writeConcern: <document>
		   })
		2： db.collection.insertMany(
  			 [ <document 1> , <document 2>, ... ],
		   {
			  writeConcern: <document>,
			  ordered: <boolean>
		   })
		参数说明：
			document：要写入的文档。
			writeConcern：写入策略，默认为 1，即要求确认写操作，0 是不要求。
			ordered：指定是否按顺序写入，默认 true，按顺序写入。

六：MongoDB 更新文档
	MongoDB 使用 update() 和 save() 方法来更新集合中的文档。接下来让我们详细来看下两个函数的应用及其区别。
	1：update() 方法
			格式：db.collection.update(
					   <query>,
					   <update>,
					   {
						 upsert: <boolean>,
						 multi: <boolean>,
						 writeConcern: <document>
					   })
			参数说明：
				query : update的查询条件，类似sql update查询内where后面的。
				update : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
				upsert : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
				multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
				writeConcern :可选，抛出异常的级别。
			db.col.update({'title':'MongoDB 教程'},{$set:{'title':'MongoDB'}})  ： 以上语句只会修改第一条发现的文档，如果你要修改多条相同的文档，则需要设置 multi 参数为 true。

	 2： save() 方法通过传入的文档来替换已有文档，_id 主键存在就更新，不存在就插入。语法格式如下：
			格式： db.collection.save(
					   <document>,
					   {
						 writeConcern: <document>
					   })
			参数说明：
				document : 文档数据。
				writeConcern :可选，抛出异常的级别。



七; MongoDB 删除文档
	MongoDB remove() 函数是用来移除集合中的数据。
	MongoDB 数据更新可以使用 update() 函数。在执行 remove() 函数前先执行 find() 命令来判断执行的条件是否正确，这是一个比较好的习惯。
	1: remove() 方法的基本语法格式如下所示：
			 MongoDB 是 2.6 版本以后的 : db.collection.remove(
				   <query>,
				   {
					 justOne: <boolean>,
					 writeConcern: <document>
				   })
			 参数说明：
				query :（可选）删除的文档的条件。
				justOne : （可选）如果设为 true 或 1，则只删除一个文档，如果不设置该参数，或使用默认值 false，则删除所有匹配条件的文档。
				writeConcern :（可选）抛出异常的级别。
	2: 如果你想删除所有数据，可以使用以下方式（类似常规 SQL 的 truncate 命令）：
		db.col.remove({})


八：MongoDB 查询文档
	1：find() 方法以非结构化的方式来显示所有文档
		语法： db.collection.find(query, projection)
		参数：
			query ：可选，使用查询操作符指定查询条件
			projection ：可选，使用投影操作符指定返回的键。查询时返回文档中所有键值， 只需省略该参数即可（默认省略）。

	2： 如果你需要以易读的方式来读取数据，可以使用 pretty() 方法，语法格式如下：
		db.col.find().pretty()
	3： 除了 find() 方法之外，还有一个 findOne() 方法，它只返回一个文档。

	todo 查询条件参考where语句比较.png
	4：MongoDB AND 条件
		MongoDB 的 find() 方法可以传入多个键(key)，每个键(key)以逗号隔开，即常规 SQL 的 AND 条件。
		db.col.find({key1:value1, key2:value2}).pretty()  以上实例中类似于 WHERE 语句：WHERE key1='value1' AND key2='value2'
	5; MongoDB OR 条件
		MongoDB OR 条件语句使用了关键字 $or,语法格式如下
			db.col.find({
					  $or: [{key1: value1}, {key2:value2}]
				   }).pretty()
	6； AND 和 OR 联合使用
		db.col.find({"likes": {$gt:50}, $or: [{"by": "菜鸟教程"},{"title": "MongoDB 教程"}]}).pretty()
	    以下实例演示了 AND 和 OR 联合使用，类似常规 SQL 语句为： 'where likes>50 AND (by = '菜鸟教程' OR title = 'MongoDB 教程')'
九：MongoDB 条件操作符
	1： (>) 大于 - $gt
		(<) 小于 - $lt
		(>=) 大于等于 - $gte
		(<= ) 小于等于 - $lte
		db.test01.find({"age":{$gte:19}}).pretty(); 查找age>=19的
	2：MongoDB 使用 (<) 和 (>) 查询 - $lt 和 $gt  db.col.find({likes : {$lt :200, $gt : 100}})


十： MongoDB $type 操作符
	$type操作符是基于BSON类型来检索集合中匹配的数据类型，并返回结果。 MongoDB 中可以使用的类型如下表所示 参考图 $type.png
	1： MongoDB 操作符 - $type 实例
		如果想获取 "col" 集合中 title 为 String 的数据，你可以使用以下命令：
		db.col.find({"title" : {$type : 2}}) 或者 db.col.find({"title" : {$type : 'string'}})


十一： MongoDB Limit与Skip方法
	1：Limit()方法
		如果你需要在MongoDB中读取指定数量的数据记录，可以使用MongoDB的Limit方法，limit()方法接受一个数字参数，该参数指定从MongoDB中读取的记录条数。
			db.COLLECTION_NAME.find().limit(NUMBER)
			db.col.find({},{"title":1,_id:0}).limit(2)： 以下实例为显示查询文档中的两条记录： 注：如果你们没有指定limit()方法中的参数则显示集合中的所有数据。

	2：MongoDB Skip() 方法
		我们除了可以使用limit()方法来读取指定数量的数据外，还可以使用skip()方法来跳过指定数量的数据，skip方法同样接受一个数字参数作为跳过的记录条数。
		skip() 方法脚本语法格式如下：
			db.COLLECTION_NAME.find().limit(NUMBER).skip(NUMBER)  注:skip()方法默认参数为 0

十二： MongoDB sort() 方法
	在 MongoDB 中使用 sort() 方法对数据进行排序，sort() 方法可以通过参数指定排序的字段，并使用 1 和 -1 来指定排序的方式，其中 1 为升序排列，而 -1 是用于降序排列。
	sort()方法基本语法如下所示：
		db.COLLECTION_NAME.find().sort({KEY:1})
		以下实例演示了 col 集合中的数据按字段 likes 的降序排列： db.col.find({},{"title":1,_id:0}).sort({"likes":-1})

	todo skip(), limilt(), sort()三个放在一起执行的时候，执行的顺序是先 sort(), 然后是 skip()，最后是显示的 limit()。


十三： MongoDB 聚合
	1; aggregate() 方法  db.COLLECTION_NAME.aggregate(AGGREGATE_OPERATION)
		db.test01.aggregate([{$group : {_id : "$age", num_tutorial : {$sum : 1}}}]);   select age, count(*) from mycol group by age
	常用聚合函数参考图: 聚合函数.png

	2： 聚合管道的概念
		MongoDB的聚合管道将MongoDB文档在一个管道处理完毕后将结果传递给下一个管道处理。管道操作是可以重复的。
		表达式：处理输入文档并输出。表达式是无状态的，只能用于计算当前聚合管道的文档，不能处理其它的文档。这里我们介绍一下聚合框架中常用的几个操作：
			1：$project：修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。
				示例：db.article.aggregate({ $project : {
										title : 1 ,
										author : 1 ,}});
					这样的话结果中就只还有_id,tilte和author三个字段了，默认情况下_id字段是被包含的，如果要想不包含_id话可以这样:   _id : 0 ,

			2：$match：用于过滤数据，只输出符合条件的文档。$match使用MongoDB的标准查询操作。
				db.articles.aggregate( [
                        { $match : { score : { $gt : 70, $lte : 90 } } },
                        { $group: { _id: null, count: { $sum: 1 } } }
                       ] );
				$match用于获取分数大于70小于或等于90记录，然后将符合条件的记录送到下一阶段$group管道操作符进行处理。
			3：$limit：用来限制MongoDB聚合管道返回的文档数。
			4：$skip：在聚合管道中跳过指定数量的文档，并返回余下的文档。
			5：$unwind：将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。
			6：$group：将集合中的文档分组，可用于统计结果。
			7：$sort：将输入文档排序后输出。
			8：$geoNear：输出接近某一地理位置的有序文档。



*/
