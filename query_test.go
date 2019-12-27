package varys

import (
	"github.com/CharLemAznable/gokits"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppToken(t *testing.T) {
	ConfigInstance = NewConfig()

	ConfigInstance.Address = ""
	_, err := AppToken("codeName")
	if nil == err || "未配置Varys.Address" != err.Error() {
		t.Errorf("Should has error: 未配置Varys.Address")
	}

	testServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
	ConfigInstance.Address = testServer.URL
	_, err = AppToken("codeName")
	if nil == err || "请求AppToken失败" != err.Error() {
		t.Errorf("Should has error: 请求AppToken失败")
	}

	appTokenResp := AppTokenResp{Token: "token", AppId: "appId"}
	testServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Except 'Get' got '%s'", r.Method)
			}

			if r.URL.EscapedPath() != "/query-wechat-app-token/codeName" {
				t.Errorf("Except to path '/query-wechat-app-token/codeName',got '%s'", r.URL.EscapedPath())
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(gokits.Json(appTokenResp)))
		}))
	ConfigInstance.Address = testServer.URL
	resp, err := AppToken("codeName")
	if nil != err {
		t.Errorf("Should has no error")
	}
	if resp.Token != appTokenResp.Token ||
		resp.AppId != appTokenResp.AppId {
		t.Errorf("Response mismatch")
	}
}

func TestAppAuthorizerToken(t *testing.T) {
	ConfigInstance = NewConfig()

	testServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
	ConfigInstance.Address = testServer.URL
	_, err := AppAuthorizerToken("codeName", "authorizerAppId")
	if nil == err || "请求AppAuthorizerToken失败" != err.Error() {
		t.Errorf("Should has error: 请求AppAuthorizerToken失败")
	}

	appAuthorizerTokenResp := AppAuthorizerTokenResp{Token: "token", AppId: "appId", AuthorizerAppId: "authorizerAppId"}
	testServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Except 'Get' got '%s'", r.Method)
			}

			if r.URL.EscapedPath() != "/query-wechat-app-authorizer-token/codeName/authorizerAppId" {
				t.Errorf("Except to path '/query-wechat-app-authorizer-token/codeName/authorizerAppId',got '%s'", r.URL.EscapedPath())
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(gokits.Json(appAuthorizerTokenResp)))
		}))
	ConfigInstance.Address = testServer.URL
	resp, err := AppAuthorizerToken("codeName", "authorizerAppId")
	if nil != err {
		t.Errorf("Should has no error")
	}
	if resp.Token != appAuthorizerTokenResp.Token ||
		resp.AppId != appAuthorizerTokenResp.AppId ||
		resp.AuthorizerAppId != appAuthorizerTokenResp.AuthorizerAppId {
		t.Errorf("Response mismatch")
	}
}

func TestCorpToken(t *testing.T) {
	ConfigInstance = NewConfig()

	testServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
	ConfigInstance.Address = testServer.URL
	_, err := CorpToken("codeName")
	if nil == err || "请求CorpToken失败" != err.Error() {
		t.Errorf("Should has error: 请求CorpToken失败")
	}

	corpTokenResp := CorpTokenResp{Token: "token", CorpId: "corpId"}
	testServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Except 'Get' got '%s'", r.Method)
			}

			if r.URL.EscapedPath() != "/query-wechat-corp-token/codeName" {
				t.Errorf("Except to path '/query-wechat-corp-token/codeName',got '%s'", r.URL.EscapedPath())
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(gokits.Json(corpTokenResp)))
		}))
	ConfigInstance.Address = testServer.URL
	resp, err := CorpToken("codeName")
	if nil != err {
		t.Errorf("Should has no error")
	}
	if resp.Token != corpTokenResp.Token ||
		resp.CorpId != corpTokenResp.CorpId {
		t.Errorf("Response mismatch")
	}
}

func TestCorpAuthorizerToken(t *testing.T) {
	ConfigInstance = NewConfig()

	testServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
	ConfigInstance.Address = testServer.URL
	_, err := CorpAuthorizerToken("codeName", "corpId")
	if nil == err || "请求CorpAuthorizerToken失败" != err.Error() {
		t.Errorf("Should has error: 请求CorpAuthorizerToken失败")
	}

	corpAuthorizerTokenResp := CorpAuthorizerTokenResp{Token: "token", CorpId: "appId", SuiteId: "suiteId"}
	testServer = httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Except 'Get' got '%s'", r.Method)
			}

			if r.URL.EscapedPath() != "/query-wechat-corp-authorizer-token/codeName/corpId" {
				t.Errorf("Except to path '/query-wechat-corp-authorizer-token/codeName/corpId',got '%s'", r.URL.EscapedPath())
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(gokits.Json(corpAuthorizerTokenResp)))
		}))
	ConfigInstance.Address = testServer.URL
	resp, err := CorpAuthorizerToken("codeName", "corpId")
	if nil != err {
		t.Errorf("Should has no error")
	}
	if resp.Token != corpAuthorizerTokenResp.Token ||
		resp.CorpId != corpAuthorizerTokenResp.CorpId ||
		resp.SuiteId != corpAuthorizerTokenResp.SuiteId {
		t.Errorf("Response mismatch")
	}
}
