package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
//type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
//type Engine struct {
//	router map[string]HandlerFunc
//}
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	//return &Engine{router: make(map[string]HandlerFunc)}
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//key := method + "-" + pattern
	//engine.router[key] = handler
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//key := req.Method + "-" + req.URL.Path
	//if handler, ok := engine.router[key]; ok {
	//	handler(w, req)
	//} else {
	//	w.WriteHeader(http.StatusNotFound)
	//	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	//}
	c := newContext(w, req)
	engine.router.handle(c)

}

/**
首先定义了类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法。
我们在Engine中，添加了一张路由映射表router，key 由请求方法和静态路由地址构成，
例如GET-/、GET-/hello、POST-/hello，这样针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。

当用户调用(*Engine).GET()方法时，会将路由和处理方法注册到映射表 router 中，(*Engine).Run()方法，是 ListenAndServe 的包装。

Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
*/
