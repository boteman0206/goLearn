package git

/**


有时候会莫名其妙的出现一些拉取不到github的问题，可能是开代理把github的端口改了

可以在  .ssh/目录下新增config文件 配置代理端口和github的地址端口 20.205.243.166 22，可以解决

需要windows客户端配合开启代理


Host github.com
        User git
        ProxyCommand connect -H 127.0.0.1:9910 20.205.243.166 22



*/
