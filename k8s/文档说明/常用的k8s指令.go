package 文档说明

/**

创建更新pod资源
kubectl apply -f deployment.yaml


pod映射本机端口
kubectl port-forward hellok8s-deployment-694687d76f-gnnk2   3000:3000

会滚上一个版本
kubectl rollout undo deployment hellok8s-deployment

获取pod信息
kubectl get pods
# NAME                                   READY   STATUS    RESTARTS   AGE
# hellok8s-deployment-77bffb88c5-cvm5c   1/1     Running   0          39s
# hellok8s-deployment-77bffb88c5-lktbl   1/1     Running   0          41s
# hellok8s-deployment-77bffb88c5-nh82z   1/1     Running   0          37s

查看pod描述信息
kubectl describe pod hellok8s-deployment-77bffb88c5-cvm5c
# Image: guangzhengli/hellok8s:v1


查看历史版本 除了上面的命令，还可以用 history 来查看历史版本，--to-revision=2 来回滚到指定版本。
kubectl rollout history deployment hellok8s-deployment
kubectl rollout undo deployment/hellok8s-deployment --to-revision=2


来观察 pod 的创建销毁情况
kubectl get pods --watch


存活探针 (livenessProb)
存活探测器来确定什么时候要重启容器。 例如，存活探测器可以探测到应用死锁（应用程序在运行，但是无法继续执行后面的步骤）情况。
重启这种状态下的容器有助于提高应用的可用性，即使其中存在缺陷。-- LivenessProb



显示资源列表
	# kubectl get 资源类型
	#获取类型为Deployment的资源列表
	kubectl get deployments

	#获取类型为Pod的资源列表
	kubectl get pods

	#获取类型为Node的资源列表
	kubectl get nodes

名称空间
在命令后增加 -A 或 --all-namespaces 可查看所有 名称空间中 的对象，使用参数 -n 可查看指定名称空间的对象，例如
	# 查看所有名称空间的 Deployment
	kubectl get deployments -A
	kubectl get deployments --all-namespaces
	# 查看 kube-system 名称空间的 Deployment
	kubectl get deployments -n kube-system

kubectl describe - 显示有关资源的详细信息
	# kubectl describe 资源类型 资源名称
	#查看名称为nginx-XXXXXX的Pod的信息
	kubectl describe pod nginx-XXXXXX
	#查看名称为nginx的Deployment的信息
	kubectl describe deployment nginx

kubectl logs - 查看pod中的容器的打印日志（和命令docker logs 类似）
	# kubectl logs Pod名称
	#查看名称为nginx-pod-XXXXXXX的Pod内的容器打印的日志
	#本案例中的 nginx-pod 没有输出日志，所以您看到的结果是空的
	kubectl logs -f nginx-pod-XXXXXXX

kubectl exec - 在pod中的容器环境内执行命令(和命令docker exec 类似)
	# kubectl exec Pod名称 操作命令
	# 在名称为nginx-pod-xxxxxx的Pod中运行bash
	kubectl exec -it nginx-pod-xxxxxx /bin/bash



查看指定名称空间下的pods信息
注意： 当您使用 kubectl get pods 命令而没有显式指定命名空间时，它会默认查询当前上下文下的默认命名空间
kubectl get pods -n <namespace>
kubectl get pods -n my-namespace -o wide
kubectl describe pod <pod-name> -n <namespace>





kubernetes 提供了一种名叫 Service 的资源帮助解决这些问题，它为 pod 提供一个稳定的 Endpoint。Service 位于 pod 的前面，负责接收请求并将它们传递给它后面的所有pod。一旦服务中的 Pod 集合发生更改，Endpoints 就会被更新，请求的重定向自然也会导向最新的 pod。
service设置：
	1: 先需要编写一个service.yaml的文件，kubectl apply -f 启动对应的资源
	2: 被selector选中的app会被绑定起来
		kubectl get endpoints
		# NAME                         ENDPOINTS                                          AGE
		# service-hellok8s-clusterip   172.17.0.10:3000,172.17.0.2:3000,172.17.0.3:3000   10s

	3: kubectl get pod -o wide
	4: kubectl get service


*/
