proxyClient.go tcp模式, 去掉头部直连服务器; http模式, 直连http代理
tunnel.go 加解密隧道, 包含客户端和服务器, 客户端加密, 服务端解密
tcpClient.go, tcpServer.go 自己用来测试的小例子



tunnel 会话， tcp共享会话

1, 客户端发起TCP连接

2, 客户端发起建立加密隧道数据
加密算法类型
客户端签名（用户名，公钥）

3, 服务器回复确认加密隧道数据
用客户端公钥加密的随机密码
服务端签名（用户名，公钥）

4, 客户端
使用客户端私钥解密密码
使用流加密加密数据
传输数据



常见问题
1, socket: too many open files

查看链接限制 ulimit -a

sysctl -w fs.file-max=2000500
sysctl -w fs.nr_open=2000500
sysctl -w net.nf_conntrack_max=2000500
ulimit -n 2000500

sysctl -w net.ipv4.tcp_tw_recycle=1
sysctl -w net.ipv4.tcp_tw_reuse=1

2, connect: cannot assign requested address

一个客户端占用的端口数量为65535个， 解决办法是客户端使用docker, 这样可以产生n个65535,
服务器只用了一个端口

作为一个tcp长链接的隧道， 最多支持65534个客户端同时在线， 65534个客户端同时链接一个服务端口，服务器再链接65534个不同的目标服务器

有一个问题是“一台机器可以建立多少个连接”， 看似与我们遇到的问题一样，其实不一样。
多少连接和多少端口是不一样的， 一个服务器的连接数受内存影响很大，支持百万连接没有什么问题，却只使用一个端口
多少端口是系统定死的， “一台机器可以连接多少个服务器”？

要想tcp隧道支持百万连接，那必然是分布式的，前端服务器的一个端口负责百万连接， 后端n的服务器负责连接目标服务器（同一机器上部署10个docker, 可支持60万长连接）



