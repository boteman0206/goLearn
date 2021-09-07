package mysql学习

/**
1: 用于删除重复行的sql
	1.1：使用极值函数（在mysql中会报错误，需要早子查询的sql中再套接一层select * from 即可）
		DELETE FROM Products P1
			 WHERE rowid < ( SELECT MAX(P2.rowid)
			 FROM Products P2
			 WHERE P1.name = P2. name
			 AND P1.price = P2.price ) ;
	1.2：使用非等值连接 （在mysql中会报错误，需要早子查询的sql中再套接一层select * from 即可）
		DELETE FROM Products P1
			 WHERE EXISTS ( SELECT *
			 FROM Products P2
			 WHERE P1.name = P2.name
			 AND P1.price = P2.price
			 AND P1.rowid < P2.rowid );
2：特殊的排序（除了用窗口函数之外还可以使用自连接）
	2.1： 窗口函数
			SELECT name, price,
			 RANK() OVER (ORDER BY price DESC) AS rank_1,
			 DENSE_RANK() OVER (ORDER BY price DESC) AS rank_2
			 FROM Products;
	2.2： 子查询方式： 排序从 1 开始。如果已出现相同位次，则跳过之后的位次
			SELECT P1.name,	P1.price, (
				SELECT COUNT(P2.price)
				 FROM Products P2
				 WHERE P2.price > P1.price) + 1 AS rank_1
				 FROM Products P1
				 ORDER BY rank_1;
		  扩展：例如去掉标量子查询后边的 +1，就可以从 0 开始给商品排序，而且如果修改成COUNT(DISTINCT P2.price)，
		  那么存在相同位次的记录时，就可以不跳过之后的位次，而是连续输出（相当于 DENSE_RANK 函数）。
	2.3： 使用自连接方式
			SELECT P1.name, MAX(P1.price) AS price, COUNT(P2.name) +1 AS rank_1
				 	FROM Products P1
					LEFT OUTER JOIN Products P2 ON P1.price < P2.price
				 GROUP BY P1.name
				 ORDER BY rank_1;

3：mysql的三值运算（三值运算图.jpg）
    3.1： NOT IN 和 NOT EXISTS 不是等价的
		（NOT IN 子查询中用到的表里被选择的列中存在 NULL，则 SQL 语句整体的查询结果永远是空。这是很可怕的现象。
          为了得到正确的结果，我们需要使用 EXISTS 谓词。）
		SELECT * FROM Class_A A
			 WHERE NOT EXISTS ( SELECT *
			 FROM Class_B B
			 WHERE A.age = B.age
			 AND B.city = '东京' );


4： 限定谓词和极值函数不是等价的: ALL 谓词和ANY谓语和极值函数大多数时候都是一样的，但是在有null的情况下是不同的
	原因：ALL 谓词其实是多个以 AND 连接的逻辑表达式的省略写法，所以有null的时候还是会返回不确定unknown
         极值函数max, min等在统计时会把为 NULL 的数据排除掉


5：按照现在的 SQL 标准来说，HAVING子句是可以单独使用的
 	5.1：having寻找缺失的编号
	-- 如果有查询结果，说明存在缺失的编号
		SELECT '存在缺失的编号' AS gap
		 FROM SeqTbl
		HAVING COUNT(*) <> MAX(seq);
	 -- 查询缺失编号的最小值
		SELECT MIN(seq + 1) AS gap
		 FROM SeqTbl
		 WHERE (seq+ 1) NOT IN ( SELECT seq FROM SeqTbl);
	5.2： 用 HAVING 子句进行子查询 ：求众数（它指的是在群体中出现次数最多的值，）
	 -- 求众数的 SQL 语句 (1)：使用谓词
		SELECT income, COUNT(*) AS cnt
		 FROM Graduates
		 GROUP BY income
		HAVING COUNT(*) >= ALL ( SELECT COUNT(*)
		 FROM Graduates GROUP BY income);


*/
