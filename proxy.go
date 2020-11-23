package varys

import (
    "fmt"
    "github.com/CharLemAznable/gokits"
)

type ProxyReq struct {
    httpReq *gokits.HttpReq
}

func NewProxyReq(baseUrl string) *ProxyReq {
    proxyReq := new(ProxyReq)
    proxyReq.httpReq = gokits.NewHttpReq(baseUrl)
    return proxyReq
}

func (proxyReq *ProxyReq) Params(name string, value string, more ...string) *ProxyReq {
    proxyReq.httpReq.Params(name, value, more...)
    return proxyReq
}

func (proxyReq *ProxyReq) ParamsMapping(params map[string]string) *ProxyReq {
    proxyReq.httpReq.ParamsMapping(params)
    return proxyReq
}

func (proxyReq *ProxyReq) RequestBody(requestBody string) *ProxyReq {
    proxyReq.httpReq.RequestBody(requestBody)
    return proxyReq
}

func (proxyReq *ProxyReq) Prop(name string, value string) *ProxyReq {
    proxyReq.httpReq.Prop(name, value)
    return proxyReq
}

func (proxyReq *ProxyReq) Get() (string, error) {
    return proxyReq.httpReq.Get()
}

func (proxyReq *ProxyReq) Post() (string, error) {
    return proxyReq.httpReq.Post()
}

func WechatApp(codeName, proxyPathTemplate string, proxyPathArgs ...interface{}) *ProxyReq {
    return NewProxyReq(ConfigInstance.Path("/proxy-wechat-app/",
        codeName, proxyPath(proxyPathTemplate, proxyPathArgs...)))
}

func WechatTp(codeName, proxyPathTemplate string, proxyPathArgs ...interface{}) *ProxyReq {
    return NewProxyReq(ConfigInstance.Path("/proxy-wechat-tp/",
        codeName, proxyPath(proxyPathTemplate, proxyPathArgs...)))
}

func WechatTpAuth(codeName, authorizerAppId, proxyPathTemplate string, proxyPathArgs ...interface{}) *ProxyReq {
    return NewProxyReq(ConfigInstance.Path("/proxy-wechat-tp-auth/",
        codeName, authorizerAppId, proxyPath(proxyPathTemplate, proxyPathArgs...)))
}

func WechatCorp(codeName, proxyPathTemplate string, proxyPathArgs ...interface{}) *ProxyReq {
    return NewProxyReq(ConfigInstance.Path("/proxy-wechat-corp/",
        codeName, proxyPath(proxyPathTemplate, proxyPathArgs...)))
}

func FengniaoApp(codeName, proxyPathTemplate string, proxyPathArgs ...interface{}) *ProxyReq {
    return NewProxyReq(ConfigInstance.Path("/proxy-fengniao-app/",
        codeName, proxyPath(proxyPathTemplate, proxyPathArgs...)))
}

func proxyPath(proxyPathTemplate string, proxyPathArgs ...interface{}) string {
    proxyPath := proxyPathTemplate
    if len(proxyPathArgs) > 0 {
        proxyPath = fmt.Sprintf(proxyPathTemplate, proxyPathArgs...)
    }
    return proxyPath
}
