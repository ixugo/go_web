package geb

import (
	"net/http"
)

// HandleFunc 请求处理
type HandleFunc func(*Context)

// Engine 实现 http.Handle
type Engine struct {
	router *router
}

// New Engine 构造函数
func New() *Engine {
	return &Engine{newRouter()}
}

// GET method
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.router.addRoute(http.MethodGet, pattern, handler)
}

// POST method
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.router.addRoute(http.MethodPost, pattern, handler)
}

// Run start server
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
