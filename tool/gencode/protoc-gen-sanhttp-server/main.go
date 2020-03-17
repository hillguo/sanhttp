package main

import (
	"bytes"
	"fmt"
	"github.com/hillguo/sanhttp/tool/gencode"
	"strings"
	"text/template"
)

var format_server = `package main

import (
	"github.com/hillguo/sanhttp"
)

func main() {
	app := sanhttp.Default()
`
var formmat_route = `
	app.Any("/%s/%s", %sCGIFactoryHandler)
`
var end= `
	app.Run("0.0.0.0:8080")
}
`






var factory1 = `package main

import (
	"github.com/hillguo/sanhttp/ctx"
	"github.com/hillguo/sanhttp/basecgi"
)
`

var factoryFunc = `
type %sCGIFactory struct {

}

func(f *%sCGIFactory) Process(c *ctx.Context) {

	cgi := f.CreateCgi()
	cgi.Execute(c)
}

func(f *%sCGIFactory) CreateCgi() basecgi.Cgier{
	return &%sCGI{}
}

func(f *%sCGIFactory) GetName()string{
	return "%sCGIFactory"
}

var %sCGIFactoryHandler = (&%sCGIFactory{}).Process
`

var cgitmplStr = `package main
import (
	"encoding/json"

	"github.com/hillguo/sanhttp/ctx"
	"github.com/hillguo/sanhttp/basecgi"
	"github.com/hillguo/sanhttp/errs"
)

type {{ .MethodName }}CGI struct {
	basecgi.BaseCGI

	req {{.InputType}}
	resp {{.OutputType}}
}

func(cgi *{{.MethodName}}CGI) Process(c *ctx.Context, req *{{.InputType}}, resp *{{.OutputType}}) error{

	return errs.New(100000, "method unimplemented")
}
func(cgi *{{.MethodName}}CGI) GetName() string {
	return ""
}


func(cgi *{{.MethodName}}CGI) Execute(c *ctx.Context){
	cgi.BeforeExecute(c)
	cgi.BeforeProcess(c)
	if c.ContentType() == "application/json" {
		data ,_ := c.GetRawData()
		json.Unmarshal(data, &cgi.req)
	}
	err := cgi.Process(c, &cgi.req, &cgi.resp)
	if err != nil {
		if e, ok := err.(*errs.Error) ; ok {
			c.JSON(200, e)
		} else {
			e := errs.New(1000, err.Error())
			c.JSON(200, e)
		}
	} else{
		c.JSON(200, cgi.resp)
	}


	cgi.AfterExecute(c)
}
`

func genmain(protoInfo *gencode.ProtoFileInfo) []gencode.FileNameData {
	data := format_server
	for _, methodInfo := range protoInfo.Methods {
		data += fmt.Sprintf(formmat_route, protoInfo.ModuleName, strings.ToLower(methodInfo.MethodName),
			methodInfo.MethodName)
	}
	data += end
	return []gencode.FileNameData{ {Name:"xmain.go", Data:data}}
}

func genfactory(protoInfo *gencode.ProtoFileInfo) []gencode.FileNameData{
	data := factory1
	for _, methodInfo := range protoInfo.Methods {
		data += fmt.Sprintf(factoryFunc,methodInfo.MethodName,methodInfo.MethodName,methodInfo.MethodName,
			methodInfo.MethodName,methodInfo.MethodName,methodInfo.MethodName,methodInfo.MethodName,
			methodInfo.MethodName)
	}
	return []gencode.FileNameData{ {Name:"cgifactory.go", Data:data}}
}

func gencgis(protoInfo *gencode.ProtoFileInfo) []gencode.FileNameData {
	basicTmpl, err := template.New("basic").Parse(cgitmplStr)  //创建、解析
	if err != nil {
		panic(err)
	}

	flienamedata := make([]gencode.FileNameData, 0 , len(protoInfo.Methods))
	for  _, methodInfo := range protoInfo.Methods {
		var buf bytes.Buffer
		basicTmpl.Execute(&buf, methodInfo) // 渲染
		name := strings.ToLower(methodInfo.MethodName) + "_cgi.go"
		flienamedata = append(flienamedata, gencode.FileNameData{Name:name, Data:buf.String()})
	}

	return flienamedata
}

func main() {
	gencode.Main(genmain, genfactory, gencgis)
}
