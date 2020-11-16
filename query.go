package varys

import (
    "errors"
    "github.com/CharLemAznable/gokits"
    "time"
)

type WechatAppTokenResp struct {
    Error string `json:"error"`
    AppId string `json:"appId"`
    Token string `json:"token"`
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

type ToutiaoAppTokenResp struct {
    Error string `json:"error"`
    AppId string `json:"appId"`
    Token string `json:"token"`
}

type WechatMpLoginResp struct {
    OpenId     string `json:"openid"`
    SessionKey string `json:"session_key"`
    UnionId    string `json:"unionid"`
    Errcode    int    `json:"errcode"`
    Errmsg     string `json:"errmsg"`
}

var wechatAppTokenCache *gokits.CacheTable
var wechatTpTokenCache *gokits.CacheTable
var wechatTpAuthTokenCache *gokits.CacheTable
var wechatCorpTokenCache *gokits.CacheTable
var wechatCorpTpAuthTokenCache *gokits.CacheTable
var toutiaoAppTokenCache *gokits.CacheTable

func init() {
    wechatAppTokenCache = gokits.CacheExpireAfterWrite("WechatAppTokenCache")
    wechatAppTokenCache.SetDataLoader(wechatAppTokenLoader)
    wechatTpTokenCache = gokits.CacheExpireAfterWrite("WechatTpTokenCache")
    wechatTpTokenCache.SetDataLoader(wechatTpTokenLoader)
    wechatTpAuthTokenCache = gokits.CacheExpireAfterWrite("WechatTpAuthTokenCache")
    wechatTpAuthTokenCache.SetDataLoader(wechatTpAuthTokenLoader)
    wechatCorpTokenCache = gokits.CacheExpireAfterWrite("WechatCorpTokenCache")
    wechatCorpTokenCache.SetDataLoader(wechatCorpTokenLoader)
    wechatCorpTpAuthTokenCache = gokits.CacheExpireAfterWrite("WechatCorpTpAuthTokenCache")
    wechatCorpTpAuthTokenCache.SetDataLoader(wechatCorpTpAuthTokenLoader)
    toutiaoAppTokenCache = gokits.CacheExpireAfterWrite("ToutiaoAppTokenCache")
    toutiaoAppTokenCache.SetDataLoader(toutiaoAppTokenLoader)
}

func wechatAppTokenLoader(codeName interface{}, _ ...interface{}) (*gokits.CacheItem, error) {
    resp, err := httpGet(gokits.PathJoin("/query-wechat-app-token/", codeName.(string)))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatAppToken失败")
    }
    return gokits.NewCacheItem(codeName,
        ConfigInstance.WechatAppTokenCacheDuration*time.Minute,
        gokits.UnJson(resp, new(WechatAppTokenResp))), nil
}

func wechatTpTokenLoader(codeName interface{}, _ ...interface{}) (*gokits.CacheItem, error) {
    resp, err := httpGet(gokits.PathJoin("/query-wechat-tp-token/", codeName.(string)))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatTpToken失败")
    }
    return gokits.NewCacheItem(codeName,
        ConfigInstance.WechatTpTokenCacheDuration*time.Minute,
        gokits.UnJson(resp, new(WechatTpTokenResp))), nil
}

type WechatTpAuthKey struct {
    CodeName        string
    AuthorizerAppId string
}

func wechatTpAuthTokenLoader(key interface{}, _ ...interface{}) (*gokits.CacheItem, error) {
    authKey, ok := key.(WechatTpAuthKey)
    if !ok {
        return nil, errors.New("WechatTpAuthKey type error")
    }
    resp, err := httpGet(gokits.PathJoin(
        "/query-wechat-tp-auth-token/",
        authKey.CodeName, authKey.AuthorizerAppId))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatTpAuthToken失败")
    }
    return gokits.NewCacheItem(key,
        ConfigInstance.WechatTpAuthTokenCacheDuration*time.Minute,
        gokits.UnJson(resp, new(WechatTpAuthTokenResp))), nil
}

func wechatCorpTokenLoader(codeName interface{}, _ ...interface{}) (*gokits.CacheItem, error) {
    resp, err := httpGet(gokits.PathJoin("/query-wechat-corp-token/", codeName.(string)))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatCorpToken失败")
    }
    return gokits.NewCacheItem(codeName,
        ConfigInstance.WechatCorpTokenCacheDuration*time.Minute,
        gokits.UnJson(resp, new(WechatCorpTokenResp))), nil
}

type WechatCorpTpAuthKey struct {
    CodeName string
    CorpId   string
}

func wechatCorpTpAuthTokenLoader(key interface{}, _ ...interface{}) (*gokits.CacheItem, error) {
    authKey, ok := key.(WechatCorpTpAuthKey)
    if !ok {
        return nil, errors.New("WechatCorpTpAuthKey type error")
    }
    resp, err := httpGet(gokits.PathJoin(
        "/query-wechat-corp-tp-auth-token/",
        authKey.CodeName, authKey.CorpId))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatCorpTpAuthToken失败")
    }
    return gokits.NewCacheItem(key,
        ConfigInstance.WechatCorpTpAuthTokenCacheDuration*time.Minute,
        gokits.UnJson(resp, new(WechatCorpTpAuthTokenResp))), nil
}

func toutiaoAppTokenLoader(codeName interface{}, _ ...interface{}) (*gokits.CacheItem, error) {
    resp, err := httpGet(gokits.PathJoin("/query-toutiao-app-token/", codeName.(string)))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求ToutiaoAppToken失败")
    }
    return gokits.NewCacheItem(codeName,
        ConfigInstance.ToutiaoAppTokenCacheDuration*time.Minute,
        gokits.UnJson(resp, new(ToutiaoAppTokenResp))), nil
}

func httpGet(subpath string) (string, error) {
    if 0 == len(ConfigInstance.Address) {
        return "", errors.New("未配置Varys.Address")
    }
    return gokits.NewHttpReq(ConfigInstance.Path(subpath)).Get()
}

func WechatAppToken(codeName string) (*WechatAppTokenResp, error) {
    value, err := wechatAppTokenCache.Value(codeName)
    if nil != err {
        return nil, err
    }
    return value.Data().(*WechatAppTokenResp), nil
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

func WechatMpLogin(codeName, jsCode string) (*WechatMpLoginResp, error) {
    resp, err := httpGet(gokits.PathJoin("/proxy-wechat-mp-login/", codeName) + "?js_code=" + jsCode)
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求WechatMpLogin失败")
    }
    return gokits.UnJson(resp, new(WechatMpLoginResp)).(*WechatMpLoginResp), nil
}
