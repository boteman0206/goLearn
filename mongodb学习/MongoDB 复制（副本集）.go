package mongodb学习

/**

一： MongoDB复制原理
	mongodb的复制至少需要两个节点。其中一个是主节点，负责处理客户端请求，其余的都是从节点，负责复制主节点上的数据。
	mongodb各个节点常见的搭配方式为：一主一从、一主多从。

	主节点记录在其上的所有操作oplog，从节点定期轮询主节点获取这些操作，然后对自己的数据副本执行这些操作，从而保证从节点的数据与主节点一致。

	MongoDB复制结构图如下参考： 复制结构图.png


二：副本集特征：
	N 个节点的集群
	任何节点可作为主节点
	所有写入操作都在主节点上
	自动故障转移
	自动恢复

三; MongoDB 副本集可用分两种:
	1: Master-Slave 主从复制
		实现数据同步只需要在某一台服务器启动时加上"-master"参数，以指明此服务器的角色是primary；另一台服务器加上"-slave"和"-source"参数，以指明此服务器的角色是slave。
		主从复制的优点如下：
			从服务器可以执行查询工作，降低主服务器访问压力。
			在从服务器执行备份，避免备份期间锁定主服务器的数据。
			当主服务器出现故障时，可以快速切换到从服务器，减少当机时间。
		注意：MongoDB 的最新版本已不再推荐此方案。主从复制虽然可以承受一定的负载压力，但这种方式仍然是一个单点，如果主库挂了，数据写入就成了风险。
	2: Replica Sets复制集
		MongoDB 在 1.6 版本对开发了新功能replica set，这比之前的replication 功能要强大一 些，增加了故障自动切换和自动修复成员节点，各个DB 之间数据完全一致，大大降低了维 护成功。auto shard 已经明确说明不支持replication paris，建议使用replica set，replica set 故障切换完全自动。
		Replica Sets的结构类似一个集群，完全可以把它当成一个集群，因为它确实与集群实现的作用是一样的：如果其中一个节点出现故障，其他节点马上会将业务接管过来而无须停机操作。

三： MongoDB副本集设置
	1、关闭正在运行的MongoDB服务器。
		现在我们通过指定 --replSet 选项来启动mongoDB。--replSet 基本语法格式如下：
		mongod --port "PORT" --dbpath "YOUR_DB_DATA_PATH" --replSet "REPLICA_SET_INSTANCE_NAME"
			mongod --port 27017 --dbpath "D:\set up\mongodb\data" --replSet rs0
		以上实例会启动一个名为rs0的MongoDB实例，其端口号为27017。
		启动后打开命令提示框并连接上mongoDB服务。
		在Mongo客户端使用命令rs.initiate()来启动一个新的副本集。
		我们可以使用rs.conf()来查看副本集的配置
		查看副本集状态使用 rs.status() 命令

	2：副本集添加成员
		添加副本集的成员，我们需要使用多台服务器来启动mongo服务。进入Mongo客户端，并使用rs.add()方法来添加副本集的成员
		语法
			rs.add() 命令基本语法格式如下：  rs.add(HOST_NAME:PORT)

	3：MongoDB中你只能通过主节点将Mongo服务添加到副本集中， 判断当前运行的Mongo服务是否为主节点可以使用命令db.isMaster() 。
		MongoDB的副本集与我们常见的主从有所不同，主从在主机宕机后所有服务将停止，而副本集在主机宕机后，副本会接管主节点成为主节点，不会出现宕机的情况。


*/
