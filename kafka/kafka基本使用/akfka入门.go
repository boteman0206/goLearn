package kafka基本使用

/**


消息队列的两种基本工作模式：
	1： 点对点  一对一
    2： 发布订阅  一对多

消费者消费方式
	1： 服务端主动推送
	2： 消费者去pull拉取消息
	Kafka只有消息的拉取，没有推送，可以通过轮询实现消息的推送。



kafka概念：
	broker： 相当于一个kafka的进程或者线程，类似于一抬服务器
	topic： 可以理解为一个队列，不同的数据放到不同的topic里面，生产者和消费者面向的都是一个topic
	partition： 分区，为了实现扩展性。非常大的一个topic可以分到多个broker上，一个topic可以划分为多个分区，每个分区是一个有序的队列。
	replica：副本，为保证集群中的某一个节点发生故障，该节点上的partition数据不丢失，切kafka能够正常工作而提供的一个副本机制，一个leader和若干个follower。
    leader：每个分区多个副本的主，生产者发送数据的对象，以及消费者消费数据的对象都是leader。
	follower：每个分区多个副本的从节点，实时从leader中同步数据，保持和leader数据的同步。leader发生故障的时候，某个follower会成为新的leader。
	groupName： 消费者组，将多个消费者归于一个组下（消费者组保证了组内消息不会被重复消费）
	consumer： 消费者
    offset： 消息的偏移量，消费的位置（0.9版本之前保存在zookeeper里面，高版本的直接存在kafka的一个系统topic里面）


主题和分区：
	Kafka的消息通过主题进行分类。主题可比是数据库的表或者文件系统里的文件夹。主题可以被分为若干分区，一个主题通过分区分布于Kafka集群中，提供了横向扩展的能力。

生产者和消费者：（生产者创建消息。消费者消费消息。）
	一个消息被发布到一个特定的主题上。生产者在默认情况下把消息均衡地分布到主题的所有分区上：
		1. 直接指定消息的分区
		2. 根据消息的key散列取模得出分区
		3. 轮询指定分区。
	消费者通过偏移量来区分已经读过的消息，从而消费消息。 消费者是消费组的一部分。消费组保证每个分区只能被一个消费者使用，避免重复消费。
		1. 消费者订阅一个或多个主题，并按照消息生成的顺序读取它们。
        2. 消费者通过检查消息的偏移量来区分已经读取过的消息。偏移量是另一种元数据，它是一个不断递增的整数值，在创建消息时，Kafka 会把它添加到消息里。在给定的分区里，每个消息的偏移量都是唯一的。消费者把每个分区最后读取的消息偏移量保存在Zookeeper 或Kafka上，如果消费者关闭或重启，它的读取状态不会丢失。
		3. 消费者是消费组的一部分。群组保证每个分区只能被一个消费者使用。
		4. 如果一个消费者失效，消费组里的其他消费者可以接管失效消费者的工作，再平衡，分区重新分配。
Broker：
	一个独立的Kafka 服务器被称为broker。 broker 为消费者提供服务，对读取分区的请求作出响应，返回已经提交到磁盘上的消息。
		1. 如果某topic有N个partition，集群有N个broker，那么每个broker存储该topic的一个partition。

topic：
	每条发布到Kafka集群的消息都有一个类别，这个类别被称为Topic。 物理上不同Topic的消息分开存储。 主题就好比数据库的表，尤其是分库分表之后的逻辑表。

Partition
	1：主题可以被分为若干个分区，一个分区就是一个提交日志。
	2：消息以追加的方式写入分区，然后以先入先出的顺序读取。
	3：无法在整个主题范围内保证消息的顺序，但可以保证消息在单个分区内的顺序。
	4：Kafka 通过分区来实现数据冗余和伸缩性。
	5：在需要严格保证消息的消费顺序的场景下，需要将partition数目设为1

Offset：
	生产者Offset 消息写入的时候，每一个分区都有一个offset，这个offset就是生产者的offset，同时也是这个分区的最新最大的offset。



发送数据可靠性保证
	为保证 producer 发送的数据，能可靠的发送到指定的 topic，topic 的每个 partition 收到 producer 发送的数据后，都需要向 producer 发送 ack（acknowledgement 确认收到），如果 producer 收到 ack，就会进行下一轮的发送，否则重新发送数据。

生产者数据的一致性：
副本数据同步策略
	1: 半数以上   需要 2n+1个副本
	2： 全部完成   需要n+1个副本
	Kafka 选择了第二种方案，原因如下：
		1：同样为了容忍 n 台节点的故障，第一种方案需要 2n+1 个副本，而第二种方案只需要 n+1 个副本，而 Kafka 的每个分区都有大量的数据，第一种方案会造成大量数据的冗余。
		2：虽然第二种方案的网络延迟会比较高，但网络延迟对 Kafka 的影响较小（同一网络环境下的传输）。

ISR: 同步副本（包含leader+符合条件的follower）
	因为采用数据同步的第二种方案的话，可能由于其中一个副本响应时间很慢，就会影响leader的ack延迟时间，严重影响效率
	这里会维护一个动态的in-sync replication set , 意为和leader保持同步的follower集合，当isr中的follower数据同步完成之后，leader就会给follower发送ack。如果follower长时间没有向leader同步数据，则该follower将被剔除isr，该时间阀值有参数
	replica.lag.time.max.ms参数设定。leader发生故障之后，就会从isr中选择新的leader。

ack机制：
	0： producer 不等待broker的ack，最低延迟，发完就不管了，但是broker接受到还没有写入磁盘就已经返回，有可能因为故障丢失数据。
	1： 只等待leader的ack，如果follower同步成功之前leader故障，可能会丢失数据
    -1（all）：producer等待broker的ack，partition的leader和（isr中的）follower全部落盘成功后才返回ack。但是如果follower同步完成后。broker发送ack之前，leader发生故障，可能造成数据重复。
			  isr里面只剩下一个leader的时候也是会丢数据的。



副本的数据一致性保证：
	log文件的hw（hight watermark）和leo（log end offset）
	LEO: 每一个副本的最后一个offset 可能不一样
	HW： 对于consumer而言，只有最小的的一个LEO才可以被消费。
	follower故障的时候：
		被剔除isr，等待故障恢复之后，follower会读取本地的hw，将log文件高于hw的部分截取。从hw开始向leader同步数据，等该follower的leo大于等于该partition的HW之后，就认为该follower追上了leader，就可以重新加入isr集合了。
	leader故障之后：
		会从isr中选取新的leader，之后为了保证各个副本之间数据的一致性，其余的follower会先将各自的log文件高于hw的部分截取。然后从新的leader同步数据。
	这里只是保证副本的数据一致性，数据丢不丢是由ack决定的



消息中间的 exactly once恰好一次：
	ack设置0： at most once 可能会丢数据
	ack设置 -1 ： at least once  数据可能会重复
	at least once + 冪等 = exactly once  0.11版本引入的一个冪等特性




分区分配策略：消费者改变的时候（增加或者减少的时候）
	roundrobin 按照消费者组来划分的
    range（默认）  按照topic主题划分的







*/
