package geb

import (
	"fmt"
	"net/http"
)

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func getKey(m, p string) string {
	return fmt.Sprintf("%s-%s", m, p)
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	key := getKey(method, pattern)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := getKey(c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%s\n", c.Path)
	}

}
