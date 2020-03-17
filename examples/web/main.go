package main

import (
	"github.com/hillguo/sanhttp"
	"github.com/hillguo/sanhttp/ctx"
)

func main() {
	app := sanhttp.Default()
	app.GET("/", echo)
	app.Any("/test", sanhttp.HF(Test))
	app.Run("0.0.0.0:1234")
}

func echo(c *ctx.Context) {
	c.JSON(200, 123)
}



func Test(c *ctx.Context, req *TestReq, resp *TestResp) error {
	resp.C = "woaini"
	return nil
}
