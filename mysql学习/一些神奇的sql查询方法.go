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


*/
