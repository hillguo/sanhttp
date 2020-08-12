package main

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/hillguo/sanhttp/tool/gencode"
)

var maintmplStr = `package main

import (
	"github.com/hillguo/sanhttp"
)

func main() {
	app := sanhttp.Default()
	{{range .Methods}}
	app.Any("/{{$.ServiceName}}/{{.MethodName}}", sanhttp.HF({{.MethodName}}))
	{{end}}
	app.Run("0.0.0.0:8888")
}
`

var handlertplStr = `package main

import (
	"github.com/hillguo/sanhttp/ctx"
)

func {{.MethodName}}(c *ctx.Context, req *{{.InputType}}, resp *{{.OutputType}}) error {

	return nil
}

`

func genmain(protoInfo *gencode.ProtoFileInfo) []gencode.FileNameData {
	basicTmpl, err := template.New("basic").Parse(maintmplStr) //创建、解析
	if err != nil {
		panic(err)
	}

	flienamedata := make([]gencode.FileNameData, 0, len(protoInfo.Methods))
	var buf bytes.Buffer
	basicTmpl.Execute(&buf, protoInfo)
	flienamedata = append(flienamedata, gencode.FileNameData{Name: "main.go", Data: buf.String()})

	// handlers

	handlerTmpl, err := template.New("basic").Parse(handlertplStr) //创建、解析
	if err != nil {
		panic(err)
	}

	for _, v := range protoInfo.Methods {
		buf.Reset()
		handlerTmpl.Execute(&buf, v)
		flienamedata = append(flienamedata, gencode.FileNameData{Name:"handler_"+ strings.ToLower(v.MethodName) + ".go", Data: buf.String()})
	}
	return flienamedata
}

func main() {
	gencode.Main(genmain)
}
