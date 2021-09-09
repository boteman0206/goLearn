package mysql学习

/**


https://zhuanlan.zhihu.com/p/164511591
https://cloud.tencent.com/developer/article/1491329
https://tech.meituan.com/2014/06/30/mysql-index.html
https://www.gxlcms.com/mysql-463427.html

https://zhuanlan.zhihu.com/p/352436463


一：mysql的分层：
	1：通过对mysql的两张图： mysql架构图.jpg和mysql的执行图.jpg我们知道 简单来说，MySQL主要分为 Server层 和 存储引擎层：
		Server层：主要包括连接器、查询缓存（MySQL8.0移除）、分析器、优化器、执行器等，所有的跨存储引擎的功能都在这一层实现，
			     比如存储过程、触发器、视图、函数等，还有一个通用的日志模块binglog
		存储引擎层：主要负责数据的存储和读取，采用可以替换的插件式架构，支持InnoDB、MyISAM、Memory等多个存储引擎，其中InnnoDB有属于自己的日志模块（下文会介绍到）。
				  「现在最常用的存储引擎是InnoDB，它从MySQL5.5.5版本开始被当做默认的存储引擎了。」



	2：查询缓存」（MySQL8.0后移除） 查询缓存主要用来缓存我们所执行的 SELECT 语句以及该语句的结果集。
		为什么MySQL8.0后要移除呢？
		因为查询缓存失效在实际业务场景中可能会非常频繁，假如你对一个表更新的话，这个表上的所有的查询缓存都会被清空。对于不经常更新的数据来说，使用缓存还是可以的。所以，一般在大多数情况下我们都是不推荐去使用查询缓存的。

	3：binlog（归档日志）是MySQL的Server层有的。 逻辑日志 主要记录用户对数据库操作的SQL语句

*/
