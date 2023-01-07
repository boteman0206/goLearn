package kafka基本使用

/**
	windows下： https://blog.csdn.net/qq_40708522/article/details/123842282

   todo： 最新版本的zookeeper是自带的，下载报的时候要注意选择scala的那个
          解压之后运行的时候需要注意bin里面还有个windows的文件夹是保存的.bat命令

    1： 启动zookeeper
  	D:\download\kafka_2.12-3.3.1>.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties
	2：启动kafka
	D:\download\kafka_2.12-3.3.1>.\bin\windows\kafka-server-start.bat .\config\server.properties

	采坑：别忘了打开这个监听端口的配置
		listeners=PLAINTEXT://localhost:9092

	topic命令：（新版本和老版本有区别需要注意）
		创建： kafka-topics.bat --create --topic firsttopic --bootstrap-server localhost:9092
		查询：  kafka-topics.bat --list  --bootstrap-server localhost:9092 // 列出所有的主题
              kafka-topics.bat --describe  --bootstrap-server localhost:9092 --topic firsttopic  // 查看某一个主题的详情

     	删除： --delete  删除可能会报错，需要吧日志文件删除

	生产者：
		控制台开启生产者：
		>.\kafka-console-producer.bat --broker-list localhost:9092 --topic firsttopic
		>hello worls
	消费者：
		>.\kafka-console-consumer.bat --topic firsttopic --bootstrap-server localhost:9092 --from-beginning
		 --from-beginning ： 消费者从最开始进行消费，消费者起启动晚于生产者

*/
