package mysql学习

/**
1: 参数是子查询时，使用 EXISTS 代替 IN
	使用 EXISTS 时更快的原因有以下两个。
		● 如果连接列（id）上建立了索引，那么查询 Class_B 时不用查实际的表，只需查索引就可以了。
		● 如果使用 EXISTS，那么只要查到一行数据满足条件就会终止查询，不用像使用 IN 时一样扫描全表。在这一点上 NOT EXISTS 也一样。
2: 参数是子查询时，使用连接代替 IN
	-- 使用连接代替 IN
			SELECT A.id, A.name
			 FROM Class_A A INNER JOIN Class_B B
			 ON A.id = B.id;
3: 避免排序
	会进行排序的代表性的运算有下面这些。
		● GROUP BY 子句
		● ORDER BY 子句
		● 聚合函数（SUM、COUNT、AVG、MAX、MIN）
		● DISTINCT
		● 集合运算符（UNION、INTERSECT、EXCEPT）
		● 窗口函数（RANK、ROW_NUMBER 等）
4: 灵活使用集合运算符的 ALL 可选项
	如果不在乎结果中是否有重复数据，或者事先知道不会有重复数据，请使用 UNION ALL 代替 UNION。这样就不会进行排序了。

5: 使用 EXISTS 代替 DISTINCT
	不好：SELECT DISTINCT I.item_no FROM Items I INNER JOIN SalesHistory SH ON I. item_no = SH. item_no;
	好的写法：SELECT item_no FROM Items I WHERE EXISTS (SELECT * FROM SalesHistory SH WHERE I.item_no = SH.item_no);

6： 在极值函数中使用索引（MAX/MIN）

7： 能写在 WHERE 子句里的条件不要写在 HAVING 子句里

8： 在 GROUP BY 子句和 ORDER BY 子句中使用索引

9： 使用索引时，条件表达式的左侧应该是原始字段
	不好的写法： WHERE col_1 * 1.1 > 100;  好的写法： WHERE col_1 > 100 / 1.1
10： 使用联合索引时，列的顺序错误（最左匹配原则）

11： 使用 LIKE 谓词进行后方一致或中间一致的匹配 使用 LIKE 谓词时，只有前方一致的匹配才能用到索引 列如：SELECT * FROM SomeTable WHERE col_1 LIKE 'a%';

12： 进行默认的类型转换 默认的类型转换不仅会增加额外的性能开销，还会导致索引不可用
	■ 对 char 类型的“col_1”列指定条件的示例（如果使用数字，则不走索引）
		错误： SELECT * FROM SomeTable WHERE col_1 = 10;
		正确： SELECT * FROM SomeTable WHERE col_1 = '10';
		正确： SELECT * FROM SomeTable WHERE col_1 = CAST(10, AS CHAR(2));
13：需要对多个字段使用 IN 谓词时，将它们汇总到一处
	不好的用法：SELECT id, state, city FROM Addresses1 A1 WHERE state IN (SELECT state FROM Addresses2 A2 WHERE A1.id = A2.id)
									               AND city IN (SELECT city FROM Addresses2 A2 WHERE A1.id = A2.id);
	好的使用： SELECT * FROM Addresses1 A1 WHERE (id, state, city) IN (SELECT id, state, city FROM Addresses2 A2);

14： 在指定 IS NULL、IS NOT NULL 的时候，不会用到索引，因而SQL 语句执行起来性能低下。




结论：

	当in()种的数据很大时，不走索引

	当查询的列是char类型没有加引号，mysql优化器会自动给填充引号，同时也会导致索引失效

	当in()中存在子查询、格式化函数等同样也会使索引失效！


explain关键字： https://zhuanlan.zhihu.com/p/114182767
	id:选择标识符 SELECT识别符。
		这是SELECT的查询序列号  我的理解是SQL执行的顺序的标识，SQL从大到小的执行
			1. id相同时，执行顺序由上至下
			2. 如果是子查询，id的序号会递增，id值越大优先级越高，越先被执行
			3. id如果相同，可以认为是一组，从上往下顺序执行；在所有组中，id值越大，优先级越高，越先执行

	select_type:表示查询的类型。  查询中每个select子句的类型
			(1) SIMPLE(简单SELECT，不使用UNION或子查询等)
			(2) PRIMARY(子查询中最外层查询，查询中若包含任何复杂的子部分，最外层的select被标记为PRIMARY)
			(3) UNION(UNION中的第二个或后面的SELECT语句)
			(4) DEPENDENT UNION(UNION中的第二个或后面的SELECT语句，取决于外面的查询)
			(5) UNION RESULT(UNION的结果，union语句中第二个select开始后面所有select)
			(6) SUBQUERY(子查询中的第一个SELECT，结果不依赖于外部查询)
			(7) DEPENDENT SUBQUERY(子查询中的第一个SELECT，依赖于外部查询)
			(8) DERIVED(派生表的SELECT, FROM子句的子查询)
			(9) UNCACHEABLE SUBQUERY(一个子查询的结果不能被缓存，必须重新评估外链接的第一行)

	table:输出结果集的表
			显示这一步所访问数据库中表名称（显示这一行的数据是关于哪张表的），有时不是真实的表名字，可能是简称，例如上面的e，d，也可能是第几步执行的结果的简称
	partitions:匹配的分区

	type:表示表的连接类型
			常用的类型有： 依次从最优到最差分别为：system > const > eq_ref > ref > fulltext > ref_or_null > index_merge > unique_subquery > index_subquery > range > index > ALL
	possible_keys:表示查询时，可能使用的索引
			指出MySQL能使用哪个索引在表中找到记录，查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询使用（该查询可以利用的索引，如果没有任何索引显示 null）
	key:表示实际使用的索引
			key列显示MySQL实际决定使用的键（索引），必然包含在possible_keys中
	key_len:索引字段的长度
			表示索引中使用的字节数，可通过该列计算查询中使用的索引的长度（key_len显示的值为索引字段的最大可能长度，并非实际使用长度，即key_len是根据表定义计算而得，不是通过表内检索出的）
			key_len计算规则如下：
				字符串
					char(n)：n字节长度
					varchar(n)：2字节存储字符串长度，如果是utf-8，则长度 3n + 2
				数值类型
					tinyint：1字节
					smallint：2字节
					int：4字节
					bigint：8字节
				时间类型
					date：3字节
					timestamp：4字节
					datetime：8字节
				如果字段允许为 NULL，需要1字节记录是否为 NULL
	ref:列与索引的比较
			列与索引的比较，表示上述表的连接匹配条件，即哪些列或常量被用于查找索引列上的值
	rows:扫描出的行数(估算的行数)
			估算出结果集行数，表示MySQL根据表统计信息及索引选用情况，估算的找到所需的记录所需要读取的行数
	filtered:按表条件过滤的行百分比
	Extra:执行情况的描述和说明
			该列包含MySQL解决查询的详细信息,有以下几种情况：

				1：Using where：mysql服务器将在存储引擎检索行后再进行过滤。就是先读取整行数据，再按 where 条件进行检查，符合就留下，不符合就丢弃。
				2：Using temporary：mysql需要创建一张临时表来处理查询。出现这种情况一般是要进行优化的，首先是想到用索引来优化。
				3：Using filesort：mysql 会对结果使用一个外部索引排序，而不是按索引次序从表里读取行。此时mysql会根据联接类型浏览所有符合条件的记录，并保存排序关键字和行指针，然后排序关键字并按顺序检索行信息。这种情况下一般也是要考虑使用索引来优化的。
				4：Using join buffer：改值强调了在获取连接条件时没有使用索引，并且需要连接缓冲区来存储中间结果。如果出现了这个值，那应该注意，根据查询的具体情况可能需要添加索引来改进能。
				5： using index： 这发生在对表的请求列都是同一索引的部分的时候，返回的列数据只使用了索引中的信息，而没有再去访问表中的行记录。是性能高的表现。
				6：distinct: 一旦mysql找到了与行相联合匹配的行，就不再搜索了

type：连接类型（建议记到小本本上）
    system          表只有一行
    const           表最多只有一行匹配，通用用于主键或者唯一索引比较时
    eq_ref          每次与之前的表合并行都只在该表读取一行，这是除了system，const之外最好的一种，
                    特点是使用=，而且索引的所有部分都参与join且索引是主键或非空唯一键的索引
    ref             如果每次只匹配少数行，那就是比较好的一种，使用=或<=>，可以是左覆盖索引或非主键或非唯一键
    fulltext        全文搜索
    ref_or_null     与ref类似，但包括NULL
    index_merge     表示出现了索引合并优化(包括交集，并集以及交集之间的并集)，但不包括跨表和全文索引。
                    这个比较复杂，目前的理解是合并单表的范围索引扫描（如果成本估算比普通的range要更优的话）
    unique_subquery 在in子查询中，就是value in (select...)把形如“select unique_key_column”的子查询替换。
                    PS：所以不一定in子句中使用子查询就是低效的！
    index_subquery  同上，但把形如”select non_unique_key_column“的子查询替换
    range           常数值的范围
    index           a.当查询是索引覆盖的，即所有数据均可从索引树获取的时候（Extra中有Using Index）；
                    b.以索引顺序从索引中查找数据行的全表扫描（无 Using Index）；
                    c.如果Extra中Using Index与Using Where同时出现的话，则是利用索引查找键值的意思；
                    d.如单独出现，则是用读索引来代替读行，但不用于查找
    all             全表扫描
*/
