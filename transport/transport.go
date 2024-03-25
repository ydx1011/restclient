package transport

import (
	"net"
	"net/http"
	"time"
)

const (
	ConnectTimeout        = 30 * time.Second
	KeepaliveTime         = 30 * time.Second
	MaxIdleConn           = 100
	MaxIdleConnPerHost    = 5
	IdleConnTimeout       = 90 * time.Second
	TlsHandshakeTimeout   = 10 * time.Second
	ExpectContinueTimeout = 1 * time.Second
)

type Opt func(*http.Transport)

func New(opts ...Opt) *http.Transport {
	ret := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   ConnectTimeout,
			KeepAlive: KeepaliveTime,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          MaxIdleConn,
		MaxIdleConnsPerHost:   MaxIdleConnPerHost,
		IdleConnTimeout:       IdleConnTimeout,
		TLSHandshakeTimeout:   TlsHandshakeTimeout,
		ExpectContinueTimeout: ExpectContinueTimeout,
	}
	for i := range opts {
		opts[i](ret)
	}
	return ret
}
