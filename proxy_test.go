package varys

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestWechatApp(t *testing.T) {
    ConfigInstance = NewConfig()

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            if r.Method != "GET" {
                t.Errorf("Except 'Get' got '%s'", r.Method)
            }

            if r.URL.EscapedPath() != "/proxy-wechat-app/codeName/wechatApp" {
                t.Errorf("Except to path '/proxy-wechat-app/codeName/wechatApp',got '%s'", r.URL.EscapedPath())
            }

            if r.Header.Get("a") != "b" {
                t.Errorf("Except Header 'a': 'b',got '%s'", r.Header.Get("a"))
            }

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatAppResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatApp("codeName", "/wechatApp").Prop("a", "b").Get()
    if nil != err {
        t.Errorf("Should has no error")
    }
    if resp != "defaultWechatAppResp" {
        t.Errorf("Response mismatch")
    }

    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            if r.Method != "POST" {
                t.Errorf("Except 'Get' got '%s'", r.Method)
            }

            if r.URL.EscapedPath() != "/proxy-wechat-app/codeName/wechatAppParam/testParam" {
                t.Errorf("Except to path '/proxy-wechat-app/codeName/wechatAppParam/testParam',got '%s'", r.URL.EscapedPath())
            }

            bytes, _ := ioutil.ReadAll(r.Body)
            if string(bytes) != "requestBody" {
                t.Errorf("Except RequestBody 'requestBody',got '%s'", string(bytes))
            }

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatAppParamResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err = WechatApp("codeName", "wechatAppParam/%s", "testParam").RequestBody("requestBody").Post()
    if nil != err {
        t.Errorf("Should has no error")
    }
    if resp != "defaultWechatAppParamResp" {
        t.Errorf("Response mismatch")
    }
}

func TestWechatCorp(t *testing.T) {
    ConfigInstance = NewConfig()

    testServer := httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            if r.Method != "GET" {
                t.Errorf("Except 'Get' got '%s'", r.Method)
            }

            if r.URL.EscapedPath() != "/proxy-wechat-corp/codeName/wechatCorp" {
                t.Errorf("Except to path '/proxy-wechat-corp/codeName/wechatCorp',got '%s'", r.URL.EscapedPath())
            }

            if r.FormValue("a") != "b" {
                t.Errorf("Except FormValue 'a': 'b',got '%s'", r.FormValue("a"))
            }

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatCorpResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err := WechatCorp("codeName", "/wechatCorp").Params("a", "b").Get()
    if nil != err {
        t.Errorf("Should has no error")
    }
    if resp != "defaultWechatCorpResp" {
        t.Errorf("Response mismatch")
    }

    testServer = httptest.NewServer(http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            if r.Method != "POST" {
                t.Errorf("Except 'Get' got '%s'", r.Method)
            }

            if r.URL.EscapedPath() != "/proxy-wechat-corp/codeName/wechatCorpParam/testParam" {
                t.Errorf("Except to path '/proxy-wechat-corp/codeName/wechatCorpParam/testParam',got '%s'", r.URL.EscapedPath())
            }

            if r.FormValue("a") != "b" {
                t.Errorf("Except FormValue 'a': 'b',got '%s'", r.FormValue("a"))
            }

            w.WriteHeader(http.StatusOK)
            _, _ = w.Write([]byte("defaultWechatCorpParamResp"))
        }))
    ConfigInstance.Address = testServer.URL
    resp, err = WechatCorp("codeName", "wechatCorpParam/%s", "testParam").ParamsMapping(map[string]string{"a": "b"}).Post()
    if nil != err {
        t.Errorf("Should has no error")
    }
    if resp != "defaultWechatCorpParamResp" {
        t.Errorf("Response mismatch")
    }
}
