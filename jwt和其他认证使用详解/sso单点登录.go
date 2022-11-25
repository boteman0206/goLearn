package jwt和其他认证使用详解

/**

https://juejin.cn/post/7044328327762411534


一： 什么是 SSO？
	SSO 英文全称 Single Sign On，单点登录。SSO 是在多个应用系统中，用户只需要登录一次就可以访问所有相互信任的应用系统。
	例如你登录网易账号中心（https://reg.163.com/ ）之后访问以下站点都是登录状态。
		网易直播 https://v.163.comopen in new window
		网易博客 https://blog.163.comopen in new window
		网易花田 https://love.163.comopen in new window
		网易考拉 https://www.kaola.comopen in new window
		网易 Lofter http://www.lofter.comopen in new window


二、早期的多系统登录解决方案
	单系统登录解决方案的核心是cookie，cookie携带会话id在浏览器与服务器之间维护会话状态。但cookie是有限制的，这个限制就是cookie的域（通常对应网站的域名），浏览器发送http请求时会自动携带与该域匹配的cookie，而不是所有cookie
	解决：
		既然这样，为什么不将web应用群中所有子系统的域名统一在一个顶级域名下，例如“*.baidu.com”，然后将它们的cookie域设置为“baidu.com”，这种做法理论上是可以的，
		甚至早期很多多系统登录就采用这种同域名共享cookie的方式。
	共享cookie的方式存在众多局限：
		1: 应用群域名得统一
		2: 应用群各系统使用的技术（至少是web服务器）要相同，不然cookie的key值（tomcat为JSESSIONID）不同，无法维持会话，共享cookie的方式是无法实现跨语言技术平台登录的，比如java、php、.net系统之间
		3: cookie本身不安全。




*/
