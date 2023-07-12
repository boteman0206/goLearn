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


GET /megacorp/_doc/_search  // 查询所有，不带参数 默认都是返回10条数据的

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


查询文档的数量
GET index/_count
{
  "query": {
    "match_all": {}
  }
}

短语搜索 为此对 match 查询稍作调整，使用一个叫做 match_phrase 的查询：
GET /megacorp/employee/_search
{
    "query" : {
        "match_phrase" : {
            "about" : "rock climbing"
        }
    }
}
match_phrase: 是只查询 rock climbing的短语，不会进行分词查询
match：  在全文属性上搜索并返回相关性最强的结果 会将rock 和 climbing都查询出来



高亮搜索
再次执行前面的查询，并增加一个新的 highlight 参数：


自动生成id 使用post，不能使用put
POST /website/blog/
{
  "title": "My second blog entry",
  "text":  "Still trying this out...",
  "date":  "2014/01/01"
}


返回文档的一部分
GET 默认返回文档的所有字段，可以指定source返回文档的部分字段
GET /megacorp/_doc/1?_source=first_name,last_name
或者，如果你只想得到 _source 字段，不需要任何元数据，你能使用 _source 端点：
GET /megacorp/_doc/1/_source


检查文档是否存在
如果只想检查一个文档是否存在--根本不想关心内容—​那么用 HEAD 方法来代替 GET 方法。 HEAD 请求没有返回体，只返回一个 HTTP 请求报头：
HEAD /megacorp/_doc/13



并发版本控制，使用if_seq_no和if_primary_term来控制
_seq_no属于整个索引，而不是只属于单一的文档，也就是说每次对索引的任意文档进行更新/新增操作时，索引的_seq_no就会+1（但是要注意的是，虽然索引的_seq_no在不断递增，但是文档中也有一个_seq_no，该文档的_seq_no的值是对他进行最后一次更新后索引的_seq_no值，如果后续不更新，该值并不会改变）
POST /test_index/_doc/2?if_seq_no=19&if_primary_term=1
{
  "doc": {
    "field1": "updated value client 89"
  }
}
查看某一个文档的seq_no
GET /test_index/_doc/2


更新部分字段 // 会在原来的字段上加上 tags和views
POST /test_index/_doc/1/_update
{
   "doc" : {
      "tags" : [ "testing" ],
      "views": 0
   }
}



假设我们需要在 Elasticsearch 中存储一个页面访问量计数器。 每当有用户浏览网页，我们对该页面的计数器进行累加。但是，如果它是一个新网页，我们不能确定计数器已经存在。 如果我们尝试更新一个不存在的文档，那么更新操作将会失败。
在这样的情况下，我们可以使用 upsert 参数，指定如果文档不存在就应该先创建它：
POST /website/_doc/1/_update
{
   "script" : "ctx._source.views+=1",
   "upsert": {
       "views": 1
   }
}

加上retry_on_conflict参数，重试5次
POST /website/pageviews/1/_update?retry_on_conflict=5
{
   "script" : "ctx._source.views+=1",
   "upsert": {
       "views": 0
   }
}


取回多个文档,可以取出不同索引的不同id
GET /_mget
{
   "docs" : [
      {
         "_index" : "test_index",
         "_type" :  "_doc",
         "_id" :    2
      },
      {
         "_index" : "website",
         "_type" :  "_doc",
         "_id" :    1,
         "_source": "views"
      }
   ]
}
如果想检索的数据都在相同的 _index 中（甚至相同的 _type 中），则可以在 URL 中指定默认的 /_index 或者默认的 /_index/_type 。
GET /test_index/_doc/_mget
{
   "ids" : [ "2", "1" ]
}



批量操作  index 相当与没有就创建，有就更新
POST /_bulk
{ "create": { "_index": "website", "_type": "_doc", "_id": "123" }}
{ "title":    "Cannot create - it already exists" }
{ "index":  { "_index": "website", "_type": "_doc", "_id": "1233" }}
{ "title":    "But we can update it" }
：通过批量索引典型文档，并不断增加批量大小进行尝试。 当性能开始下降，那么你的批量大小就太大了。一个好的办法是开始时将 1,000 到 5,000 个文档作为一个批次, 如果你的文档非常大，那么就减少批量的文档个数。
密切关注你的批量请求的物理大小往往非常有用，一千个 1KB 的文档是完全不同于一千个 1MB 文档所占的物理大小。 一个好的批量大小在开始处理后所占用的物理大小约为 5-15 MB。



*/
