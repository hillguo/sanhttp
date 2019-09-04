package middleware

import (
	"fmt"
	 "github.com/hillguo/sanhttp/ctx"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
)

// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
func Recovery() ctx.HandlerFunc {
	return func(c *ctx.Context) {
		defer func() {
			var rawReq []byte
			if err := recover(); err != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				if c.Request != nil {
					rawReq, _ = httputil.DumpRequest(c.Request, false)
				}
				pl := fmt.Sprintf("http call panic: %s\n%v\n%s\n", string(rawReq), err, buf)
				fmt.Fprintf(os.Stderr, pl)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
