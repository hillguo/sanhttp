package basecgi

import "time"

type ErrType uint32

const (
	NoErr ErrType = iota
	Logic
	Sys
	Ignore
)

type FrameType uint32

const (
	JSON FrameType = iota
	HTML
	RAW
)

type CheckSessionType uint32

const (
	Default CheckSessionType = iota
	JWT
)

type cgi struct {
	CmdID        uint32
	StatusCode   uint32
	ReasonPhrase string
	CgiName      string
}

type metadata struct {
	start_time uint64
}

type BaseCGI struct {
	metadata
	CmdID        uint32
	StatusCode   uint32
	ReasonPhrase string
	CgiName      string
}

func (cgi *BaseCGI) Execute() {
	cgi.BeforeExecute()
	cgi.DoExecute()
	cgi.AfterExecute()
}

func (cgi *BaseCGI) BeforeExecute() {
	_ = time.Now().Unix()
	cgi.InitUserInfo()

	// 上报 cgi 调用数量

}

func (cgi *BaseCGI) DoExecute() {
	cgi.InitPb()
	cgi.BeforeProcess()
	cgi.InitCommParam()
	cgi.CheckSession()
	cgi.AfterCheckSession()
	cgi.SpamFreq() // 频率限制
	cgi.Process()
	cgi.BeforeResponse()
}

func (cgi *BaseCGI) SpamFreq() {

}

func (cgi *BaseCGI) AfterExecute() {

}

func (cgi *BaseCGI) BeforeProcess() {

}

func (cgi *BaseCGI) Process() {

}

func (cgi *BaseCGI) InitPb() {

}

func (cgi *BaseCGI) AfterCheckSession() {

}

func (cgi *BaseCGI) BeforeResponse() {

}

func (cgi *BaseCGI) GetServiceType() uint32 {
	return 0
}

func (cgi *BaseCGI) HandleExceptionResult() {

}

func (cgi *BaseCGI) Report() {

}

func (cgi *BaseCGI) InitCommParam() {

}

func (cgi *BaseCGI) CheckSession() {

}

func (cgi *BaseCGI) InitUserInfo() {
	// cookie 中获取用户信息
}
