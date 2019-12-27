package varys

import (
	"errors"
	"github.com/CharLemAznable/gokits"
	"time"
)

type AppTokenResp struct {
	Token string `json:"token"`
	Error string `json:"error"`
	AppId string `json:"appId"`
}

type AppAuthorizerTokenResp struct {
	Token           string `json:"token"`
	Error           string `json:"error"`
	AppId           string `json:"appId"`
	AuthorizerAppId string `json:"authorizerAppId"`
}

type CorpTokenResp struct {
	Token  string `json:"token"`
	Error  string `json:"error"`
	CorpId string `json:"corpId"`
}

type CorpAuthorizerTokenResp struct {
	Token   string `json:"token"`
	Error   string `json:"error"`
	SuiteId string `json:"suiteId"`
	CorpId  string `json:"corpId"`
}

var appTokenCache *gokits.CacheTable
var appAuthorizerTokenCache *gokits.CacheTable
var corpTokenCache *gokits.CacheTable
var corpAuthorizerTokenCache *gokits.CacheTable

func init() {
	appTokenCache = gokits.CacheExpireAfterWrite("appTokenCache")
	appTokenCache.SetDataLoader(appTokenLoader)
	appAuthorizerTokenCache = gokits.CacheExpireAfterWrite("appAuthorizerTokenCache")
	appAuthorizerTokenCache.SetDataLoader(appAuthorizerTokenLoader)
	corpTokenCache = gokits.CacheExpireAfterWrite("corpTokenCache")
	corpTokenCache.SetDataLoader(corpTokenLoader)
	corpAuthorizerTokenCache = gokits.CacheExpireAfterWrite("corpAuthorizerTokenCache")
	corpAuthorizerTokenCache.SetDataLoader(corpAuthorizerTokenLoader)
}

//noinspection GoUnusedParameter
func appTokenLoader(codeName interface{}, args ...interface{}) (*gokits.CacheItem, error) {
	resp, err := httpGet(gokits.PathJoin(ConfigInstance.AppTokenCachePath, codeName.(string)))
	if nil != err {
		return nil, err
	}
	if 0 == len(resp) {
		return nil, errors.New("请求AppToken失败")
	}
	return gokits.NewCacheItem(codeName,
		ConfigInstance.AppTokenCacheDurationInMinutes*time.Minute,
		gokits.UnJson(resp, new(AppTokenResp))), nil
}

type AppAuthorizerTokenKey struct {
	CodeName        string
	AuthorizerAppId string
}

//noinspection GoUnusedParameter
func appAuthorizerTokenLoader(key interface{}, args ...interface{}) (*gokits.CacheItem, error) {
	tokenKey, ok := key.(AppAuthorizerTokenKey)
	if !ok {
		return nil, errors.New("AppAuthorizerTokenKey type error")
	}
	resp, err := httpGet(gokits.PathJoin(
		ConfigInstance.AppAuthorizerTokenCachePath,
		tokenKey.CodeName, tokenKey.AuthorizerAppId))
	if nil != err {
		return nil, err
	}
	if 0 == len(resp) {
		return nil, errors.New("请求AppAuthorizerToken失败")
	}
	return gokits.NewCacheItem(key,
		ConfigInstance.AppAuthorizerTokenCacheDurationInMinutes*time.Minute,
		gokits.UnJson(resp, new(AppAuthorizerTokenResp))), nil
}

//noinspection GoUnusedParameter
func corpTokenLoader(codeName interface{}, args ...interface{}) (*gokits.CacheItem, error) {
	resp, err := httpGet(gokits.PathJoin(ConfigInstance.CorpTokenCachePath, codeName.(string)))
	if nil != err {
		return nil, err
	}
	if 0 == len(resp) {
		return nil, errors.New("请求CorpToken失败")
	}
	return gokits.NewCacheItem(codeName,
		ConfigInstance.CorpTokenCacheDurationInMinutes*time.Minute,
		gokits.UnJson(resp, new(CorpTokenResp))), nil
}

type CorpAuthorizerTokenKey struct {
	CodeName string
	CorpId   string
}

//noinspection GoUnusedParameter
func corpAuthorizerTokenLoader(key interface{}, args ...interface{}) (*gokits.CacheItem, error) {
	tokenKey, ok := key.(CorpAuthorizerTokenKey)
	if !ok {
		return nil, errors.New("CorpAuthorizerTokenKey type error")
	}
	resp, err := httpGet(gokits.PathJoin(
		ConfigInstance.CorpAuthorizerTokenCachePath,
		tokenKey.CodeName, tokenKey.CorpId))
	if nil != err {
		return nil, err
	}
	if 0 == len(resp) {
		return nil, errors.New("请求CorpAuthorizerToken失败")
	}
	return gokits.NewCacheItem(key,
		ConfigInstance.CorpAuthorizerTokenCacheDurationInMinutes*time.Minute,
		gokits.UnJson(resp, new(CorpAuthorizerTokenResp))), nil
}

func httpGet(subpath string) (string, error) {
	if 0 == len(ConfigInstance.Address) {
		return "", errors.New("未配置Varys.Address")
	}
	return gokits.NewHttpReq(ConfigInstance.Path(subpath)).Get()
}

func AppToken(codeName string) (*AppTokenResp, error) {
	cache, err := appTokenCache.Value(codeName)
	if nil != err {
		return nil, err
	}
	return cache.Data().(*AppTokenResp), nil
}

func AppAuthorizerToken(codeName, authorizerAppId string) (*AppAuthorizerTokenResp, error) {
	cache, err := appAuthorizerTokenCache.Value(AppAuthorizerTokenKey{
		CodeName: codeName, AuthorizerAppId: authorizerAppId})
	if nil != err {
		return nil, err
	}
	return cache.Data().(*AppAuthorizerTokenResp), err
}

func CorpToken(codeName string) (*CorpTokenResp, error) {
	cache, err := corpTokenCache.Value(codeName)
	if nil != err {
		return nil, err
	}
	return cache.Data().(*CorpTokenResp), nil
}

func CorpAuthorizerToken(codeName, corpId string) (*CorpAuthorizerTokenResp, error) {
	cache, err := corpAuthorizerTokenCache.Value(CorpAuthorizerTokenKey{
		CodeName: codeName, CorpId: corpId})
	if nil != err {
		return nil, err
	}
	return cache.Data().(*CorpAuthorizerTokenResp), nil
}
