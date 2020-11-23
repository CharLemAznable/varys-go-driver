package varys

import (
    "github.com/CharLemAznable/gokits"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestWechatAppToken(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatAppToken("codeName")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatAppToken("codeName")
    a.NotNil(err)
    a.Equal("请求WechatAppToken失败", err.Error())

    appTokenResp := WechatAppTokenResp{AppId: "appId", Token: "token", Ticket: "ticket"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-wechat-app-token/codeName", r.URL.EscapedPath())

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(appTokenResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatAppToken("codeName")
    a.Nil(err)
    a.Equal(appTokenResp.AppId, resp.AppId)
    a.Equal(appTokenResp.Token, resp.Token)
    a.Equal(appTokenResp.Ticket, resp.Ticket)
}

func TestWechatMpLogin(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatMpLogin("codeName", "js_code")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatMpLogin("codeName", "js_code")
    a.NotNil(err)
    a.Equal("请求WechatMpLogin失败", err.Error())

    mpLoginResp := WechatMpLoginResp{OpenId: "open_id", SessionKey: "session_key", UnionId: "unionid"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/proxy-wechat-mp-login/codeName", r.URL.EscapedPath())
            a.Equal("js_code", r.FormValue("js_code"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(mpLoginResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatMpLogin("codeName", "js_code")
    a.Nil(err)
    a.Equal(mpLoginResp.OpenId, resp.OpenId)
    a.Equal(mpLoginResp.SessionKey, resp.SessionKey)
    a.Equal(mpLoginResp.UnionId, resp.UnionId)
}

func TestWechatAppJsConfig(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatAppJsConfig("codeName", "http://varys.com/test")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatAppJsConfig("codeName", "http://varys.com/test")
    a.NotNil(err)
    a.Equal("请求WechatAppJsConfig失败", err.Error())

    jsConfigResp := WechatJsConfigResp{AppId: "appId", NonceStr: "nonceStr", Timestamp: 100, Signature: "signature"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-wechat-app-js-config/codeName", r.URL.EscapedPath())
            a.Equal("http://varys.com/test", r.FormValue("url"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(jsConfigResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatAppJsConfig("codeName", "http://varys.com/test")
    a.Nil(err)
    a.Equal(jsConfigResp.AppId, resp.AppId)
    a.Equal(jsConfigResp.NonceStr, resp.NonceStr)
    a.Equal(jsConfigResp.Timestamp, resp.Timestamp)
    a.Equal(jsConfigResp.Signature, resp.Signature)
}

func TestWechatTpToken(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatTpToken("codeName")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatTpToken("codeName")
    a.NotNil(err)
    a.Equal("请求WechatTpToken失败", err.Error())

    tpTokenResp := WechatTpTokenResp{AppId: "appId", Token: "token"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-wechat-tp-token/codeName", r.URL.EscapedPath())

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(tpTokenResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatTpToken("codeName")
    a.Nil(err)
    a.Equal(tpTokenResp.AppId, resp.AppId)
    a.Equal(tpTokenResp.Token, resp.Token)
}

func TestWechatTpAuthToken(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatTpAuthToken("codeName", "authorizerAppId")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatTpAuthToken("codeName", "authorizerAppId")
    a.NotNil(err)
    a.Equal("请求WechatTpAuthToken失败", err.Error())

    tpAuthTokenResp := WechatTpAuthTokenResp{AppId: "appId", AuthorizerAppId: "authorizerAppId", Token: "token", Ticket: "ticket"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-wechat-tp-auth-token/codeName/authorizerAppId", r.URL.EscapedPath())

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(tpAuthTokenResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatTpAuthToken("codeName", "authorizerAppId")
    a.Nil(err)
    a.Equal(tpAuthTokenResp.AppId, resp.AppId)
    a.Equal(tpAuthTokenResp.AuthorizerAppId, resp.AuthorizerAppId)
    a.Equal(tpAuthTokenResp.Token, resp.Token)
    a.Equal(tpAuthTokenResp.Ticket, resp.Ticket)

    _, err = wechatTpAuthTokenCache.Value("codeName:authorizerAppId")
    a.NotNil(err)
    a.Equal("WechatTpAuthKey type error", err.Error())
}

func TestWechatTpAuthJsConfig(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatTpAuthJsConfig("codeName", "http://varys.com/test")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatTpAuthJsConfig("codeName", "http://varys.com/test")
    a.NotNil(err)
    a.Equal("请求WechatTpAuthJsConfig失败", err.Error())

    jsConfigResp := WechatJsConfigResp{AppId: "appId", NonceStr: "nonceStr", Timestamp: 100, Signature: "signature"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-wechat-tp-auth-js-config/codeName", r.URL.EscapedPath())
            a.Equal("http://varys.com/test", r.FormValue("url"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(jsConfigResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatTpAuthJsConfig("codeName", "http://varys.com/test")
    a.Nil(err)
    a.Equal(jsConfigResp.AppId, resp.AppId)
    a.Equal(jsConfigResp.NonceStr, resp.NonceStr)
    a.Equal(jsConfigResp.Timestamp, resp.Timestamp)
    a.Equal(jsConfigResp.Signature, resp.Signature)
}

func TestWechatCorpToken(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatCorpToken("codeName")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatCorpToken("codeName")
    a.NotNil(err)
    a.Equal("请求WechatCorpToken失败", err.Error())

    corpTokenResp := WechatCorpTokenResp{CorpId: "corpId", Token: "token", }
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-wechat-corp-token/codeName", r.URL.EscapedPath())

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(corpTokenResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatCorpToken("codeName")
    a.Nil(err)
    a.Equal(corpTokenResp.CorpId, resp.CorpId)
    a.Equal(corpTokenResp.Token, resp.Token)
}

func TestWechatCorpTpAuthToken(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := WechatCorpTpAuthToken("codeName", "corpId")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = WechatCorpTpAuthToken("codeName", "corpId")
    a.NotNil(err)
    a.Equal("请求WechatCorpTpAuthToken失败", err.Error())

    corpTpAuthTokenResp := WechatCorpTpAuthTokenResp{CorpId: "appId", SuiteId: "suiteId", Token: "token"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-wechat-corp-tp-auth-token/codeName/corpId", r.URL.EscapedPath())

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(corpTpAuthTokenResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatCorpTpAuthToken("codeName", "corpId")
    a.Nil(err)
    a.Equal(corpTpAuthTokenResp.CorpId, resp.CorpId)
    a.Equal(corpTpAuthTokenResp.SuiteId, resp.SuiteId)
    a.Equal(corpTpAuthTokenResp.Token, resp.Token)

    _, err = wechatCorpTpAuthTokenCache.Value("codeName:corpId")
    a.NotNil(err)
    a.Equal("WechatCorpTpAuthKey type error", err.Error())
}

func TestToutiaoAppToken(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := ToutiaoAppToken("codeName")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = ToutiaoAppToken("codeName")
    a.NotNil(err)
    a.Equal("请求ToutiaoAppToken失败", err.Error())

    appTokenResp := ToutiaoAppTokenResp{Token: "token", AppId: "appId"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-toutiao-app-token/codeName", r.URL.EscapedPath())

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(appTokenResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := ToutiaoAppToken("codeName")
    a.Nil(err)
    a.Equal(appTokenResp.Token, resp.Token)
    a.Equal(appTokenResp.AppId, resp.AppId)
}

func TestFengniaoAppToken(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    ConfigInstance.Address = ""
    _, err := FengniaoAppToken("codeName")
    a.NotNil(err)
    a.Equal("未配置Varys.Address", err.Error())

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(http.StatusInternalServerError)
        }))
    ConfigInstance.Address = testServer.URL
    _, err = FengniaoAppToken("codeName")
    a.NotNil(err)
    a.Equal("请求FengniaoAppToken失败", err.Error())

    appTokenResp := ToutiaoAppTokenResp{AppId: "appId", Token: "token"}
    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/query-fengniao-app-token/codeName", r.URL.EscapedPath())

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte(gokits.Json(appTokenResp)))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := FengniaoAppToken("codeName")
    a.Nil(err)
    a.Equal(appTokenResp.AppId, resp.AppId)
    a.Equal(appTokenResp.Token, resp.Token)
}
