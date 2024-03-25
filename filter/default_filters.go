package filter

import (
	"github.com/xfali/xlog"
	"github.com/ydx1011/restclient/buffer"
	"io"
	"net/http"
	"time"
)

type Log struct {
	Log  xlog.Logger
	Tag  string
	pool buffer.Pool
}

func NewLog(log xlog.Logger, tag string) *Log {
	if tag == "" {
		tag = "restclient"
	}
	return &Log{
		Log:  log,
		Tag:  tag,
		pool: buffer.NewPool(),
	}
}

func (log *Log) Filter(request *http.Request, fc FilterChain) (*http.Response, error) {
	reqBuf := buffer.NewReadWriteCloser(log.pool)

	var reqData []byte
	if request.Body != nil {
		_, err := io.Copy(reqBuf, request.Body)
		if err != nil {
			return nil, err
		}
		reqData = reqBuf.Bytes()
		// close old request body
		request.Body.Close()
		request.Body = reqBuf
	}

	now := time.Now()
	id := RandomId(10)
	log.Log.Infof("[%s request %s]: url: %s , method: %s , header: %v , body: %s \n",
		log.Tag, id, request.URL.String(), request.Method, request.Header, string(reqData))

	resp, err := fc.Filter(request)
	var (
		status   int
		header   http.Header
		respData []byte
	)
	if resp != nil {
		status = resp.StatusCode
		header = resp.Header

		if resp.Body != nil {
			respBuf := buffer.NewReadWriteCloser(log.pool)

			_, rspErr := io.Copy(respBuf, resp.Body)
			resp.Body.Close()
			if rspErr == nil {
				respData = respBuf.Bytes()
			}
			resp.Body = respBuf
		}
	}
	if err != nil {
		log.Log.Infof("[%s response %s]: use time: %d ms, status: %d , header: %v, result: %s, error: %v \n",
			log.Tag, id, time.Since(now)/time.Millisecond, status, header, string(respData), err)
	} else {
		log.Log.Infof("[%s response %s]: use time: %d ms, status: %d , header: %v, result: %s \n",
			log.Tag, id, time.Since(now)/time.Millisecond, status, header, string(respData))
	}

	return resp, err
}
