package 高阶知识点

/**
参考书籍 https://share.weiyun.com/5z6Ow6i

一：消息何去何从?
	1：mandatory和immediate是push发送消息方法中的两个参数，他们都有当消息传递过程中不可达目的地时将消息返回给生产者的功能。
		mandatory设为true时，交换器无法根据自身的类型和路由键找到一个符合条件的队列，那么mq会调用basic.return命令将消息返回给生产者
		mandatory设为false时,出先以上情况直接丢弃

		immediate为true时，如果交换器在将消息路由到队列时候并不存在任何消费者，消息不会假日队列，该消息会通过basic.Return返回给生产者
		immediate为false时，则丢弃消息

	2：备份交换机
		rabbitMq提供的备份交换器将可以将未能被交换器路由的消息（没有绑定队列或者没有匹配的绑定）存储起来，而不用返回给客户端
		不想使用mandatory参数，那样需要添加returnListener的编辑逻辑，生产者代码将变得复杂，这时候可以使用备份交换器，需要在申明交换器的时候添加alternate-exchange参数来实现，也可以通过策略的policy的方式实现，
		两者同时使用的时候，前者优先级更高，会覆盖掉policy的设置

二：过期时间TTL
	1： 设置消息的ttl
		mq只会判断队列的头消息是否过期，所有即使消息过期，如果不在头部也不会立即丢弃，因为要扫描所有的队列消息，性能损耗大
	2：设置队列的ttl（ttl超时会进入死信队列）
		申明队列的时候添加参数： x-message-ttl
		如果不设置ttl标识消息不过期，如果设置ttl=0则表示此时可以直接将消息投递到消费者，否则直接将消息丢弃


三： 死信队列
	1;消息被拒绝
	2：消息过期
	3;消息队列达到最大长度
	通过参数x-dead-letter-exchange来设置死信交换机，x-dead-letter-routing-key 指定路由键



四：延迟队列
	1：基于死信队列和ttl实现的延迟队列
	2; 基于插件实现的


五：优先级队列
	具有高优先级的队列具有很高的优先权，优先级高的消息具备优先被消费的特权。可以通过x-max-priority参实现
	1： 声明队列时增加一个参数
		// 官方允许是 0-255 之间 此处设置10 允许优化级范围0-10 不要设置过大 浪费CPU与内存
		arguments.put("x-max-priority", 10);

	2：发布消息时设置优先级 不能高于声明队列时设置的参数
		// 设置优先级, 不得高于 x-max-priority 设置的值
		AMQP.BasicProperties basicProperties = new AMQP.BasicProperties().builder().priority(5).build();



*/
