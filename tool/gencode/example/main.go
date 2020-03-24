package main

import (
	"github.com/hillguo/sanhttp"
)

func main() {
	app := sanhttp.Default()

	app.Any("/Test/Test", sanhttp.HF(Test))

	app.Any("/Test/GetVidByName", sanhttp.HF(GetVidByName))

	app.Run("0.0.0.0:8888")
}
