package varys

import (
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestWechatApp(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/proxy-wechat-app/codeName/wechatApp", r.URL.EscapedPath())
            a.Equal("b", r.Header.Get("a"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatAppResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatApp("codeName", "/wechatApp").Prop("a", "b").Get()
    a.Nil(err)
    a.Equal("defaultWechatAppResp", resp)

    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("POST", r.Method)
            a.Equal("/proxy-wechat-app/codeName/wechatAppParam/testParam", r.URL.EscapedPath())
            bytes, _ := ioutil.ReadAll(r.Body)
            a.Equal("requestBody", string(bytes))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatAppParamResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err = WechatApp("codeName", "wechatAppParam/%s", "testParam").RequestBody("requestBody").Post()
    a.Nil(err)
    a.Equal("defaultWechatAppParamResp", resp)
}

func TestWechatMp(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/proxy-wechat-mp/codeName/wechatMp", r.URL.EscapedPath())
            a.Equal("b", r.Header.Get("a"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatMpResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatMp("codeName", "/wechatMp").Prop("a", "b").Get()
    a.Nil(err)
    a.Equal("defaultWechatMpResp", resp)

    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("POST", r.Method)
            a.Equal("/proxy-wechat-mp/codeName/wechatMpParam/testParam", r.URL.EscapedPath())
            bytes, _ := ioutil.ReadAll(r.Body)
            a.Equal("requestBody", string(bytes))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatMpParamResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err = WechatMp("codeName", "wechatMpParam/%s", "testParam").RequestBody("requestBody").Post()
    a.Nil(err)
    a.Equal("defaultWechatMpParamResp", resp)
}

func TestWechatTp(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/proxy-wechat-tp/codeName/wechatTp", r.URL.EscapedPath())
            a.Equal("b", r.Header.Get("a"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatTpResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatTp("codeName", "/wechatTp").Prop("a", "b").Get()
    a.Nil(err)
    a.Equal("defaultWechatTpResp", resp)

    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("POST", r.Method)
            a.Equal("/proxy-wechat-tp/codeName/wechatTpParam/testParam", r.URL.EscapedPath())
            bytes, _ := ioutil.ReadAll(r.Body)
            a.Equal("requestBody", string(bytes))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatTpParamResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err = WechatTp("codeName", "wechatTpParam/%s", "testParam").RequestBody("requestBody").Post()
    a.Nil(err)
    a.Equal("defaultWechatTpParamResp", resp)
}

func TestWechatCorp(t *testing.T) {
    a := assert.New(t)
    ConfigInstance = NewConfig()

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("GET", r.Method)
            a.Equal("/proxy-wechat-corp/codeName/wechatCorp", r.URL.EscapedPath())
            a.Equal("b", r.FormValue("a"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatCorpResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatCorp("codeName", "/wechatCorp").Params("a", "b").Get()
    a.Nil(err)
    a.Equal("defaultWechatCorpResp", resp)

    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            a.Equal("POST", r.Method)
            a.Equal("/proxy-wechat-corp/codeName/wechatCorpParam/testParam", r.URL.EscapedPath())
            a.Equal("b", r.FormValue("a"))

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatCorpParamResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err = WechatCorp("codeName", "wechatCorpParam/%s", "testParam").ParamsMapping(map[string]string{"a": "b"}).Post()
    a.Nil(err)
    a.Equal("defaultWechatCorpParamResp", resp)
}
