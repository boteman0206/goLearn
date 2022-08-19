package mysql学习

/**

查看正在执行的sql
select * from information_schema.processlist WHERE info IS NOT NULL ORDER BY TIME desc;

查看当前的连接数
show status like  'Threads%';




修改表的字段的字符集 // 直接修改库不生效，需要这样才行
alter table dc_oms.product_sku_third convert to character set utf8 COLLATE utf8_general_ci;


*/
