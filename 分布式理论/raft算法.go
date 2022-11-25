package 分布式理论

/**

https://javaguide.cn/distributed-system/theorem&algorithm&protocol/raft-algorithm.html#_1-1-%E6%8B%9C%E5%8D%A0%E5%BA%AD%E5%B0%86%E5%86%9B
http://thesecretlivesofdata.com/raft/#election  演示图

Raft 源于“Reliable, Replicated, Redundant, And Fault-Tolerant” 是一种用于替代 Paxos 的共识算法。它的目标是确保集群内任意节点的状态保持一致，同时力保整个算法过程易于理解。



一： 拜占庭将军
	假设多位拜占庭将军中没有叛军，信使的信息可靠但有可能被暗杀的情况下，将军们如何达成是否要进攻的一致性决定？
		解决方案大致可以理解成：先在所有的将军中选出一个大将军，用来做出所有的决定。


二： 共识算法
	共识是可容错系统中的一个基本问题：即使面对故障，服务器也可以在共享状态上达成一致。


三： raft节点类型
一个 Raft 集群包括若干服务器，以典型的 5 服务器集群举例。在任意的时间，每个服务器一定会处于以下三个状态中的一个：
	1：Leader：负责发起心跳，响应客户端，创建日志，同步日志。
	2：Candidate：Leader 选举过程中的临时角色，由 Follower 转化而来，发起投票参与竞选。
	3：Follower：接受 Leader 的心跳和日志同步数据，投票给 Candidate。
	在正常的情况下，只有一个服务器是 Leader，剩下的服务器是 Follower。Follower 是被动的，它们不会发送任何请求，只是响应来自 Leader 和 Candidate 的请求。


四：领导人选举
raft 使用心跳机制来触发 Leader 的选举。
	1： 如果一台服务器能够收到来自 Leader 或者 Candidate 的有效信息，那么它会一直保持为 Follower 状态，并且刷新自己的 electionElapsed，重新计时。
	2： Leader 会向所有的 Follower 周期性发送心跳来保证自己的 Leader 地位。如果一个 Follower 在一个周期内没有收到心跳信息，就叫做选举超时，然后它就会认为此时没有可用的 Leader，并且开始进行一次选举以选出一个新的 Leader。
	3： 为了开始新的选举，Follower 会自增自己的 term 号并且转换状态为 Candidate。然后他会向所有节点发起 RequestVoteRPC 请求， Candidate 的状态会持续到以下情况发生：
		赢得选举
			赢得选举的条件是：一个 Candidate 在一个任期内收到了来自集群内的多数选票（N/2+1），就可以成为 Leader。
		其他节点赢得选举
		一轮选举结束，无人胜出
	由于可能同一时刻出现多个 Candidate，导致没有 Candidate 获得大多数选票，如果没有其他手段来重新分配选票的话，那么可能会无限重复下去。
		raft 使用了随机的选举超时时间来避免上述情况。每一个 Candidate 在发起选举后，都会随机化一个新的枚举超时时间，这种机制使得各个服务器能够分散开来，在大多数情况下只有一个服务器会率先超时；它会在其他服务器超时之前赢得选举。# 4 日志复


五：任期号
term： Raft 协议将时间划分成一个个任期（Term），任期用连续的整数表示，每个任期从一次选举开始，赢得选举的节点在该任期内充当 Leader 的职责，随着时间的消逝，集群可能会发生新的选举，任期号也会单调递增。
	    	通过任期号，可以比较各个节点的数据新旧、识别过期的 Leader 等，它在 Raft 算法中充当逻辑时钟，发挥着重要作用。

六： 日志
	1：entry：每一个事件成为 entry，只有 Leader 可以创建 entry。entry 的内容为<term,index,cmd>其中 cmd 是可以应用到状态机的操作。
	2：log：由 entry 构成的数组，每一个 entry 都有一个表明自己在 log 中的 index。只有 Leader 才可以改变其他节点的 log。
		entry 总是先被 Leader 添加到自己的 log 数组中，然后再发起共识请求，获得同意后才会被 Leader 提交给状态机。Follower 只能从 Leader 获取新日志和当前的 commitIndex，然后把对应的 entry 应用到自己的状态机中。


六：选举时机
	1：一种就是当 Raft 集群启动之后开始第一次选举，
	2：另外一种就是 Leader 出现故障，无法使用心跳机制维持自己的统治， 导致选举超时机制触发，节点开始尝试新的一轮的选举(需要注意的是，
		选举有可能在一个 Term 中没有结果，那么就在下个 Term 中继续选举直到选出 Leader)。


七： 选举过程
	1：如果追随者没有收到领导的消息，那么他们就可以成为候选人。
	2：然后候选人从其他节点请求选票。
	3：节点将用他们的投票进行回复。如果候选人获得了大多数节点的选票，就会成为领先者。这个过程被称为领导人选举。
	4：成为leader之后进行日志复制


八：consul的raft算法实现








*/
