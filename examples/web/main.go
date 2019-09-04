package main

import (
	"github.com/hillguo/sanhttp"
	"github.com/hillguo/sanhttp/ctx"
)

func main(){
	app := sanhttp.Default()
	app.GET("/",echo)
	app.Run("127.0.0.1:1234")
}


func echo(c *ctx.Context){
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte("123"))
}