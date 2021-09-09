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


*/
