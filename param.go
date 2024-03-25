package restclient

import (
	"context"
	"github.com/ydx1011/restclient/filter"
	"github.com/ydx1011/restclient/request"
	"net/http"
)

type defaultParam struct {
	ctx           context.Context
	method        string
	header        http.Header
	filterManager filter.FilterManager

	reqBody  interface{}
	result   interface{}
	response *http.Response
	respFlag bool
}

func emptyParam() *defaultParam {
	return &defaultParam{
		method: http.MethodGet,
		ctx:    context.Background(),
		header: make(http.Header),
	}
}

func (p *defaultParam) Set(key string, value interface{}) {
	switch key {
	case request.KeyMethod:
		p.method = value.(string)
	case request.KeyAddFilter:
		p.filterManager.Add(value.([]filter.Filter)...)
	case request.KeyRequestContext:
		p.ctx = value.(context.Context)
	case request.KeyRequestHeader:
		p.header = value.(http.Header)
	case request.KeyRequestAddHeader:
		ss := value.([]string)
		p.header.Add(ss[0], ss[1])
	case request.KeyRequestAddCookie:
		p.addCookies(value.([]*http.Cookie))
	case request.KeyRequestBody:
		p.reqBody = value
	case request.KeyResult:
		p.result = value
	case request.KeyResponse:
		rs := value.([]interface{})
		p.response = rs[0].(*http.Response)
		p.respFlag = rs[1].(bool)
	}
}

func (p *defaultParam) addCookies(cookies []*http.Cookie) {
	if len(cookies) == 0 {
		return
	}
	req := &http.Request{
		Header: make(http.Header),
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}

	for k, vs := range req.Header {
		for _, v := range vs {
			p.header.Add(k, v)
		}
	}
}
