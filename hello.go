package hello

import "github.com/valyala/fasthttp"

var proxyClient = &fasthttp.HostClient{
	IsTLS: true,
	Addr:  "domain.net"}

//Hello ...
func Hello() string {
	return "hello"
}

//GetProxyClient ...
func GetProxyClient() *fasthttp.HostClient {
	return proxyClient
}
