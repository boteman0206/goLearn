package http和tcp

/**


1: TCP 四元组可以唯一的确定一个连接，四元组包括如下：
	源地址
	源端口
	目的地址
	目的端口
2： 什么是 TCP 连接？  (面向链接的，可靠的，字节流)
	简单来说就是，用于保证可靠性和流量控制维护的某些状态信息，这些信息的组合，包括Socket、序列号和窗口大小称为连接。
	所以我们可以知道，建立一个 TCP 连接是需要客户端与服务端端达成上述三个信息的共识。
	Socket：由 IP 地址和端口号组成
	序列号：用来解决乱序问题等
	窗口大小：用来做流量控制


2: 有一个 IP 的服务端监听了一个端口，它的 TCP 的最大连接数是多少？
	对 IPv4，客户端的 IP 数最多为 2 的 32 次方，客户端的端口数最多为 2 的 16 次方，也就是服务端单机最大 TCP 连接数，约为 2 的 48 次方。
	当然，服务端最大并发 TCP 连接数远不能达到理论上限，会受以下因素影响：
		文件描述符限制，每个 TCP 连接都是一个文件，如果文件描述符被占满了，会发生 too many open files。Linux 对可打开的文件描述符的数量分别作了三个方面的限制：
		系统级：当前系统可打开的最大数量，通过 cat /proc/sys/fs/file-max 查看；
		用户级：指定用户可打开的最大数量，通过 cat /etc/security/limits.conf 查看；
		进程级：单个进程可打开的最大数量，通过 cat /proc/sys/fs/nr_open 查看；
		内存限制，每个 TCP 连接都要占用一定内存，操作系统的内存是有限的，如果内存资源被占满后，会发生 OOM
3: UDP 和 TCP 有什么区别呢？分别的应用场景是？
	3.1. 连接
		TCP 是面向连接的传输层协议，传输数据前先要建立连接。
		UDP 是不需要连接，即刻传输数据。
	3.2. 服务对象
		TCP 是一对一的两点服务，即一条连接只有两个端点。
		UDP 支持一对一、一对多、多对多的交互通信
	3.3. 可靠性
		TCP 是可靠交付数据的，数据可以无差错、不丢失、不重复、按序到达。
		UDP 是尽最大努力交付，不保证可靠交付数据。但是我们可以基于 UDP 传输协议实现一个可靠的传输协议，比如 QUIC 协议
	3.4. 拥塞控制、流量控制
		TCP 有拥塞控制和流量控制机制，保证数据传输的安全性。
		UDP 则没有，即使网络非常拥堵了，也不会影响 UDP 的发送速率
	3.5. 首部开销
		TCP 首部长度较长，会有一定的开销，首部在没有使用「选项」字段时是 20 个字节，如果使用了「选项」字段则会变长的。
		UDP 首部只有 8 个字节，并且是固定不变的，开销较小。
	3.6. 传输方式
		TCP 是流式传输，没有边界，但保证顺序和可靠。
		UDP 是一个包一个包的发送，是有边界的，但可能会丢包和乱序。
	3.7. 分片不同
		TCP 的数据大小如果大于 MSS 大小，则会在传输层进行分片，目标主机收到后，也同样在传输层组装 TCP 数据包，如果中途丢失了一个分片，只需要传输丢失的这个分片。
		UDP 的数据大小如果大于 MTU 大小，则会在 IP 层进行分片，目标主机收到后，在 IP 层组装完数据，接着再传给传输层。



4：TCP 和 UDP 可以使用同一个端口吗？
答案：可以的。 TCP/UDP 各自的端口号也相互独立，如 TCP 有一个 80 号端口，UDP 也可以有一个 80 号端口，二者并不冲突。

5：TCP 三次握手过程是怎样的？
	5.1： 一开始，客户端和服务端都处于 CLOSE 状态。先是服务端主动监听某个端口，处于 LISTEN 状态
	5.2：客户端会随机初始化序号（client_isn），将此序号置于 TCP 首部的「序号」字段中，同时把 SYN 标志位置为 1 ，表示 SYN 报文。接着把第一个 SYN 报文发送给服务端，表示向服务端发起连接，该报文不包含应用层数据，之后客户端处于 SYN-SENT 状态。
	5.3：服务端收到客户端的 SYN 报文后，首先服务端也随机初始化自己的序号（server_isn），将此序号填入 TCP 首部的「序号」字段中，其次把 TCP 首部的「确认应答号」字段填入 client_isn + 1, 接着把 SYN 和 ACK 标志位置为 1。最后把该报文发给客户端，该报文也不包含应用层数据，之后服务端处于 SYN-RCVD 状态。
	5.4： 客户端收到服务端报文后，还要向服务端回应最后一个应答报文，首先该应答报文 TCP 首部 ACK 标志位置为 1 ，其次「确认应答号」字段填入 server_isn + 1 ，最后把报文发送给服务端，这次报文可以携带客户到服务端的数据，之后客户端处于 ESTABLISHED 状态。
	5.5： 服务端收到客户端的应答报文后，也进入 ESTABLISHED 状态。

6： 为什么是三次握手？不是两次、四次？
	6.1：原因一：避免历史连接  三次握手才可以阻止重复历史连接的初始化（主要原因）

	6.2： 原因二：同步双方初始序列号
		TCP 协议的通信双方， 都必须维护一个「序列号」， 序列号是可靠传输的一个关键因素，它的作用：
		接收方可以去除重复的数据；
		接收方可以根据数据包的序列号按序接收；
		可以标识发送出去的数据包中， 哪些是已经被对方收到的（通过 ACK 报文中的序列号知道）；

	6.3： 三次握手才可以避免资源浪费
		如果只有「两次握手」，当客户端发生的 SYN 报文在网络中阻塞，客户端没有接收到 ACK 报文，就会重新发送 SYN ，由于没有第三次握手，服务端不清楚客户端是否收到了自己回复的 ACK 报文，所以服务端每收到一个 SYN 就只能先主动建立一个连接，
		如果客户端发送的 SYN 报文在网络中阻塞了，重复发送多次 SYN 报文，那么服务端在收到请求后就会建立多个冗余的无效链接，造成不必要的资源浪费。

	总结：不使用「两次握手」和「四次握手」的原因：
	「两次握手」：无法防止历史连接的建立，会造成双方资源的浪费，也无法可靠的同步双方序列号；
	「四次握手」：三次握手就已经理论上最少可靠连接建立，所以不需要使用更多的通信次数。

7： 为什么每次建立 TCP 连接时，初始化的序列号都要求不一样呢？
	为了防止历史报文被下一个相同四元组的连接接收（主要方面）；
	为了安全性，防止黑客伪造的相同序列号的 TCP 报文被对方接收；




8： TCP 四次挥手
	具体过程：
	客户端主动调用关闭连接的函数，于是就会发送 FIN 报文，这个 FIN 报文代表客户端不会再发送数据了，进入 FIN_WAIT_1 状态；
	服务端收到了 FIN 报文，然后马上回复一个 ACK 确认报文，此时服务端进入 CLOSE_WAIT 状态。在收到 FIN 报文的时候，TCP 协议栈会为 FIN 包插入一个文件结束符 EOF 到接收缓冲区中，服务端应用程序可以通过 read 调用来感知这个 FIN 包，这个 EOF 会被放在已排队等候的其他已接收的数据之后，所以必须要得继续 read 接收缓冲区已接收的数据；
	接着，当服务端在 read 数据的时候，最后自然就会读到 EOF，接着 read() 就会返回 0，这时服务端应用程序如果有数据要发送的话，就发完数据后才调用关闭连接的函数，如果服务端应用程序没有数据要发送的话，可以直接调用关闭连接的函数，这时服务端就会发一个 FIN 包，这个 FIN 报文代表服务端不会再发送数据了，之后处于 LAST_ACK 状态；
	客户端接收到服务端的 FIN 包，并发送 ACK 确认包给服务端，此时客户端将进入 TIME_WAIT 状态；
	服务端收到 ACK 确认包后，就进入了最后的 CLOSE 状态；
	客户端经过 2MSL 时间之后，也进入 CLOSE 状态；
	你可以看到，每个方向都需要一个 FIN 和一个 ACK，因此通常被称为四次挥手
9： TCP 四次挥手，可以变成三次吗？
	当被动关闭方（上图的服务端）在 TCP 挥手过程中，「没有数据要发送」并且「开启了 TCP 延迟确认机制」，那么第二和第三次挥手就会合并传输，这样就出现了三次挥手。
10： TCP 序列号和确认号是如何变化的？



11： TCP 序列号和确认号是如何变化的
公式一：序列号 = 上一次发送的序列号 + len（数据长度）。特殊情况，如果上一次发送的报文是 SYN 报文或者 FIN 报文，则改为 上一次发送的序列号 + 1。
公式二：确认号 = 上一次收到的报文中的序列号 + len（数据长度）。特殊情况，如果收到的是 SYN 报文或者 FIN 报文，则改为上一次收到的报文中的序列号 + 1。


问题
1：服务端大量处于 TIME_WAIT 状态连接的原因。
	TIME_WAIT 状态是「主动关闭连接方」才会出现的状态。而且 TIME_WAIT 状态会持续 2MSL 时间才会进入到 close 状态。在 Linux 上 2MSL 的时长是 60 秒，也就是说停留在 TIME_WAIT 的时间为固定的 60 秒。

	1.1： 为什么需要 TIME_WAIT 状态？
		一： 保证「被动关闭连接」的一方，能被正确的关闭。
			TCP 协议在关闭连接的四次挥手中，在主动关闭方发送的最后一个 ACK 报文，有可能丢失，这时被动方会重新发 FIN 报文, 如果这时主动方处于 CLOSE 状态 ，就会响应 RST 报文而不是 ACK 报文。所以主动方要处于 TIME_WAIT 状态，而不能是 CLOSE。
		二： 防止历史连接中的数据，被后面相同四元组的连接错误的接收。
			为了避免这个情 况， TIME_WAIT 状态需要持续 2MSL，因为这样就可以保证当成功建立一个 TCP 连接的时候，来自连接先前化身的重复报文已经在网络中消逝。
    谁先关闭连接的，它就是主动关闭方，那么 TIME_WAIT 就会出现在主动关闭方。
	解决答案：如果服务端出现大量的 TIME_WAIT 状态的 TCP 连接，就是说明服务端主动断开了很多 TCP 连接。
		第一个场景：HTTP 没有使用长连接
			当服务端出现大量的 TIME_WAIT 状态连接的时候，可以排查下是否客户端和服务端都开启了 HTTP Keep-Alive，因为任意一方没有开启  HTTP Keep-Alive，都会导致服务端在处理完一个 HTTP 请求后，就主动关闭连接，此时服务端上就会出现大量的 TIME_WAIT 状态的连接。
		第二个场景：HTTP 长连接超时
		第三个场景：HTTP 长连接的请求数量达到上限
	过多的 TIME-WAIT 状态主要的危害有两种：
		第一是占用系统资源，比如文件描述符、内存资源、CPU 资源等；
		第二是占用端口资源，端口资源也是有限的，一般可以开启的端口为 32768～61000，也可以通过 net.ipv4.ip_local_port_range参数指定范围



2：服务端大量处于 CLOSE_WAIT 状态连接的原因。
	当服务端出现大量 CLOSE_WAIT 状态的连接的时候，说明服务端的程序没有调用 close 函数关闭连接。
	答案： https://mp.weixin.qq.com/s/sK2caRVxmkXInKcxtDsTVg
		第一个原因：第 2 步没有做，没有将服务端 socket 注册到 epoll，这样有新连接到来时，服务端没办法感知这个事件，也就无法获取到已连接的 socket，那服务端自然就没机会对 socket 调用 close 函数了。不过这种原因发生的概率比较小，这种属于明显的代码逻辑 bug，在前期 read view 阶段就能发现的了。
		第二个原因：第 3 步没有做，有新连接到来时没有调用 accpet 获取该连接的 socket，导致当有大量的客户端主动断开了连接，而服务端没机会对这些 socket 调用 close 函数，从而导致服务端出现大量 CLOSE_WAIT 状态的连接。		发生这种情况可能是因为服务端在执行 accpet  函数之前，代码卡在某一个逻辑或者提前抛出了异常。
		第三个原因：第 4 步没有做，通过 accpet 获取已连接的 socket 后，没有将其注册到 epoll，导致后续收到 FIN 报文的时候，服务端没办法感知这个事件，那服务端就没机会调用 close 函数了。发生这种情况可能是因为服务端在将已连接的 socket 注册到 epoll 之前，代码卡在某一个逻辑或者提前抛出了异常。之前看到过别人解决 close_wait 问题的实践文章，感兴趣的可以看看：一次 Netty 代码不健壮导致的大量 CLOSE_WAIT 连接原因分析
		第四个原因：第 6 步没有做，当发现客户端关闭连接后，服务端没有执行 close 函数，可能是因为代码漏处理，或者是在执行 close 函数之前，代码卡在某一个逻辑，比如发生死锁等等。可以发现，当服务端出现大量 CLOSE_WAIT 状态的连接的时候，通常都是代码的问题，这时候我们需要针对具体的代码一步一步的进行排查和定位，主要分析的方向就是服务端为什么没有调用 close。

*/
