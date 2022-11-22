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
	mongodb不支持事务，所以，在你的项目中应用时，要注意这点。无论什么设计，都不要要求mongodb保证数据的完整性。
	但是mongodb提供了许多原子操作，比如文档的保存，修改，删除等，都是原子操作。
	所谓原子操作就是要么这个文档保存到Mongodb，要么没有保存到Mongodb，不会出现查询到的文档没有保存完整的情况。

	原子操作常用命令
		$set 用来指定一个键并更新键值，若键不存在并创建。  { $set : { field : value } }
		$unset 用来删除一个键。 { $unset : { field : 1} }
		$inc $inc可以对文档的某个值为数字型（只能为满足要求的数字）的键进行增减的操作。  { $inc : { field : value } }
		$push { $push : { field : value } } 把value追加到field里面去，field一定要是数组类型才行，如果field不存在，会新增一个数组类型加进去。
		$pushAll  同$push,只是一次可以追加多个值到一个数组字段内。  { $pushAll : { field : value_array } }
		$pull 从数组field内删除一个等于value值。 { $pull : { field : _value } }
		$addToSet 增加一个值到数组内，而且只有当这个值不在数组内才增加。
		$pop 删除数组的第一个或最后一个元素 { $pop : { field : 1 } }
		$rename 修改字段名称 { $rename : { old_field_name : new_field_name } }
		$bit 位操作，integer类型 {$bit : { field : {and : 5}}}


五： MongoDB 高级索引
	1： 索引数组字段
		假设我们基于标签来检索用户，为此我们需要对集合中的数组 tags 建立索引，组中创建索引，需要对数组中的每个字段依次建立索引。所以在我们为数组 tags 创建索引时，
		会为 music、cricket、blogs三个值建立单独的索引。

	2： 索引子文档字段
		假设我们需要通过city、state、pincode字段来检索文档，由于这些字段是子文档的字段，所以我们需要对子文档建立索引。为子文档的三个字段创建索引，命令如下：
		db.users.ensureIndex({"address.city":1,"address.state":1,"address.pincode":1})

六： MongoDB 索引限制
	每个索引占据一定的存储空间，在进行插入，更新和删除操作时也需要对索引进行操作。所以，如果你很少对集合进行读取操作，建议不使用索引。
	由于索引是存储在内存(RAM)中,你应该确保该索引的大小不超过内存的限制。如果索引的大小大于内存的限制，MongoDB会删除一些索引，这将导致性能下降。
	1：查询限制
		索引不能被以下的查询使用：
			正则表达式及非操作符，如 $nin, $not, 等。
			算术运算符，如 $mod, 等。
			$where 子句
	2：索引键限制
		从2.6版本开始，如果现有的索引字段的值超过索引键的限制，MongoDB中不会创建索引。
	3：插入文档超过索引键限制
		如果文档的索引字段值超过了索引键的限制，MongoDB不会将任何文档转换成索引的集合。与mongorestore和mongoimport工具类似。
	4：最大范围
		集合中索引不能超过64个
		索引名的长度不能超过128个字符
		一个复合索引最多可以有31个字段

七：MongoDB ObjectId
	ObjectId 是一个12字节 BSON 类型数据，有以下格式：
		前4个字节表示时间戳
		接下来的3个字节是机器标识码
		紧接的两个字节由进程id组成（PID）
		最后三个字节是随机数。
	MongoDB中存储的文档必须有一个"_id"键。这个键的值可以是任何类型的，默认是个ObjectId对象。
	在一个集合里面，每个文档都有唯一的"_id"值，来确保集合里面每个文档都能被唯一标识。
	MongoDB采用ObjectId，而不是其他比较常规的做法（比如自动增加的主键）的主要原因，因为在多个 服务器上同步自动增加主键值既费力还费时。
	创建id： newObjectId = ObjectId()
	ObjectId 转换为字符串： new ObjectId().str


八： MongoDB Map Reduce 参考图片: mapreduce流程图.png
	Map-Reduce是一种计算模型，简单的说就是将大批量的工作（数据）分解（MAP）执行，然后再将结果合并成最终结果（REDUCE）。
	MongoDB提供的Map-Reduce非常灵活，对于大规模数据分析也相当实用。
	1：以下是MapReduce的基本语法：
		>db.collection.mapReduce(
			   function() {emit(key,value);},  //map 函数
			   function(key,values) {return reduceFunction},   //reduce 函数
			   {
				  out: collection,
				  query: document,
				  sort: document,
				  limit: number
			   })
	使用 MapReduce 要实现两个函数 Map 函数和 Reduce 函数,Map 函数调用 emit(key, value), 遍历 collection 中所有的记录, 将 key 与 value 传递给 Reduce 函数进行处理。
	Map 函数必须调用 emit(key, value) 返回键值对。
	参数说明:
		map ：映射函数 (生成键值对序列,作为 reduce 函数参数)。
		reduce 统计函数，reduce函数的任务就是将key-values变成key-value，也就是把values数组变成一个单一的值value。。
		out 统计结果存放集合 (不指定则使用临时集合,在客户端断开后自动删除)。
		query 一个筛选条件，只有满足条件的文档才会调用map函数。（query。limit，sort可以随意组合）
		sort 和limit结合的sort排序参数（也是在发往map函数前给文档排序），可以优化分组机制
		limit 发往map函数的文档数量的上限（要是没有limit，单独使用sort的用处不大）

	2: mapReduce 输出结果为:
		{
			"result" : "post_total",
			"timeMillis" : 23,
			"counts" : {
					"input" : 5,
					"emit" : 5,
					"reduce" : 1,
					"output" : 2},
			"ok" : 1
		}
		具体参数说明：
			result：储存结果的collection的名字,这是个临时集合，MapReduce的连接关闭后自动就被删除了。
			timeMillis：执行花费的时间，毫秒为单位
			input：满足条件被发送到map函数的文档个数
			emit：在map函数中emit被调用的次数，也就是所有集合中的数据总量
			output：结果集合中的文档个数（count对调试非常有帮助）
			ok：是否成功，成功为1
			err：如果失败，这里可以有失败原因，不过从经验上来看，原因比较模糊，作用不大


九： MongoDB 全文检索
全文检索对每一个词建立一个索引，指明该词在文章中出现的次数和位置，当用户查询时，检索程序就根据事先建立的索引进行查找，并将查找的结果反馈给用户的检索方式。
这个过程类似于通过字典中的检索字表查字的过程。
	1：启用全文检索 （MongoDB 在 2.6 版本以后是默认开启全文检索的，如果你使用之前的版本，你需要使用以下代码来启用全文检索）
		>db.adminCommand({setParameter:true,textSearchEnabled:true})
		或者： mongod --setParameter textSearchEnabled=true
	2：创建全文索引
		考虑以下 posts 集合的文档数据，包含了文章内容（post_text）及标签(tags)：
			{"post_text": "enjoy the mongodb articles on Runoob",
			   "tags": [
				  "mongodb",
				  "runoob"
			   ]}
		建立全文索引： db.posts.ensureIndex({post_text:"text"})
	3： 使用全文索引 ： 现在我们已经对 post_text 建立了全文索引，我们可以搜索文章中的关键词 runoob：
		>db.posts.find({$text:{$search:"runoob"}})
		>db.posts.runCommand("text",{search:"runoob"}) // 旧版本
	4：删除全文索引
		>db.posts.getIndexes()  通过以上命令获取索引名，本例的索引名为post_text_text，执行以下命令来删除索引：
		>db.posts.dropIndex("post_text_text")


十：MongoDB 正则表达式
	MongoDB 使用 $regex 操作符来设置匹配字符串的正则表达式。MongoDB使用PCRE (Perl Compatible Regular Expression) 作为正则表达式语言。不同于全文检索，我们使用正则表达式不需要做任何配置。
	考虑以下 posts 集合的文档结构，该文档包含了文章内容和标签：
		{  "post_text": "enjoy the mongodb articles on runoob",
		   "tags": [
			  "mongodb",
			  "runoob"]
		}
	1： 使用正则表达式
		以下命令使用正则表达式查找包含 runoob 字符串的文章：
		>db.posts.find({post_text:{$regex:"runoob"}})
		>db.posts.find({post_text:/runoob/}) // 也可以写成这样
	2：不区分大小写的正则表达式
		如果检索需要不区分大小写，我们可以设置 $options 为 $i。
			>db.posts.find({post_text:{$regex:"runoob",$options:"$i"}})
	3：数组元素使用正则表达式
		我们还可以在数组字段中使用正则表达式来查找内容。 这在标签的实现上非常有用，如果你需要查找包含以 run 开头的标签数据(ru 或 run 或 runoob)， 你可以使用以下代码：
		>db.posts.find({tags:{$regex:"run"}})
	4： 优化正则表达式查询
		如果你的文档中字段设置了索引，那么使用索引相比于正则表达式匹配查找所有的数据查询速度更快。
		如果正则表达式是前缀表达式，所有匹配的数据将以指定的前缀字符串为开始。例如： 如果正则表达式为 ^tut ，查询语句将查找以 tut 为开头的字符串。

	5：这里面使用正则表达式有两点需要注意：
		正则表达式中使用变量。一定要使用eval将组合的字符串进行转换，不能直接将字符串拼接后传入给表达式。否则没有报错信息，只是结果为空！实例如下：
		var name=eval("/" + 变量值key +"/i");
		title:eval("/"+title+"/i")    // 等同于 title:{$regex:title,$Option:"$i"}

十二：MongoDB 固定集合（Capped Collections）
MongoDB 固定集合（Capped Collections）是性能出色且有着固定大小的集合，对于大小固定，我们可以想象其就像一个环形队列，当集合空间用完后，再插入的元素就会覆盖最初始的头部的元素！
	1：创建固定集合
		我们通过createCollection来创建一个固定集合，且capped选项设置为true：
			>db.createCollection("cappedLogCollection",{capped:true,size:10000})
		还可以指定文档个数,加上max:1000属性：
			>db.createCollection("cappedLogCollection",{capped:true,size:10000,max:1000})
		判断集合是否为固定集合:
			>db.cappedLogCollection.isCapped()
		如果需要将已存在的集合转换为固定集合可以使用以下命令：
			>db.runCommand({"convertToCapped":"posts",size:10000})
	2：固定集合查询
		固定集合文档按照插入顺序储存的,默认情况下查询就是按照插入顺序返回的,也可以使用$natural调整返回顺序。
		>db.cappedLogCollection.find().sort({$natural:-1})
	3: 固定集合的功能特点
		可以插入及更新,但更新不能超出collection的大小,否则更新失败,不允许删除,但是可以调用drop()删除集合中的所有行,但是drop后需要显式地重建集合。
		在32位机子上一个cappped collection的最大值约为482.5M,64位上只受系统文件大小的限制。
	4:固定集合属性及用法
		属性
			属性1:对固定集合进行插入速度极快
			属性2:按照插入顺序的查询输出速度极快
			属性3:能够在插入最新数据时,淘汰最早的数据
		用法
			用法1:储存日志信息
			用法2:缓存一些少量的文档


*/
