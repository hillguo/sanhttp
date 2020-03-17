package main

import (
	"github.com/hillguo/sanhttp/ctx"
	"github.com/hillguo/sanhttp/basecgi"
)

type TestCGIFactory struct {

}

func(f *TestCGIFactory) Process(c *ctx.Context) {

	cgi := f.CreateCgi()
	cgi.Execute(c)
}

func(f *TestCGIFactory) CreateCgi() basecgi.Cgier{
	return &TestCGI{}
}

func(f *TestCGIFactory) GetName()string{
	return "TestCGIFactory"
}

var TestCGIFactoryHandler = (&TestCGIFactory{}).Process

type GetVidByNameCGIFactory struct {

}

func(f *GetVidByNameCGIFactory) Process(c *ctx.Context) {

	cgi := f.CreateCgi()
	cgi.Execute(c)
}

func(f *GetVidByNameCGIFactory) CreateCgi() basecgi.Cgier{
	return &GetVidByNameCGI{}
}

func(f *GetVidByNameCGIFactory) GetName()string{
	return "GetVidByNameCGIFactory"
}

var GetVidByNameCGIFactoryHandler = (&GetVidByNameCGIFactory{}).Process
