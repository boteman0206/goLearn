package mysql学习

/**

查看正在执行的sql
select * from information_schema.processlist WHERE info IS NOT NULL ORDER BY TIME desc;

查看当前的连接数
show status like  'Threads%';



*/
