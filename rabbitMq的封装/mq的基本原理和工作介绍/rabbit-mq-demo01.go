package mq的基本原理和工作介绍

/**

http://rabbitmq.mr-ping.com/AMQP/AMQP_0-9-1_Model_Explained.html#amqp-0-9-1-%E7%AE%80%E4%BB%8B
http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9CRabbitMQ/

一： 基本概念(参照图示表示)
    Broker：标识消息队列服务器实体.
    Virtual Host：虚拟主机。标识一批交换机、消息队列和相关对象。虚拟主机是共享相同的身份认证和加密环境的独立服务器域。每个vhost本质上就是一个mini版的RabbitMQ服务器，拥有自己的队列、交换器、绑定和权限机制。vhost是AMQP概念的基础，必须在链接时指定，RabbitMQ默认的vhost是 /。
    Exchange：交换器，用来接收生产者发送的消息并将这些消息路由给服务器中的队列。
    Queue：消息队列，用来保存消息直到发送给消费者。它是消息的容器，也是消息的终点。一个消息可投入一个或多个队列。消息一直在队列里面，等待消费者连接到这个队列将其取走。
    Banding：绑定，用于消息队列和交换机之间的关联。一个绑定就是基于路由键将交换机和消息队列连接起来的路由规则，所以可以将交换器理解成一个由绑定构成的路由表。
    Channel：信道，多路复用连接中的一条独立的双向数据流通道。新到是建立在真实的TCP连接内地虚拟链接，AMQP命令都是通过新到发出去的，不管是发布消息、订阅队列还是接收消息，这些动作都是通过信道完成。因为对于操作系统来说，建立和销毁TCP都是非常昂贵的开销，所以引入了信道的概念，以复用一条TCP连接。
    Connection：网络连接，比如一个TCP连接。
    Publisher：消息的生产者，也是一个向交换器发布消息的客户端应用程序。
    Consumer：消息的消费者，表示一个从一个消息队列中取得消息的客户端应用程序。
    Message：消息，消息是不具名的，它是由消息头和消息体组成。消息体是不透明的，而消息头则是由一系列的可选属性组成，这些属性包括routing-key(路由键)、priority(优先级)、delivery-mode(消息可能需要持久性存储[消息的路由模式])等。



二：工作模式
    1：simple模式 简单工作模式
        消息产生着§将消息放入队列
        消息的消费者(consumer) 监听(while) 消息队列,如果队列中有消息,就消费掉,消息被拿走后,
        自动从队列中删除(隐患 消息可能没有被消费者正确处理,已经从队列中消失了,造成消息的丢失)应用场景:聊天(中间有一个过度的服务器;p端,c端)

    2： work模式
        消息产生者将消息放入队列消费者可以有多个,消费者1,消费者2,同时监听同一个队列,消息被消费?C1 C2共同争抢当前的消息队列内容,
        谁先拿到谁负责消费消息(隐患,高并发情况下,默认会产生某一个消息被多个消费者共同使用,可以设置一个开关(syncronize,与同步锁的性能不一样) 保证一条消息只能被一个消费者使用)
        应用场景:红包;大项目中的资源调度(任务分配系统不需知道哪一个任务执行系统在空闲,直接将任务扔到消息队列中,空闲的系统自动争抢)
    3： publish模式
        X代表交换机rabbitMQ内部组件,erlang 消息产生者是代码完成,代码的执行效率不高,消息产生者将消息放入交换机,交换机发布订阅把消息发送到所有消息队列中,对应消息队列的消费者拿到消息进行消费
        相关场景:邮件群发,群聊天,广播(广告)
    4： routing模式
         消息生产者将消息发送给交换机按照路由判断,路由是字符串(info) 当前产生的消息携带路由字符(对象的方法),交换机根据路由的key,只能匹配上路由key对应的消息队列,对应的消费者才能消费消息;
         根据业务功能定义路由字符串
         从系统的代码逻辑中获取对应的功能字符串,将消息任务扔到对应的队列中业务场景:error 通知;
         EXCEPTION;错误通知的功能;传统意义的错误通知;客户通知;利用key路由,可以将程序中的错误封装成消息传入到消息队列中,开发者可以自定义消费者,实时接收错误;
    5： topic模式
         星号井号代表通配符
         星号代表多个单词,井号代表一个单词
         路由功能添加模糊匹配
         消息产生者产生消息,把消息交给交换机
         交换机根据key的规则模糊匹配到对应的队列,由队列的监听消费者接收消息消费

    6： 发布确认模式（先不做处理）






三： rabbitMQ 四种类型交换器 Fanout,Direct,Topic和 headers
    1：Fanout Exchange
        不处理路由键。你只需要简单的将队列绑定到交换机上。一个发送到交换机的消息都会被转发到与该交换机绑定的所有队列上。
        很像子网广播，每台子网内的主机都获得了一份复制的消息。Fanout交换机转发消息是最快的
    2：Direct Exchange
        处理路由键。需要将一个队列绑定到交换机上，要求该消息与一个特定的路由键完全匹配。
        这是一个完整的匹配。如果一个队列绑定到该交换机上要求路由键 “test”，则只有被标记为“test”的消息才被转发，不会转发test.aaa，也不会转发dog.123，只会转发test。
    3：Topic Exchange
        将路由键和某模式进行匹配。此时队列需要绑定要一个模式上。符号“#”匹配一个或多个词，符号“*”匹配不多不少一个词。因此“audit.#”能够匹配到“audit.irs.corporate”，但是“audit.*” 只会匹配到“audit.irs”。
    4： Headers
        Headers类型的exchange使用的比较少，它也是忽略routingKey的一种路由方式。是使用Headers来匹配的。Headers是一个键值对，可以定义成Hashtable。
        发送者在发送的时候定义一些键值对，接收者也可以再绑定时候传入一些键值对，两者匹配的话，则对应的队列就可以收到消息。匹配有两种方式all和any。
        这两种方式是在接收端必须要用键值"x-mactch"来定义。all代表定义的多个键值对都要满足，
        而any则代码只要满足一个就可以了。fanout，direct，topic exchange的routingKey都需要要字符串形式的，而headers exchange则没有这个要求，因为键值对的值可以是任何类型。



四： 常用属性
	1：交换机的属性：
		交换机时还可以附带许多其他的属性，其中最重要的几个分别是：
		1.1： Name 交换机的名称
		1.2： Durability （消息代理重启后，交换机是否还存在）  持久（durable）、暂存（transient）  持久化的交换机会在消息代理（broker）重启后依旧存在，而暂存的交换机则不会（它们需要在代理再次上线后重新被声明）。
		1.3： Auto-delete （当所有与之绑定的消息队列都完成了对此交换机的使用后，删掉它）
		1.4:  Arguments（依赖代理本身）

	2:队列
		AMQP中的队列（queue）跟其他消息队列或任务队列中的队列是很相似的：它们存储着即将被应用消费掉的消息。队列跟交换机共享某些属性，但是队列也有一些另外的属性。
		2.1:  Name 名称
		2.2:  Durable（消息代理重启后，队列依旧存在）
		2.3： Exclusive（只被一个连接（connection）使用，而且当连接关闭后队列即被删除）
		2.4： Auto-delete（当最后一个消费者退订后即被删除）
		2.5： Arguments（一些消息代理用他来完成类似与TTL的某些额外功能）
		队列在声明（declare）后才能被使用。如果一个队列尚不存在，声明一个队列会创建它。如果声明的队列已经存在，并且属性完全相同，那么此次声明不会对原有队列产生任何影响。如果声明中的属性与已存在队列的属性有差异，那么一个错误代码为406的通道级异常就会被抛出。

	3：

*/
