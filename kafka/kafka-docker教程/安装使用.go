package kafka_docker教程

/**
https://www.hangge.com/blog/cache/detail_2791.html  参考文档




创建topic
docker exec kafka kafka-topics.sh --create --zookeeper 172.18.184.38:2181 --replication-factor 1 --partitions 1 --topic test

查看topic
docker exec kafka kafka-topics.sh --list --zookeeper 172.18.184.38:2181

*/
