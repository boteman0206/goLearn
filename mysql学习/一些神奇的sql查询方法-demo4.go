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
	5.3： 用 HAVING 子句进行自连接 ：求中位数
		-- 求中位数的 SQL 语句 ：在 HAVING 子句中使用非等值自连接
			（比较难理解：大意就是先交叉链接然后通过t2表中的income和t1表的income作比较，取出个数>=总数一半的那个值，如果是奇数会取出一个，如果是偶数取出一个然后取中位数）
		SELECT AVG(DISTINCT income)
			 FROM (
				 SELECT T1.income
					 FROM Graduates T1, Graduates T2
				 GROUP BY T1.income HAVING SUM(CASE WHEN T2.income >= T1.income THEN 1 ELSE 0 END)  >= COUNT(*) / 2
				 AND SUM(CASE WHEN T2.income <= T1.income THEN 1 ELSE 0 END) >= COUNT(*) / 2
				  ) TMP;
	5.4: 用外连接进行行列转换




存储过程：
	1:变量名 所有MySQL变量都必须以@开始。
	2：COMMENT关键字 本例子中的存储过程在CREATE PROCEDURE语句中包含了一个COMMENT值。它不是必需的，但如果给出，将在SHOW PROCEDURE STATUS的结果中显示。

	3：无参数的
	create procedure productpring()  -- 创建存储过程
	  begin
		select max(emp_no) from employees;
	  end;


	call productpring();  -- 执行存储过程
	drop procedure if exists productpring; -- 删除存储过程

	4：有参数的
		MySQL支持IN（传递给存储过程）、OUT（从存储过程传出，如这里所用）和INOUT（对存储过程传入和传出）类型的参数
		-- 有参数的存储过程创建
		2.1:
			create procedure productpricing(
			  out p1_min decimal(8,2),
			  out p2_max decimal(8,2),
			  out p3_avg decimal(8,2)
			)
			  begin
				select min(emp_no) into p1_min from employees;
				select max(emp_no) into p2_max from employees;
				select avg(emp_no) into p3_avg from employees;
			  end;


		call productpricing(@a, @b, @c); -- 调用
		select @a, @b, @c;

		2.2: 混合参数这次使用IN和OUT参数
			create procedure  ordertotal(
				  in in_data int,
				  out out_data decimal(10,3)
				)
				  begin
					select sum(emp_no) + in_data  from employees into out_data;
				  end;

				call ordertotal(12, @p_data);
				select @p_data;

	5：弯针版本的total
		drop procedure if exists ordertotalmain;
		create procedure ordertotalmain(
		  in onnumber int,
		  in taxable boolean,
		out ototal decimal(8,2)
		) comment '这是一个完整的存储过程示例'
		  begin
			  -- 申明一个变量 total
			declare total decimal(8,2);
			-- 申明一个int默认值
			declare tax int default 6;

			-- 查询获取total
			select sum(emp_no) from employees where emp_no=onnumber into total;

			-- is this true or false
			if taxable then
			  -- yes
			  select total + (total/100*tax) into total;
			end if;
			  -- and finally save to out
			  select total into ototal;
		  end;

		call ordertotalmain(10001, true , @total);
		select @total;

	6：检查存储过程：
		show create procedure ordertotal;
		SHOW PROCEDURE STATUS like '%ordertotal%';


游标：
	1： 只能用于存储过程 不像多数DBMS，MySQL游标只能用于存储过程（和函数）。

	2：使用示例# 使用游标 mysql的游标只能在存储过程中使用
		# 使用游标 mysql的游标只能在存储过程中使用
		create procedure processorders(out num int)
		  begin

			declare done boolean default 0;
			declare ordernumbers cursor for
			  select emp_no from employees;

			declare continue handler for sqlstate '02000' set done =1 ;  ---改变结束条件

			open ordernumbers;  -- 开启游标
			  repeat  -- 开启循环
				fetch ordernumbers into num;

			  until done  end repeat;  -- 结束循环
			close ordernumbers;  -- 关闭游标

		  end;


		call processorders(@fetch);

		select @fetch;



触发器：
	需要在某个表发生更改时自动处理。这确切地说就是触发器。触发器是MySQL响应以下任意语句而自动执行的一条MySQL语句（或位于BEGIN和END语句之间的一组语句）： DELETE；  INSERT；  UPDATE。
	1：创建触发器 保持每个数据库的触发器名唯一 用CREATE TRIGGER语句创建
		CREATE TRIGGER newproduct AFTER INSERT ON products FOR EACH ROW SELECT NEW.name INTO @asd;  // 注意点： mysql trigger的返回值必须使用变量接受，不能直接返回，否则会报错。
		insert into products values("tes1t", 19, 8);
		select @asd;
	2: 在INSERT触发器代码内，可引用一个名为NEW的虚拟表，访问被插入的行；

	3： 仅支持表 只有表才支持触发器，视图不支持（临时表也不支持）。

	4： 每个表最多支持6个触发器（每条INSERT、UPDATE和DELETE的之前和之后）。单一触发器不能与多个事件或多个表关联，所以，如果你需要一个对INSERT和UPDATE操作执行的触发器，则应该定义两个触发器。

	5：触发器失败 如果BEFORE触发器失败，则MySQL将不执行请求的操作。此外，如果BEFORE触发器或语句本身失败，MySQL将不执行AFTER触发器（如果有的话）
	6：删除触发器  drop trigger newproduct;

	7： 多语句触发器使用begin 。。。 end
		CREATE TRIGGER deleteorder BEFORE DELETE ON orders FOR EACH ROW
		  BEGIN
			INSERT INTO archive_orders (order_num, order_date, cust_id)VALUES(OLD.order num, OLD.order date, OLD.cust_id);
		  END;
*/
