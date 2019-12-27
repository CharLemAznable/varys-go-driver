package varys

import (
    "github.com/CharLemAznable/gokits"
    "strings"
    "time"
)

type Config struct {
    Address                                   string
    AppTokenCachePath                         string
    AppAuthorizerTokenCachePath               string
    CorpTokenCachePath                        string
    CorpAuthorizerTokenCachePath              string
    ProxyWechatAppPath                        string
    ProxyWechatCorpPath                       string
    AppTokenCacheDurationInMinutes            time.Duration
    AppAuthorizerTokenCacheDurationInMinutes  time.Duration
    CorpTokenCacheDurationInMinutes           time.Duration
    CorpAuthorizerTokenCacheDurationInMinutes time.Duration
}

type ConfigOptions struct {
    Address                                   string
    AppTokenCachePath                         string
    AppAuthorizerTokenCachePath               string
    CorpTokenCachePath                        string
    CorpAuthorizerTokenCachePath              string
    ProxyWechatAppPath                        string
    ProxyWechatCorpPath                       string
    AppTokenCacheDurationInMinutes            time.Duration
    AppAuthorizerTokenCacheDurationInMinutes  time.Duration
    CorpTokenCacheDurationInMinutes           time.Duration
    CorpAuthorizerTokenCacheDurationInMinutes time.Duration
}

var defaultConfigOptions = ConfigOptions{
    Address:                                   "http://127.0.0.1:4236",
    AppTokenCachePath:                         "/query-wechat-app-token/",
    AppAuthorizerTokenCachePath:               "/query-wechat-app-authorizer-token/",
    CorpTokenCachePath:                        "/query-wechat-corp-token/",
    CorpAuthorizerTokenCachePath:              "/query-wechat-corp-authorizer-token/",
    ProxyWechatAppPath:                        "/proxy-wechat-app/",
    ProxyWechatCorpPath:                       "/proxy-wechat-corp/",
    AppTokenCacheDurationInMinutes:            10,
    AppAuthorizerTokenCacheDurationInMinutes:  10,
    CorpTokenCacheDurationInMinutes:           10,
    CorpAuthorizerTokenCacheDurationInMinutes: 10,
}

type ConfigOption func(*ConfigOptions)

func WithAddress(address string) ConfigOption {
    return func(o *ConfigOptions) { o.Address = address }
}

func WithAppTokenCachePath(appTokenCachePath string) ConfigOption {
    return func(o *ConfigOptions) { o.AppTokenCachePath = appTokenCachePath }
}

func WithAppAuthorizerTokenCachePath(appAuthorizerTokenCachePath string) ConfigOption {
    return func(o *ConfigOptions) { o.AppAuthorizerTokenCachePath = appAuthorizerTokenCachePath }
}

func WithCorpTokenCachePath(corpTokenCachePath string) ConfigOption {
    return func(o *ConfigOptions) { o.CorpTokenCachePath = corpTokenCachePath }
}

func WithCorpAuthorizerTokenCachePath(corpAuthorizerTokenCachePath string) ConfigOption {
    return func(o *ConfigOptions) { o.CorpAuthorizerTokenCachePath = corpAuthorizerTokenCachePath }
}

func WithProxyWechatAppPath(proxyWechatAppPath string) ConfigOption {
    return func(o *ConfigOptions) { o.ProxyWechatAppPath = proxyWechatAppPath }
}

func WithProxyWechatCorpPath(proxyWechatCorpPath string) ConfigOption {
    return func(o *ConfigOptions) { o.ProxyWechatCorpPath = proxyWechatCorpPath }
}

func WithAppTokenCacheDurationInMinutes(appTokenCacheDurationInMinutes time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.AppTokenCacheDurationInMinutes = appTokenCacheDurationInMinutes }
}

func WithAppAuthorizerTokenCacheDurationInMinutes(appAuthorizerTokenCacheDurationInMinutes time.Duration) ConfigOption {
    return func(o *ConfigOptions) {
        o.AppAuthorizerTokenCacheDurationInMinutes = appAuthorizerTokenCacheDurationInMinutes
    }
}

func WithCorpTokenCacheDurationInMinutes(corpTokenCacheDurationInMinutes time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.CorpTokenCacheDurationInMinutes = corpTokenCacheDurationInMinutes }
}

func WithCorpAuthorizerTokenCacheDurationInMinutes(corpAuthorizerTokenCacheDurationInMinutes time.Duration) ConfigOption {
    return func(o *ConfigOptions) {
        o.CorpAuthorizerTokenCacheDurationInMinutes = corpAuthorizerTokenCacheDurationInMinutes
    }
}

func NewConfig(opts ...ConfigOption) *Config {
    options := defaultConfigOptions
    for _, o := range opts {
        o(&options)
    }
    return &Config{
        Address:                                   options.Address,
        AppTokenCachePath:                         options.AppTokenCachePath,
        AppAuthorizerTokenCachePath:               options.AppAuthorizerTokenCachePath,
        CorpTokenCachePath:                        options.CorpTokenCachePath,
        CorpAuthorizerTokenCachePath:              options.CorpAuthorizerTokenCachePath,
        ProxyWechatAppPath:                        options.ProxyWechatAppPath,
        ProxyWechatCorpPath:                       options.ProxyWechatCorpPath,
        AppTokenCacheDurationInMinutes:            options.AppTokenCacheDurationInMinutes,
        AppAuthorizerTokenCacheDurationInMinutes:  options.AppAuthorizerTokenCacheDurationInMinutes,
        CorpTokenCacheDurationInMinutes:           options.CorpTokenCacheDurationInMinutes,
        CorpAuthorizerTokenCacheDurationInMinutes: options.CorpAuthorizerTokenCacheDurationInMinutes,
    }
}

var ConfigInstance = NewConfig()

func (config *Config) Path(pathComponents ...string) string {
    if strings.HasSuffix(config.Address, "/") {
        config.Address = config.Address[:len(config.Address)-1]
    }

    subpath := gokits.PathJoin(pathComponents...)
    if !strings.HasPrefix(subpath, "/") {
        subpath = "/" + subpath
    }

    return config.Address + subpath
}
