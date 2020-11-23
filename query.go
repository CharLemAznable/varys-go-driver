package varys

import (
    "errors"
    "github.com/CharLemAznable/gokits"
    urlEncoder "net/url"
)

type WechatAppTokenResp struct {
    Error  string `json:"error"`
    AppId  string `json:"appId"`
    Token  string `json:"token"`
    Ticket string `json:"ticket"`
}

type WechatMpLoginResp struct {
    OpenId     string `json:"openid"`
    SessionKey string `json:"session_key"`
    UnionId    string `json:"unionid"`
    Errcode    int    `json:"errcode"`
    Errmsg     string `json:"errmsg"`
}

type WechatTpTokenResp struct {
    Error string `json:"error"`
    AppId string `json:"appId"`
    Token string `json:"token"`
}

type WechatTpAuthTokenResp struct {
    Error           string `json:"error"`
    AppId           string `json:"appId"`
    AuthorizerAppId string `json:"authorizerAppId"`
    Token           string `json:"token"`
    Ticket          string `json:"ticket"`
}

type WechatCorpTokenResp struct {
    Error  string `json:"error"`
    CorpId string `json:"corpId"`
    Token  string `json:"token"`
}

type WechatCorpTpAuthTokenResp struct {
    Error   string `json:"error"`
    SuiteId string `json:"suiteId"`
    CorpId  string `json:"corpId"`
    Token   string `json:"token"`
}

type WechatJsConfigResp struct {
    AppId     string `json:"appId"`
    NonceStr  string `json:"nonceStr"`
    Timestamp int64  `json:"timestamp"`
    Signature string `json:"signature"`
}

type ToutiaoAppTokenResp struct {
    Error string `json:"error"`
    AppId string `json:"appId"`
    Token string `json:"token"`
}

type FengniaoAppTokenResp struct {
    Error string `json:"error"`
    AppId string `json:"appId"`
    Token string `json:"token"`
}

func WechatAppToken(codeName string) (*WechatAppTokenResp, error) {
    value, err := wechatAppTokenCache.Value(codeName)
    if nil != err {
        return nil, err
    }
    return value.Data().(*WechatAppTokenResp), nil
}

func WechatMpLogin(codeName, jsCode string) (*WechatMpLoginResp, error) {
    resp, err := httpGet(gokits.PathJoin("/proxy-wechat-mp-login/", codeName) + "?js_code=" + urlEncoder.QueryEscape(jsCode))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatMpLogin失败")
    }
    return gokits.UnJson(resp, new(WechatMpLoginResp)).(*WechatMpLoginResp), nil
}

func WechatAppJsConfig(codeName, url string) (*WechatJsConfigResp, error) {
    resp, err := httpGet(gokits.PathJoin("/query-wechat-app-js-config/", codeName) + "?url=" + urlEncoder.QueryEscape(url))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatAppJsConfig失败")
    }
    return gokits.UnJson(resp, new(WechatJsConfigResp)).(*WechatJsConfigResp), nil
}

func WechatTpToken(codeName string) (*WechatTpTokenResp, error) {
    value, err := wechatTpTokenCache.Value(codeName)
    if nil != err {
        return nil, err
    }
    return value.Data().(*WechatTpTokenResp), nil
}

func WechatTpAuthToken(codeName, authorizerAppId string) (*WechatTpAuthTokenResp, error) {
    value, err := wechatTpAuthTokenCache.Value(WechatTpAuthKey{
        CodeName: codeName, AuthorizerAppId: authorizerAppId})
    if nil != err {
        return nil, err
    }
    return value.Data().(*WechatTpAuthTokenResp), nil
}

func WechatTpAuthJsConfig(codeName, url string) (*WechatJsConfigResp, error) {
    resp, err := httpGet(gokits.PathJoin("/query-wechat-tp-auth-js-config/", codeName) + "?url=" + urlEncoder.QueryEscape(url))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatTpAuthJsConfig失败")
    }
    return gokits.UnJson(resp, new(WechatJsConfigResp)).(*WechatJsConfigResp), nil
}

func WechatCorpToken(codeName string) (*WechatCorpTokenResp, error) {
    value, err := wechatCorpTokenCache.Value(codeName)
    if nil != err {
        return nil, err
    }
    return value.Data().(*WechatCorpTokenResp), nil
}

func WechatCorpTpAuthToken(codeName, corpId string) (*WechatCorpTpAuthTokenResp, error) {
    value, err := wechatCorpTpAuthTokenCache.Value(WechatCorpTpAuthKey{
        CodeName: codeName, CorpId: corpId})
    if nil != err {
        return nil, err
    }
    return value.Data().(*WechatCorpTpAuthTokenResp), nil
}

func ToutiaoAppToken(codeName string) (*ToutiaoAppTokenResp, error) {
    value, err := toutiaoAppTokenCache.Value(codeName)
    if nil != err {
        return nil, err
    }
    return value.Data().(*ToutiaoAppTokenResp), nil
}

func FengniaoAppToken(codeName string) (*FengniaoAppTokenResp, error) {
    value, err := fengniaoAppTokenCache.Value(codeName)
    if nil != err {
        return nil, err
    }
    return value.Data().(*FengniaoAppTokenResp), nil
}
