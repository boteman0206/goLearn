package main

import "golang.org/x/sync/errgroup"

var err errgroup.Group

func main() {

}

/**

errgroup 的使用

sync.errgroup 的由来

当我们的 goroutine 出错了，我们怎么知道出错的原因呢？

你有一组任务是并发的工作，当遇到某种错误或者你不想再输出了，你可能想取消整个 goroutine 那么 errgroup.errgroup 就是为你而设计的。

sync.errgroup 的示例

sync.errgroup 源码阅读
https://github.com/golang/sync/blob/master/errgroup/errgroup.go

sync.errgroup 的延展
https://github.com/bilibili/kratos/blob/master/pkg/sync/errgroup/errgroup.go
提供带 recover 和并行数的 errgroup，err 中包含详细堆栈信息。

sync.errgroup 开源项目使用情况
① 基于 errgroup 的封装 https://godoc.org/github.com/cockroachdb/cockroach/pkg/util/ctxgroup ② https://godoc.org/github.com/coreos/etcdlabs/cluster
③ https://godoc.org/github.com/docker/containerd/fs
④ https://godoc.org/github.com/grafana/grafana/pkg/tsdb/cloudwatch
⑤ https://godoc.org/github.com/knative/serving/pkg/pool
⑥ https://godoc.org/github.com/kubernetes/minikube/pkg/minikube/bootstrapper/kubeadm

Q&A
*/
