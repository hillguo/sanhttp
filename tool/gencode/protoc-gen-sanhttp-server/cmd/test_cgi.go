package main
import (
	"encoding/json"

	"github.com/hillguo/sanhttp/ctx"
	"github.com/hillguo/sanhttp/basecgi"
	"github.com/hillguo/sanhttp/errs"
)

type TestCGI struct {
	basecgi.BaseCGI

	req TestReq
	resp TestRsp
}

func(cgi *TestCGI) Process(c *ctx.Context, req *TestReq, resp *TestRsp) error{

	return errs.New(100000, "method unimplemented")
}
func(cgi *TestCGI) GetName() string {
	return ""
}


func(cgi *TestCGI) Execute(c *ctx.Context){
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
