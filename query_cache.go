package varys

import (
    "errors"
    "github.com/CharLemAznable/gokits"
)

var wechatAppTokenCache *gokits.CacheTable
var wechatTpTokenCache *gokits.CacheTable
var wechatTpAuthTokenCache *gokits.CacheTable
var wechatCorpTokenCache *gokits.CacheTable
var wechatCorpTpAuthTokenCache *gokits.CacheTable
var toutiaoAppTokenCache *gokits.CacheTable
var fengniaoAppTokenCache *gokits.CacheTable

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
    fengniaoAppTokenCache = gokits.CacheExpireAfterWrite("FengniaoAppTokenCache")
    fengniaoAppTokenCache.SetDataLoader(fengniaoAppTokenLoader)
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
        ConfigInstance.WechatAppTokenCacheDuration,
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
        ConfigInstance.WechatTpTokenCacheDuration,
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
        ConfigInstance.WechatTpAuthTokenCacheDuration,
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
        ConfigInstance.WechatCorpTokenCacheDuration,
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
        ConfigInstance.WechatCorpTpAuthTokenCacheDuration,
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
        ConfigInstance.ToutiaoAppTokenCacheDuration,
        gokits.UnJson(resp, new(ToutiaoAppTokenResp))), nil
}

func fengniaoAppTokenLoader(codeName interface{}, _ ...interface{}) (*gokits.CacheItem, error) {
    resp, err := httpGet(gokits.PathJoin("/query-fengniao-app-token/", codeName.(string)))
    if nil != err {
        return nil, err
    }
    if 0 == len(resp) {
        return nil, errors.New("请求FengniaoAppToken失败")
    }
    return gokits.NewCacheItem(codeName,
        ConfigInstance.FengniaoAppTokenCacheDuration,
        gokits.UnJson(resp, new(FengniaoAppTokenResp))), nil
}
