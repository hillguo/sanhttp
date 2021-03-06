package sanhttp

import (
	"encoding/json"
	"reflect"

	log "github.com/hillguo/sanlog"

	"github.com/hillguo/sanhttp/ctx"
	"github.com/hillguo/sanhttp/errs"
)

// HF 对handler 进行包装
func HF(method interface{}) ctx.HandlerFunc {
	typ := reflect.TypeOf(method)
	typVal := reflect.ValueOf(method)
	typName := typ.Name()
	if typ.Kind() != reflect.Func {
		log.Fatalf("%s in not func", typName)
	}

	if typ.NumIn() != 3 {
		log.Fatalf("%s input param num invalid, num:%d", typName, typ.NumIn())
	}

	// First arg must be context.Context
	ctxType := typ.In(0)
	if ctxType != reflect.TypeOf((*ctx.Context)(nil)) {
		log.Infof("%+v", ctxType)
		log.Info(reflect.TypeOf((*ctx.Context)(nil)).Elem())
		log.Fatalf("func %v", typName, " must use *ctx.Context as the first parameter")
	}

	// Second arg need not be a pointer.
	reqType := typ.In(1)

	// Third arg must be a pointer.
	rspType := typ.In(2)
	if reqType.Kind() != reflect.Ptr || rspType.Kind() != reflect.Ptr {
		log.Fatalf("func", typName, " reqType or rspType type not a pointer:", reqType, rspType)
	}

	// Method needs one out.
	if typ.NumOut() != 1 {
		log.Fatalf("func", typName, " has wrong number of outs:", typ.NumOut())
	}
	// The return type of the method must be error.
	if returnType := typ.Out(0); returnType != reflect.TypeOf((*error)(nil)).Elem() {
		log.Fatalf("func", typName, " returns ", returnType.String(), " not error")
	}

	return func(c *ctx.Context) {
		c.SetJSONFrame()
		req := reflect.New(reqType.Elem()).Interface()
		resp := reflect.New(rspType.Elem()).Interface()
		if c.ContentType() == "application/json" {
			data, _ := c.GetRawData()
			err := json.Unmarshal(data, req)
			if err != nil {
				log.Errorf("json unmarshal req body %s error:%s", string(data), err.Error())
			}
		}

		log.Debugf("sanhttp req: %v", req)
		returnValues := typVal.Call([]reflect.Value{reflect.ValueOf(c), reflect.ValueOf(req), reflect.ValueOf(resp)})
		log.Debugf("sanhttp resp: %v", resp)

		err := returnValues[0].Interface()
		header := errs.NewSuccess()
		param := &ctx.Param{LogicsvrIP: c.Request.RemoteAddr}
		if err != nil {
			if e, ok := err.(*errs.Error); ok {
				header = e
			} else if e, ok := err.(error); ok {
				header = errs.New(errs.ErrorUnKnow, e.Error())
			} else {
				log.Error("returnValue type error ", reflect.TypeOf(err).Name())
				header = errs.NewFrameError(errs.ErrorUnKnow, "returnValue type error")
			}
		}

		if c.GetFrameType() == ctx.RawFrame {
			return
		} else if c.GetFrameType() == ctx.ProtoBufFrame {
			c.ProtoBuf(200, resp)
		} else {
			retval := ctx.HTTPBody{Header: header, Param: param, Data: resp}
			c.IndentedJSON(200, retval)
		}
	}
}
