package route

import (
	"strings"

	"github.com/hillguo/sanhttp/ctx"
)

type PathHandlers map[string]ctx.HandlersChain

var DefaultRoute = &Route{routes: map[string]PathHandlers{}}

type Route struct {
	routes map[string]PathHandlers
}

func (r *Route) AddHandler(method, path string, handlers ctx.HandlersChain) {
	method = strings.ToLower(method)
	path = strings.ToLower(path)
	if r == nil {
		r = &Route{}
	}

	if r.routes == nil {
		r.routes = make(map[string]PathHandlers, 6)
	}
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = make(map[string]ctx.HandlersChain)
	}

	if pathHandler, ok := r.routes[method]; ok && pathHandler != nil {
		pathHandler[path] = handlers
	}
}

func (r *Route) GetHandler(method, path string) ctx.HandlersChain {
	method = strings.ToLower(method)
	path = strings.ToLower(path)
	if pathHandler, ok := r.routes[method]; ok {
		return pathHandler[path]
	}
	return nil
}
