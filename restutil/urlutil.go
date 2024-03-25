package restutil

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
)

type UrlBuilder struct {
	url        string
	leftDelim  string
	rightDelim string
	path       map[string]interface{}
	query      map[string]interface{}
}

// NewUrlBuilder URL构造器
func NewUrlBuilder(url string) *UrlBuilder {
	return &UrlBuilder{
		url:       url,
		leftDelim: ":",
	}
}

// Delims delimiters设置占位符，用于替换url的path参数
func (b *UrlBuilder) Delims(leftDelim, rightDelim string) *UrlBuilder {
	b.leftDelim = leftDelim
	b.rightDelim = rightDelim
	return b
}

// PathVariable 增加path变量参数
func (b *UrlBuilder) PathVariable(key string, value interface{}) *UrlBuilder {
	if b.path == nil {
		b.path = map[string]interface{}{}
	}
	b.path[key] = value
	return b
}

// QueryVariable 增加query参数
func (b *UrlBuilder) QueryVariable(key string, value interface{}) *UrlBuilder {
	if b.query == nil {
		b.query = map[string]interface{}{}
	}
	b.query[key] = value
	return b
}

// Build 创建url
func (b *UrlBuilder) Build() string {
	buf := strings.Builder{}
	if len(b.path) > 0 {
		buf.WriteString(ReplaceUrl(b.url, b.leftDelim, b.rightDelim, b.path))
	} else {
		buf.WriteString(b.url)
	}
	if len(b.query) > 0 {
		query := EncodeQuery(b.query)
		if b.url[len(b.url)-1] == '?' {
			buf.WriteString(query)
		} else {
			buf.WriteString("?")
			buf.WriteString(query)
		}
	}
	return buf.String()
}

func (b *UrlBuilder) String() string {
	return b.Build()
}

func ReplaceUrl(uri string, leftDelim string, rightDelim string, keyAndValue map[string]interface{}) string {
	if len(keyAndValue) == 0 {
		return uri
	}
	for k, v := range keyAndValue {
		uri = strings.Replace(uri, fmt.Sprintf("%s%v%s", leftDelim, k, rightDelim), url.QueryEscape(fmt.Sprintf("%v", v)), -1)
	}
	return uri
}

func EncodeQuery(keyAndValue map[string]interface{}) string {
	if len(keyAndValue) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for k, v := range keyAndValue {
		buf.WriteString(url.QueryEscape(fmt.Sprintf("%v", k)))
		buf.WriteString("=")
		buf.WriteString(url.QueryEscape(fmt.Sprintf("%v", v)))
		buf.WriteString("&")
	}
	format := buf.String()
	format = format[:len(format)-1]
	return format
}
