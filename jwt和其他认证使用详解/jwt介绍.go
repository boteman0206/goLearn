package jwt和其他认证使用详解

/**


一： 什么是JWT?
	JWT是JSON Web Token的缩写，它是一种开源标准(RFC 7519)，用来定义通信双方如何安全地交换信息的格式。
	重点：
		JWT之所以叫JSON Web Token，是因为其头部和载荷在编码之前都是JSON格式的数据；
		JWT是一种标准，它有很多的实现方案，比如jwt-auth，专门为php框架laravel打造，java玩家可以看下java-jwt；
		JWT规定以JSON的格式传递信息，负载payload的数据格式是JSON的，通常使用base64编码；
		JWT是自包含的，Token本身携带了验证信息，不需要借助其他工具就可以知道一个Token是否有效，以及载荷信息；
		JWT的某些实现比如黑名单机制、Token刷新等增强功能，可能也需要借助其他工具，但是这并不违背自包含特性。

二：JWT的结构
	JWT 本质上就是一组字串，通过（.）切分成三个为 Base64 编码的部分：
	1： 头部 Header
		头部本身是JSON格式的，注意这里说的是编码之前的格式。头部包括两个字段，token的类型和加密算法。注意这里说的加密算法是签名的加密算法，不是头部的加密算法，也不是载荷的加密算法。实际上头部并没有经过加密，只是通过base64编码成字符串。
		Header 通常由两部分组成：
			1：typ（Type）：令牌类型，也就是 JWT。
			2：alg（Algorithm） ：签名算法，比如 HS256。

	2： 载荷 Payload
		载荷也是JSON格式的，经过base64编码成字符串。上图例子中可以看到有sub,name,iat三个字段，实际上你可以放更多的信息，只要你需要，前提是JSON格式。
		Payload 也是 JSON 格式数据，其中包含了 Claims(声明，包含 JWT 的相关信息)。
			1： Registered Claims（注册声明） ：预定义的一些声明，建议使用，但不是强制性的。
			2： Public Claims（公有声明） ：JWT 签发方可以自定义的声明，但是为了避免冲突，应该在 IANA JSON Web Token Registryopen in new window 中定义它们。
			3： Private Claims（私有声明） ：JWT 签发方因为项目需要而自定义的声明，更符合实际项目场景使用。
			下面是一些常见的注册声明：
				1.1：iss（issuer）：JWT 签发方。
				1.2：iat（issued at time）：JWT 签发时间。
				1.3：sub（subject）：JWT 主题。
				1.4：aud（audience）：JWT 接收方。
				1.5：exp（expiration time）：JWT 的过期时间。
				1.6：nbf（not before time）：JWT 生效时间，早于该定义的时间的 JWT 不能被接受处理。
				1.7：jti（JWT ID）：JWT 唯一标识。
			Payload 部分默认是不加密的，一定不要将隐私信息存放在 Payload 当中！！！
	3: Signature
		HMACSHA256( base64UrlEncode(header) + "." +  base64UrlEncode(payload), secret)

	算出签名以后，把 Header、Payload、Signature 三个部分拼成一个字符串，每个部分之间用"点"（.）分隔，这个字符串就是 JWT 。


三：如何基于 JWT 进行身份验证？
	在基于 JWT 进行身份验证的的应用程序中，服务器通过 Payload、Header 和 Secret(密钥)创建 JWT 并将 JWT 发送给客户端。客户端接收到 JWT 之后，会将其保存在 Cookie 或者 localStorage 里面，以后客户端发出的所有请求都会携带这个令牌。
	两点建议：
		1：建议将 JWT 存放在 localStorage 中，放在 Cookie 中会有 CSRF 风险。
		2：请求服务端并携带 JWT 的常见做法是将其放在 HTTP Header 的 Authorization 字段中（Authorization: Bearer Token）。


四： JWT 的优势
	相比于 Session 认证的方式来说，使用 JWT 进行身份认证主要有下面 4 个优势。
	#无状态
		不过，也正是由于 JWT 的无状态，也导致了它最大的缺点：不可控！
	#有效避免了 CSRF 攻击
		使用 JWT 进行身份验证不需要依赖 Cookie ，因此可以避免 CSRF 攻击。
	#单点登录友好
	#适合移动端应用



五：JWT 身份认证常见问题及解决办法
	#注销登录等场景下 JWT 还有效
		退出登录;
		修改密码;
		服务端修改了某个用户具有的权限或者角色；
		用户的帐户被封禁/删除；
		用户被服务端强制注销；
		用户被踢下线；
		这个问题不存在于 Session 认证方式中，因为在 Session 认证方式中，遇到这种情况的话服务端删除对应的 Session 记录即可。但是，使用 JWT 认证的方式就不好解决了。我们也说过了，JWT 一旦派发出去，如果后端不增加其他逻辑的话，它在失效之前都是有效的。
		解决：
			1、将 JWT 存入内存数据库
			2、黑名单机制
			3、修改密钥 (Secret) :
			4、保持令牌的有效期限短并经常轮换

	#JWT 的续签问题
		1、类似于 Session 认证中的做法
		2、每次请求都返回新 JWT
		3、JWT 有效期设置到半夜
		4、用户登录返回两个 JWT




*/
