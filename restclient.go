package restclient

import "github.com/ydx1011/restclient/request"

type RestClient interface {
	// 发起请求
	// url：请求路径
	// params：请求参数，见ex_params.go具体定义
	Exchange(url string, opts ...request.Opt) Error
}
