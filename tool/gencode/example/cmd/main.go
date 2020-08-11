package main

import (
	"github.com/hillguo/sanhttp"
	"github.com/hillguo/sanhttp/tool/gencode/example"
)

func main() {
	app := sanhttp.Default()

	app.Any("/Test/Test", sanhttp.HF(example.Test))

	app.Any("/Test/GetVidByName", sanhttp.HF(example.GetVidByName))

	app.Run("0.0.0.0:8888")
}
