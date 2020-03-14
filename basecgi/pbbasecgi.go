package basecgi

type PBBaseCGI struct {
	BaseCGI
	pb_req  interface{}
	pb_resp interface{}
}

func (cgi *PBBaseCGI) GetPbReq() interface{} {
	return cgi.pb_req
}

func (cgi *PBBaseCGI) GetPbResp() interface{} {
	return cgi.pb_resp
}

func (cgi *PBBaseCGI) Process() {
	cgi.ProcessRequest(cgi.pb_req, cgi.pb_resp)
	cgi.HandleResult(cgi.pb_resp)
}

func (cgi *PBBaseCGI) ProcessRequest(pb_req, pb_resp interface{}) {

}

func (cgi *PBBaseCGI) HandleResult(pb_resp interface{}) {

}

func (cgi *PBBaseCGI) InitPb() {
	cgi.ParseInput(cgi.pb_req)
	cgi.InitAppBaseInfo(cgi.pb_req)
	cgi.InitResp(cgi.pb_resp)
}

func (cgi *PBBaseCGI) InitAppBaseInfo() {

}

func Test() {
}
