package restclient

import (
	"github.com/ydx1011/restclient/buffer"
	"github.com/ydx1011/restclient/filter"
	"net/http"
	"time"
)

// SetTimeout 设置读写超时
func SetTimeout(timeout time.Duration) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.timeout = timeout
	}
}

// SetConverters 配置初始转换器列表
func SetConverters(convs []Converter) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.converters = convs
	}
}

// AddConverters 添加初始转换器列表
func AddConverters(convs ...Converter) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.converters = append(client.converters, convs...)
	}
}

// SetRoundTripper 配置连接池
func SetRoundTripper(tripper http.RoundTripper) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.transport = tripper
	}
}

// SetClientCreator 配置http客户端创建器
func SetClientCreator(cliCreator HttpClientCreator) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.cliCreator = cliCreator
	}
}

// CookieJar 配置http.Client的CookieJar
func CookieJar(jar http.CookieJar) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.jar = jar
	}
}

// SetAutoAccept 配置是否自动添加accept
func SetAutoAccept(v AcceptFlag) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.acceptFlag = v
	}
}

// SetResponseBodyFlag 配置是否自动添加accept
func SetResponseBodyFlag(v ResponseBodyFlag) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.respFlag = v
	}
}

// SetBufferPool 配置内存池
func SetBufferPool(pool buffer.Pool) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.pool = pool
	}
}

// AddFilter 增加处理filter
func AddFilter(filters ...filter.Filter) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		client.filterManager.Add(filters...)
	}
}

// AddIFilter 增加处理filter
func AddIFilter(filters ...filter.IFilter) func(client *defaultRestClient) {
	return func(client *defaultRestClient) {
		for _, v := range filters {
			if v != nil {
				client.filterManager.Add(v.Filter)
			}
		}
	}
}
