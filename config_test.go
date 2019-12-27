package varys

import (
    "testing"
)

func TestConfigInstance(t *testing.T) {
    ConfigInstance = NewConfig()
    if ConfigInstance.Address != "http://127.0.0.1:4236" {
        t.Errorf("Default Address mismatch")
    }
    if ConfigInstance.AppTokenCachePath != "/query-wechat-app-token/" {
        t.Errorf("Default AppTokenCachePath mismatch")
    }
    if ConfigInstance.AppAuthorizerTokenCachePath != "/query-wechat-app-authorizer-token/" {
        t.Errorf("Default AppAuthorizerTokenCachePath mismatch")
    }
    if ConfigInstance.CorpTokenCachePath != "/query-wechat-corp-token/" {
        t.Errorf("Default CorpTokenCachePath mismatch")
    }
    if ConfigInstance.CorpAuthorizerTokenCachePath != "/query-wechat-corp-authorizer-token/" {
        t.Errorf("Default CorpAuthorizerTokenCachePath mismatch")
    }
    if ConfigInstance.ProxyWechatAppPath != "/proxy-wechat-app/" {
        t.Errorf("Default ProxyWechatAppPath mismatch")
    }
    if ConfigInstance.ProxyWechatCorpPath != "/proxy-wechat-corp/" {
        t.Errorf("Default ProxyWechatCorpPath mismatch")
    }
    if ConfigInstance.AppTokenCacheDurationInMinutes != 10 {
        t.Errorf("Default AppTokenCacheDurationInMinutes mismatch")
    }
    if ConfigInstance.AppAuthorizerTokenCacheDurationInMinutes != 10 {
        t.Errorf("Default AppAuthorizerTokenCacheDurationInMinutes mismatch")
    }
    if ConfigInstance.CorpTokenCacheDurationInMinutes != 10 {
        t.Errorf("Default CorpTokenCacheDurationInMinutes mismatch")
    }
    if ConfigInstance.CorpAuthorizerTokenCacheDurationInMinutes != 10 {
        t.Errorf("Default CorpAuthorizerTokenCacheDurationInMinutes mismatch")
    }

    ConfigInstance = NewConfig(
        WithAddress("1"),
        WithAppTokenCachePath("2"),
        WithAppAuthorizerTokenCachePath("3"),
        WithCorpTokenCachePath("4"),
        WithCorpAuthorizerTokenCachePath("5"),
        WithProxyWechatAppPath("6"),
        WithProxyWechatCorpPath("7"),
        WithAppTokenCacheDurationInMinutes(5),
        WithAppAuthorizerTokenCacheDurationInMinutes(5),
        WithCorpTokenCacheDurationInMinutes(5),
        WithCorpAuthorizerTokenCacheDurationInMinutes(5),
    )
    if ConfigInstance.Address != "1" {
        t.Errorf("Init Address mismatch")
    }
    if ConfigInstance.AppTokenCachePath != "2" {
        t.Errorf("Init AppTokenCachePath mismatch")
    }
    if ConfigInstance.AppAuthorizerTokenCachePath != "3" {
        t.Errorf("Init AppAuthorizerTokenCachePath mismatch")
    }
    if ConfigInstance.CorpTokenCachePath != "4" {
        t.Errorf("Init CorpTokenCachePath mismatch")
    }
    if ConfigInstance.CorpAuthorizerTokenCachePath != "5" {
        t.Errorf("Init CorpAuthorizerTokenCachePath mismatch")
    }
    if ConfigInstance.ProxyWechatAppPath != "6" {
        t.Errorf("Init ProxyWechatAppPath mismatch")
    }
    if ConfigInstance.ProxyWechatCorpPath != "7" {
        t.Errorf("Init ProxyWechatCorpPath mismatch")
    }
    if ConfigInstance.AppTokenCacheDurationInMinutes != 5 {
        t.Errorf("Init AppTokenCacheDurationInMinutes mismatch")
    }
    if ConfigInstance.AppAuthorizerTokenCacheDurationInMinutes != 5 {
        t.Errorf("Init AppAuthorizerTokenCacheDurationInMinutes mismatch")
    }
    if ConfigInstance.CorpTokenCacheDurationInMinutes != 5 {
        t.Errorf("Init CorpTokenCacheDurationInMinutes mismatch")
    }
    if ConfigInstance.CorpAuthorizerTokenCacheDurationInMinutes != 5 {
        t.Errorf("Init CorpAuthorizerTokenCacheDurationInMinutes mismatch")
    }
}

func TestConfigPath(t *testing.T) {
    ConfigInstance = NewConfig()
    if ConfigInstance.Path("/abc") != "http://127.0.0.1:4236/abc" {
        t.Errorf("Error Path http://127.0.0.1:4236 + /abc")
    }
    if ConfigInstance.Path("abc") != "http://127.0.0.1:4236/abc" {
        t.Errorf("Error Path http://127.0.0.1:4236 + abc")
    }

    ConfigInstance = NewConfig(WithAddress("http://127.0.0.1:4236/"))
    if ConfigInstance.Path("/abc") != "http://127.0.0.1:4236/abc" {
        t.Errorf("Error Path http://127.0.0.1:4236/ + /abc")
    }
    if ConfigInstance.Path("abc") != "http://127.0.0.1:4236/abc" {
        t.Errorf("Error Path http://127.0.0.1:4236/ + abc")
    }
}
