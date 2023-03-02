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
	Consumer Group：我们可以将多个消费组组成一个消费者组，在kafka的设计中同一个分 区的数据只能被消费者组中的某一个消费者消费。
	同一个消费者组的消费者可以消费同一个 topic的不同分区的数据，这也是为了提高kafka的吞吐量！每个消费者都属于某个消费者组，如果不指定，那么所有的消费者都属于默认的组


生产流程
    1.⽣产者从Kafka集群获取分区leader信息
    2.⽣产者将消息发送给leader
    3.leader将消息写入本地磁盘
    4.follower从leader拉取消息数据
    5.follower将消息写入本地磁盘后向leader发送ACK
    6.leader收到所有的follower的ACK之后向生产者发送ACK  （这里有isr同步副本的概念）
消费者和生产者都是从leader读写数据，不与follower交互。

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



//https://juejin.cn/post/7102243362471673892
消息传传递语义
	1: 最多一次
		从Producer的角度来看， At most once意味着Producer发送完一条消息后，不会确认消息是否成功送达。这样从Producer的角度来看，消息仅仅被发送一次，也就存在者丢失的可能性。
	2： 最少一次
		从Producer的角度来看，At least once意味着Producer发送完一条消息后，会确认消息是否发送成功。如果Producer没有收到Broker的ack确认消息，那么会不断重试发送消息。这样就意味着消息可能被发送不止一次，也就存在这消息重复的可能性。
	3： 恰好一次
		从Producer的角度来看，Exactly once意味着Producer消息的发送是幂等的。这意味着不论消息重发多少遍，最终Broker上记录的只有一条不重复的数据。

生产者端消息保证：
	Producer At least once配置
		Kafka默认的Producer消息送达语义就是At least once，这意味着我们不用做任何配置就能够实现At least once消息语义。原因是Kafka中默认`acks=1`并且`retries=2147483647`。
	Producer At most once配置
		acks=0。acks配置项表示Producer期望的Broker的确认数。默认值为1。可选项：[0，1，all]。如果设置为0，表示Producer发送完消息后不会等待任何Broker的确认；设置为1表示Producer会等待Broker集群中的leader的确认写入消息；设置为all表示Producer需要等待Broker集群中leader和其所有follower的确认写入消息。
	Producer Exactly once配置
		Exactly once是Kafka从版本0.11之后提供的高级特性。我们可以通过配置Producer的以下配置项来实现Exactly once语义：
			enable.idempotence=true。enable.idempotence配置项表示是否使用幂等性。当enable.idempotence配置为true时，acks必须配置为all。并且建议max.in.flight.requests.per.connection的值小于5。
			acks=all。

消费者端消息保证：
	1：从Consumer的角度来看，At most once意味着Consumer对一条消息最多消费一次，因此有可能存在消息消费失败依旧提交offset的情况。考虑下面的情况：Consumer首先读取消息，然后提交offset，最后处理这条消息。在处理消息时，Consumer宕机了，此时offset已经提交，下一次读取消息时读到的是下一条消息了，这就是At most once消费

	2：从Consumer的角度来看，At least once意味着Consumer对一条消息可能消费多次。考虑下面的情况：Consumer首先读取消息，然后处理这条消息，最后提交offset。在处理消息时成功后，Consumer宕机了，此时offset还未提交，下一次读取消息时依旧是这条消息，那么处理消息的逻辑又将被执行一遍，这就是At least once消费。
		Consumer At least once配置
			enable.auto.commit=false。禁止后台自动提交offset。
			手动调用consumer.commitSync()来提交offset。手动调用保证了offset即时更新。
			通过手动提交offset，就可以实现Consumer At least once语义。
	3：从Consumer的角度来看，Exactly once意味着消息的消费处理逻辑和offset的提交是原子性的，即消息消费成功后offset改变，消息消费失败offset也能回滚。
		Consumer Exactly once配置
			isolation.level=read_committed。isolation.level表示何种类型的message对Consumer可见。 也即事务机制，后面在理解。


Broker 端丢失场景剖析
	KafkaBroker 集群接收到数据后会将数据进行持久化存储到磁盘，为了提高吞吐量和性能，采用的是「异步批量刷盘的策略」，
	也就是说按照一定的消息量和间隔时间进行刷盘。首先会将数据存储到 「PageCache」 中，至于什么时候将 Cache 中的数据刷盘是由「操作系统」根据自己的策略决定或者调用 fsync 命令进行强制刷盘，
	如果此时 Broker 宕机 Crash 掉，且选举了一个落后 Leader Partition 很多的 Follower Partition 成为新的 Leader Partition，那么落后的消息数据就会丢失。

	1： 由于 Kafka 中并没有提供「同步刷盘」的方式，所以说从单个 Broker 来看还是很有可能丢失数据的。
	2： kafka 通过「**多 Partition （分区）多 Replica（副本）机制」**已经可以最大限度的保证数据不丢失，如果数据已经写入 PageCache 中但是还没来得及刷写到磁盘，此时如果所在 Broker 突然宕机挂掉或者停电，极端情况还是会造成数据丢失。



kafka的幂等  参考图 kafka幂等.png
	https://www.cnblogs.com/smartloli/p/11922639.html
	 当Producer发送消息(x2,y2)给Broker时，Broker接收到消息并将其追加到消息流中。
	此时，Broker返回Ack信号给Producer时，发生异常导致Producer接收Ack信号失败。对于Producer来说，会触发重试机制，将消息(x2,y2)再次发送，
	但是，由于引入了幂等性，在每条消息中附带了PID（ProducerID）和SequenceNumber。相同的PID和SequenceNumber发送给Broker，
	而之前Broker缓存过之前发送的相同的消息，那么在消息流中的消息就只有一条(x2,y2)，不会出现重复发送的情况。



Kafka 事务  https://juejin.cn/post/7122295644919693343
	幂等性也只能保证单分区、单会话内的数据不重复，如果 Kafka 挂掉，重新给生产者分配了 PID，还是有可能产生重复的数据，这就需要另一个特性来保证了——Kafka 事务。

	Kafka 事务基于幂等性实现，通过事务机制，Kafka 可以实现对多个 Topic 、多个 Partition 的原子性的写入，即处于同一个事务内的所有消息，最终结果是要么全部写成功，要么全部写失败。


*/
