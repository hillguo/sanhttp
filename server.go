// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sanhttp

import (
	"fmt"
	"github.com/hillguo/sanhttp/ctx"
	"github.com/hillguo/sanhttp/errs"
	"github.com/hillguo/sanhttp/middleware"
	"github.com/hillguo/sanhttp/router"
	"github.com/hillguo/sanhttp/utils"
	"log"
	"net"
	"net/http"
	"os"
)

const defaultMultipartMemory = 32 << 20 // 32 MB

var (
	default404Body   = []byte("404 page not found")
	default405Body   = []byte("405 method not allowed")
	defaultAppEngine bool
)

// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
type Engine struct {
	RouterGroup

	trees router.MethodTrees

	// If enabled, the url.RawPath will be used to find parameters.
	UseRawPath bool

	// If true, the path value will be unescaped.
	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,
	// as url.Path gonna be used, which is already unescaped.
	UnescapePathValues bool

	// Value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
	// method call.
	MaxMultipartMemory int64

	allNoRoute  ctx.HandlersChain
	allNoMethod ctx.HandlersChain
	noRoute     ctx.HandlersChain
	noMethod    ctx.HandlersChain
}

var _ IRouter = &Engine{}

// New returns a new blank Engine instance without any middleware attached.
// By default the configuration is:
// - ForwardedByClientIP:    true
// - UseRawPath:             false
// - UnescapePathValues:     true
func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		MaxMultipartMemory: defaultMultipartMemory,
		trees:              make(router.MethodTrees, 0, 9),
	}
	engine.RouterGroup.engine = engine
	return engine
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
	engine := New()
	engine.Use(middleware.Logger(), middleware.Recovery())
	return engine
}

// NoRoute adds handlers for NoRoute. It return a 404 code by default.
func (engine *Engine) NoRoute(handlers ...ctx.HandlerFunc) {
	engine.noRoute = handlers
	engine.rebuild404Handlers()
}

// NoMethod sets the handlers called when... TODO.
func (engine *Engine) NoMethod(handlers ...ctx.HandlerFunc) {
	engine.noMethod = handlers
	engine.rebuild405Handlers()
}

// Use attaches a global middleware to the router. ie. the middleware attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func (engine *Engine) Use(middleware ...ctx.HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}

func (engine *Engine) rebuild404Handlers() {
	engine.allNoRoute = engine.combineHandlers(engine.noRoute)
}

func (engine *Engine) rebuild405Handlers() {
	engine.allNoMethod = engine.combineHandlers(engine.noMethod)
}

func (engine *Engine) addRoute(method, path string, handlers ctx.HandlersChain) {
	root := engine.trees.Get(method)
	if root == nil {
		root = new(router.Node)
		root.FullPath = "/"
		engine.trees = append(engine.trees, router.MethodTree{Method: method, ROOT: root})
	}
	root.AddRoute(path, handlers)
}

// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) Run(addr ...string) (err error) {
	defer func() {}()

	address := utils.ResolveAddress(addr)
	log.Printf("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}

// RunTLS attaches the router to a http.Server and starts listening and serving HTTPS (secure) requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) {
	log.Fatalf("Listening and serving HTTPS on %s\n", addr)
	defer func() { log.Println(err) }()

	err = http.ListenAndServeTLS(addr, certFile, keyFile, engine)
	return
}

// RunUnix attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (ie. a file).
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) RunUnix(file string) (err error) {
	log.Printf("Listening and serving HTTP on unix:/%s", file)
	defer func() { log.Println(err) }()

	os.Remove(file)
	listener, err := net.Listen("unix", file)
	if err != nil {
		return
	}
	defer listener.Close()
	os.Chmod(file, 0777)
	err = http.Serve(listener, engine)
	return
}

// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) RunFd(fd int) (err error) {
	log.Printf("Listening and serving HTTP on fd@%d", fd)
	defer func() { log.Println(err) }()

	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd@%d", fd))
	listener, err := net.FileListener(f)
	if err != nil {
		return
	}
	defer listener.Close()
	err = http.Serve(listener, engine)
	return
}

// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &ctx.Context{
		Index: -1,
		Err:   &errs.Error{},
	}

	c.Request = req
	c.Writer = w

	engine.handleHTTPRequest(c)

}

func (engine *Engine) handleHTTPRequest(c *ctx.Context) {
	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	unescape := false

	// Find root of the tree for the given HTTP method
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].Method != httpMethod {
			continue
		}
		root := t[i].ROOT
		// Find route in tree
		var params ctx.Params
		value := root.GetValue(rPath, params, unescape)
		if value.Handlers != nil {
			c.Handlers = value.Handlers
			c.Params = value.Params
			c.SetFullPath(value.FullPath)
			c.Next()
			return
		}
		break
	}

	c.Handlers = engine.allNoRoute
}
