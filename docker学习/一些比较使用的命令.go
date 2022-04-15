package docker学习

/**

按照镜像名称删除容器：
	 docker rm $(docker ps -aq  --filter ancestor=training/webapp)   // training/webapp就是镜像的名称

*/
