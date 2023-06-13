package elasticSearch

/**

索引管理的引入：
	我们在前文中增加文档时，如下的语句会动态创建一个customer的index：
	PUT /customer/_doc/1
	{
	  "name": "John Doe"
	}


删除索引
	最后我们将创建的test-index-users删除。
	DELETE /test-index-users


一： 普通查询：
	GET /bank/_search
	{
	  "query": { "match_all": {} },
	  "sort": [
		{ "account_number": "asc" }
	  ]
	}
	结果集合：
		{
		  "took" : 29,
		  "timed_out" : false,
		  "_shards" : {
			"total" : 1,
			"successful" : 1,
			"skipped" : 0,
			"failed" : 0
		  },
		  "hits" : {
			"total" : {
			  "value" : 1000,
			  "relation" : "eq"
			},
			"max_score" : null,
			"hits" : [
			  {
				"_index" : "bank",
				"_type" : "_doc",
				"_id" : "0",
				"_score" : null,
				"_source" : {
				  "account_number" : 0,
				  "balance" : 16623,
				  "firstname" : "Bradshaw",
				  "lastname" : "Mckenzie",
				  "age" : 29,
				  "gender" : "F",
				  "address" : "244 Columbus Place",
				  "employer" : "Euron",
				  "email" : "bradshawmckenzie@euron.com",
				  "city" : "Hobucken",
				  "state" : "CO"
				},
				"sort" : [
				  0
				]
			  },

		结果字段展示：
			took – Elasticsearch运行查询所花费的时间（以毫秒为单位）
			timed_out –搜索请求是否超时
			_shards - 搜索了多少个碎片，以及成功，失败或跳过了多少个碎片的细目分类。
			max_score – 找到的最相关文档的分数
			hits.total.value - 找到了多少个匹配的文档
			hits.sort - 文档的排序位置（不按相关性得分排序时）
			hits._score - 文档的相关性得分（使用match_all时不适用）#


二： 分页查询
	GET /bank/_search
	{
	  "query": { "match_all": {} },
	  "sort": [
		{ "account_number": "asc" }
	  ],
	  "from": 10,
	  "size": 10
	}

三：指定字段查询：match
	如果要在字段中搜索特定字词，可以使用match; 如下语句将查询address 字段中包含 mill 或者 lane的数据
	GET /bank/_search
	{
	  "query": { "match": { "address": "mill lane" } }
	}
	（由于ES底层是按照分词索引的，所以上述查询结果是address 字段中包含 mill 或者 lane的数据）


三：查询段落匹配：match_phrase
	如果我们希望查询的条件是 address字段中包含 "mill lane"，则可以使用match_phrase
	GET /bank/_search
	{
	  "query": { "match_phrase": { "address": "mill lane" } }
	}

四：多条件查询: bool
	如果要构造更复杂的查询，可以使用bool查询来组合多个查询条件
	例如，以下请求在bank索引中搜索40岁客户的帐户，但不包括居住在爱达荷州（ID）的任何人




*/
