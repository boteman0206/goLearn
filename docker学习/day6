package docker学习


容器的互联和创建网络
docker network ls : 查看现有逇网络


1: 可以使用 --link 来指定连接的容器名称实现互联 可以通过name去ping,底层就是修改/etc/hosts文件




2: 通过自定义的网络来实现
docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet




启动自定义的网络
docker run -d -P --name tomcat-01-net --net mynet tomcat
docker run -d -P --name tomcat-02-net --net mynet tomcat

docker network inspect mynet  可以查看详情


docker exec -it tomcat-01-net ping tomcat-02-net  可以直接通过name去ping