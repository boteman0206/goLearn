package kafka

/**
一： Apache Kafka是 一个分布式流处理平台. 这到底意味着什么呢?
1: 可以让你发布和订阅流式的记录。这一方面与消息队列或者企业消息系统类似。
2: 可以储存流式的记录，并且有较好的容错性。
3: 可以在流式记录产生时就进行处理。

二： Kafka适合什么样的场景?
1：构造实时流数据管道，它可以在系统或应用之间可靠地获取数据。 (相当于message queue)
2：构建实时流式应用程序，对这些流数据进行转换或者影响。 (就是流处理，通过kafka stream topic和topic之间内部进行变化)


三：Kafka的一些概念
Kafka作为一个集群，运行在一台或者多台服务器上.
Kafka 通过 topic 对存储的流数据进行分类。
每条记录中包含一个key，一个value和一个timestamp（时间戳）。


四： Kafka有四个核心的API:
The Producer API 允许一个应用程序发布一串流式的数据到一个或者多个Kafka topic。
The Consumer API 允许一个应用程序订阅一个或多个 topic ，并且对发布给他们的流式数据进行处理。
The Streams API 允许一个应用程序作为一个流处理器，消费一个或者多个topic产生的输入流，然后生产一个输出流到一个或多个topic中去，在输入输出流中进行有效的转换。
The Connector API 允许构建并运行可重用的生产者或者消费者，将Kafka topics连接到已存在的应用程序或者数据系统。比如，连接到一个关系型数据库，捕捉表（table）的所有变更内容。

五：


*/
