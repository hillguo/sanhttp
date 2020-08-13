// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sanhttp

import (
	"net/http"

	log "github.com/hillguo/sanlog"

	"github.com/hillguo/sanhttp/ctx"
	"github.com/hillguo/sanhttp/errs"
	"github.com/hillguo/sanhttp/middleware"
	route "github.com/hillguo/sanhttp/router"
)

var (
	default404Body = "404 page not found"
)

func noRoute(c *ctx.Context) {
	c.String(404, default404Body)
}

// HTTPServer server is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of server, by using New() or Default()
type HTTPServer struct {
	RouterGroup

	route *route.Route

	noRoute  ctx.HandlersChain
	noMethod ctx.HandlersChain
}

// New ...
func New() *HTTPServer {
	server := &HTTPServer{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		route:   route.DefaultRoute,
		noRoute: ctx.HandlersChain{noRoute},
	}
	server.RouterGroup.server = server
	return server
}

// Default returns an server instance with the Logger and Recovery middleware already attached.
func Default() *HTTPServer {
	server := New()
	server.Use(middleware.Logger(), middleware.Recovery())
	return server
}

// NoRoute adds handlers for NoRoute. It return a 404 code by default.
func (server *HTTPServer) NoRoute(handlers ...ctx.HandlerFunc) {
	server.noRoute = handlers
	server.rebuild404Handlers()
}

// NoMethod sets the handlers called when... TODO.
func (server *HTTPServer) NoMethod(handlers ...ctx.HandlerFunc) {
	server.noMethod = handlers
	server.rebuild405Handlers()
}

// Use attaches a global middleware to the router. ie. the middleware attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func (server *HTTPServer) Use(middleware ...ctx.HandlerFunc) IRoutes {
	server.RouterGroup.Use(middleware...)
	server.rebuild404Handlers()
	server.rebuild405Handlers()
	return server
}

func (server *HTTPServer) rebuild404Handlers() {
	server.noRoute = server.combineHandlers(server.noRoute)
}

func (server *HTTPServer) rebuild405Handlers() {
	server.noMethod = server.combineHandlers(server.noMethod)
}

func (server *HTTPServer) addRoute(method, path string, handlers ctx.HandlersChain) {
	server.route.AddHandler(method, path, handlers)
}

// Run ...
func (server *HTTPServer) Run(addr string) (err error) {
	defer func() {}()

	log.Infof("Listening and serving HTTP on %s\n", addr)
	err = http.ListenAndServe(addr, server)
	return
}

// RunTLS attaches the router to a http.Server and starts listening and serving HTTPS (secure) requests.
func (server *HTTPServer) RunTLS(addr, certFile, keyFile string) (err error) {
	log.Fatalf("Listening and serving HTTPS on %s\n", addr)
	defer func() { log.Info(err) }()

	err = http.ListenAndServeTLS(addr, certFile, keyFile, server)
	return
}

// ServeHTTP conforms to the http.Handler interface.
func (server *HTTPServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	httpMethod := req.Method
	rPath := req.URL.Path
	handlers := server.route.GetHandler(httpMethod, rPath)
	if handlers == nil {
		handlers = server.noRoute
	}

	c := &ctx.Context{
		Index:    -1,
		Err:      &errs.Error{},
		Request:  req,
		Writer:   w,
		Handlers: handlers,
	}
	c.Next()
}
