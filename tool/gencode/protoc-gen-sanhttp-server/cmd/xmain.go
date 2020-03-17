package main

import (
	"github.com/hillguo/sanhttp"
)

func main() {
	app := sanhttp.Default()

	app.Any("/test/test", TestCGIFactoryHandler)

	app.Any("/test/getvidbyname", GetVidByNameCGIFactoryHandler)

	app.Run("0.0.0.0:8080")
}
