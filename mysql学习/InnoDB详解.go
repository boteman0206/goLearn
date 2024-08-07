package mysql学习

/**

小林coding: https://xiaolincoding.com/mysql/index/page.html#innodb-%E6%98%AF%E5%A6%82%E4%BD%95%E5%AD%98%E5%82%A8%E6%95%B0%E6%8D%AE%E7%9A%84

InnoDB 是如何存储数据的?
	InnoDB 的数据是按「数据页」为单位来读写的,当需要读一条记录的时候，并不是将这个记录本身从磁盘读出来，而是以页为单位，将其整体读入内存。
	数据库的 I/O 操作的最小单位是页: InnoDB 数据页的默认大小是 16KB，意味着数据库每次读写都是以 16KB 为单位的，一次最少从磁盘中读取 16K 的内容到内存中，一次最少把内存中的 16K 内容刷新到磁盘中。


为什么 MySQL 采用 B+ 树作为索引？
	磁盘是一个慢的离谱的存储设备，磁盘读写的最小单位是扇区，扇区的大小只有 512B 大小，操作系统一次会读写多个扇区，所以操作系统的最小读写单位是块（Block）。
	Linux 中的块大小为 4KB，也就是一次磁盘 I/O 操作会直接读写 8 个扇区。
*/

/*
MySQL 可重复读隔离级别，完全解决幻读了吗？
MySQL InnoDB 引擎的默认隔离级别虽然是「可重复读」，但是它很大程度上避免幻读现象（并不是完全解决了），解决的方案有两种：
1: 针对快照读（普通 select 语句），是通过 MVCC 方式解决了幻读，因为可重复读隔离级别下，事务执行过程中看到的数据，一直跟这个事务启动时看到的数据是一致的，
即使中途有其他事务插入了一条数据，是查询不出来这条数据的，所以就很好了避免幻读问题。

2:针对当前读（select ... for update 等语句），是通过 next-key lock（记录锁+间隙锁）方式解决了幻读，因为当执行 select ... for update 语句的时候，
会加上 next-key lock，如果有其他事务在 next-key lock 锁范围内插入了一条记录，那么这个插入语句就会被阻塞，无法成功插入，所以就很好了避免幻读问题。

什么是幻读？
当同一个查询在不同的时间产生不同的结果集时，事务中就会出现所谓的幻象问题。例如，如果 SELECT 执行了两次，但第二次返回了第一次没有返回的行，则该行是“幻像”行。

mysql的可重复读发生的几个场景？
1: 在一个事务a中，更新另外一个事务b中已经提交的记录，会发生幻读
	1.1：事务 A 第一次执行普通的 select 语句时生成了一个 ReadView
	1.2：之后事务 B 向表中新插入了一条 id = 5 的记录并提交
	1.3：接着，事务 A 对 id = 5 这条记录进行了更新操作，在这个时刻，这条新记录的 trx_id 隐藏列的值就变成了事务 A 的事务 id，
	1.4: 之后事务 A 再使用普通 select 语句去查询这条记录时就可以看到这条记录了，于是就发生了幻读。
2:
	2.1：T1 时刻：事务 A 先执行「快照读语句」：select * from t_test where id > 100 得到了 3 条记录。
	2.2：T2 时刻：事务 B 往插入一个 id= 200 的记录并提交；
	2.3：T3 时刻：事务 A 再执行「当前读语句」 select * from t_test where id > 100 for update 就会得到 4 条记录，此时也发生了幻读现象。
	要避免这类特殊场景下发生幻读的现象的话，就是尽量在开启事务之后，马上执行 select ... for update 这类当前读的语句，因为它会对记录加 next-key lock，从而避免其他事务插入一条新记录。

*/

/**
mysql的锁
1： 全局锁
flush table test1  with read lock;
unlock tables;  表被锁定无法进行增删改查的操作可以使用unlock tables
全局锁应用场景是什么？
全局锁主要应用于做全库逻辑备份，这样在备份数据库期间，不会因为数据或表结构的更新，而出现备份文件的数据与预期的不一样

2： 表级锁（表锁，元数据锁，意向锁，auto-inc锁）
表锁：
	//表级别的共享锁，也就是读锁；
	lock tables t_student read;
	//表级别的独占锁，也就是写锁；
	lock tables t_stuent write;
    unlock tables   // 要释放表锁
元数据锁（MDL）：
	我们不需要显示的使用 MDL，因为当我们对数据库表进行操作时，会自动给这个表加上 MDL：
		对一张表进行 CRUD 操作时，加的是 MDL 读锁；
		对一张表做结构变更操作的时候，加的是 MDL 写锁；
	MDL 是为了保证当用户对表执行 CRUD 操作时，防止其他线程对这个表结构做了变更。
	MDL 是在事务提交后才会释放，这意味着事务执行期间，MDL 是一直持有的。
意向锁：
	在使用 InnoDB 引擎的表里对某些记录加上「共享锁」之前，需要先在表级别加上一个「意向共享锁」；
	在使用 InnoDB 引擎的表里对某些纪录加上「独占锁」之前，需要先在表级别加上一个「意向独占锁」；
	也就是，当执行插入、更新、删除操作，需要先对表加上「意向独占锁」，然后对该记录加独占锁。
	意向锁的目的是为了快速判断表里是否有记录被加锁。

AUTO-INC 锁
在为某个字段声明 AUTO_INCREMENT 属性时，之后可以在插入数据时，可以不指定该字段的值，数据库会自动给该字段赋值递增的值，这主要是通过 AUTO-INC 锁实现的。
AUTO-INC 锁是特殊的表锁机制，锁不是再一个事务提交后才释放，而是再执行完插入语句后就会立即释放。

3： 行级锁（record lock， gap lock， next-key lock） InnoDB 引擎是支持行级锁的，而 MyISAM 引擎并不支持行级锁。
行级锁的类型主要有三类：
Record Lock，记录锁，也就是仅仅把一条记录锁上；
Gap Lock，间隙锁，锁定一个范围，但是不包含记录本身； 只存在于可重复读隔离级别，目的是为了解决可重复读隔离级别下幻读的现象。
Next-Key Lock：Record Lock + Gap Lock 的组合，锁定一个范围，并且锁定记录本身。
插入意向锁：一个事务在插入一条记录的时候，需要判断插入位置是否已被其他事务加了间隙锁（next-key lock 也包含间隙锁）。
		  插入意向锁名字虽然有意向锁，但是它并不是意向锁，它是一种特殊的间隙锁，属于行级别锁。



主要结论：
普通的 select 是不会加行级锁的，普通的 select 语句是利用 MVCC 实现一致性读，是无锁的。
select 也是可以对记录加共享锁和独占锁的，具体方式如下：
	//先在表上加上意向共享锁，然后对读取的记录加共享锁
	select ... lock in share mode;
	//先表上加上意向独占锁，然后对读取的记录加独占锁
	select ... for update;


*/
