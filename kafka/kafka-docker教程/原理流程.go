package kafka_docker教程

/**
https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Ckafka/kafka%E6%B7%B1%E5%B1%82%E4%BB%8B%E7%BB%8D.html

概念介绍：
Producer：Producer即生产者，消息的产生者，是消息的⼊口。
kafka cluster：kafka集群，一台或多台服务器组成
	Broker：Broker是指部署了Kafka实例的服务器节点。每个服务器上有一个或多个kafka的实 例，我们姑且认为每个broker对应一台服务器。每个kafka集群内的broker都有一个不重复的 编号，如图中的broker-0、broker-1等……
	Topic：消息的主题，可以理解为消息的分类，kafka的数据就保存在topic。在每个broker上 都可以创建多个topic。实际应用中通常是一个业务线建一个topic。
	Partition：Topic的分区，每个topic可以有多个分区，分区的作用是做负载，提高kafka的吞 吐量。同一个topic在不同的分区的数据是不重复的，partition的表现形式就是一个一个的⽂件夹！
	Replication:每一个分区都有多个副本，副本的作用是做备胎。当主分区（Leader）故障的 时候会选择一个备胎（Follower）上位，成为Leader。在kafka中默认副本的最大数量是10 个，且副本的数量不能大于Broker的数量，follower和leader绝对是在不同的机器，同一机 器对同一个分区也只可能存放一个副本（包括自己）。
Consumer：消费者，即消息的消费方，是消息的出口。
	Consumer Group：我们可以将多个消费组组成一个消费者组，在kafka的设计中同一个分 区的数据只能被消费者组中的某一个消费者消费。同一个消费者组的消费者可以消费同一个 topic的不同分区的数据，这也是为了提高kafka的吞吐量！


生产流程
    1.⽣产者从Kafka集群获取分区leader信息
    2.⽣产者将消息发送给leader
    3.leader将消息写入本地磁盘
    4.follower从leader拉取消息数据
    5.follower将消息写入本地磁盘后向leader发送ACK
    6.leader收到所有的follower的ACK之后向生产者发送ACK  （这里有isr同步副本的概念）


选择partition的原则
	那在kafka中，如果某个topic有多个partition，producer⼜怎么知道该将数据发往哪个partition呢？ kafka中有几个原则：
		1.partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition。
    	2.如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition。
   	 	3.如果既没指定partition，又没有设置key，则会采用轮询⽅式，即每次取一小段时间的数据写入某个partition，下一小段的时间写入下一个partition


ACK应答机制
	producer在向kafka写入消息的时候，可以设置参数来确定是否确认kafka接收到数据，这个参数可设置 的值为 0,1,all

	0代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效率最高。
	1代表producer往集群发送数据只要leader应答就可以发送下一条，只确保leader发送成功。
	all代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，确保 leader发送成功和所有的副本都完成备份。安全性最⾼高，但是效率最低。
	最后要注意的是，如果往不存在的topic写数据，kafka会⾃动创建topic，partition和replication的数量 默认配置都是1

消费数据
	多个消费者实例可以组成⼀个消费者组，并⽤⼀个标签来标识这个消费者组。⼀个消费者组中的不同消 费者实例可以运⾏在不同的进程甚⾄不同的服务器上。
	如果所有的消费者实例都在同⼀个消费者组中，那么消息记录会被很好的均衡的发送到每个消费者实例。
	如果所有的消费者实例都在不同的消费者组，那么每⼀条消息记录会被⼴播到每⼀个消费者实例。

	参考图： 消费数据.png
	如上图所示⼀个两个节点的Kafka集群上拥有⼀个四个partition（P0-P3）的topic。
	有两个 消费者组都在消费这个topic中的数据，消费者组A有两个消费者实例，消费者组B有四个消费者实例。 从图中我们可以看到，在同⼀个消费者组中，每个消费者实例可以消费多个分区，但是每个分区最多只 能被消费者组中的⼀个实例消费。
	也就是说，如果有⼀个4个分区的主题，那么消费者组中最多只能有4 个消费者实例去消费，多出来的都不会被分配到分区。
	其实这也很好理解，如果允许两个消费者实例同 时消费同⼀个分区，那么就⽆法记录这个分区被这个消费者组消费的offset了。
	如果在消费者组中动态 的上线或下线消费者，那么Kafka集群会⾃动调整分区与消费者实例间的对应关系。




*/
