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

*/
