package varys

import (
    "errors"
    "github.com/CharLemAznable/gokits"
    "strings"
    "time"
)

type Config struct {
    Address                            string
    WechatAppTokenCacheDuration        time.Duration
    WechatTpTokenCacheDuration         time.Duration
    WechatTpAuthTokenCacheDuration     time.Duration
    WechatCorpTokenCacheDuration       time.Duration
    WechatCorpTpAuthTokenCacheDuration time.Duration
    ToutiaoAppTokenCacheDuration       time.Duration
    FengniaoAppTokenCacheDuration      time.Duration
}

type ConfigOptions struct {
    Address                            string
    WechatAppTokenCacheDuration        time.Duration
    WechatTpTokenCacheDuration         time.Duration
    WechatTpAuthTokenCacheDuration     time.Duration
    WechatCorpTokenCacheDuration       time.Duration
    WechatCorpTpAuthTokenCacheDuration time.Duration
    ToutiaoAppTokenCacheDuration       time.Duration
    FengniaoAppTokenCacheDuration      time.Duration
}

var defaultConfigOptions = ConfigOptions{
    Address:                            "http://127.0.0.1:4236",
    WechatAppTokenCacheDuration:        time.Minute * 10,
    WechatTpTokenCacheDuration:         time.Minute * 10,
    WechatTpAuthTokenCacheDuration:     time.Minute * 10,
    WechatCorpTokenCacheDuration:       time.Minute * 10,
    WechatCorpTpAuthTokenCacheDuration: time.Minute * 10,
    ToutiaoAppTokenCacheDuration:       time.Minute * 10,
    FengniaoAppTokenCacheDuration:      time.Minute * 10,
}

type ConfigOption func(*ConfigOptions)

func WithAddress(address string) ConfigOption {
    return func(o *ConfigOptions) { o.Address = address }
}

func WithWechatAppTokenCacheDuration(wechatAppTokenCacheDuration time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.WechatAppTokenCacheDuration = wechatAppTokenCacheDuration }
}

func WithWechatTpTokenCacheDuration(wechatTpTokenCacheDuration time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.WechatTpTokenCacheDuration = wechatTpTokenCacheDuration }
}

func WithWechatTpAuthTokenCacheDuration(wechatTpAuthTokenCacheDuration time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.WechatTpAuthTokenCacheDuration = wechatTpAuthTokenCacheDuration }
}

func WithWechatCorpTokenCacheDuration(wechatCorpTokenCacheDuration time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.WechatCorpTokenCacheDuration = wechatCorpTokenCacheDuration }
}

func WithWechatCorpTpAuthTokenCacheDuration(wechatCorpTpAuthTokenCacheDuration time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.WechatCorpTpAuthTokenCacheDuration = wechatCorpTpAuthTokenCacheDuration }
}

func WithToutiaoAppTokenCacheDuration(toutiaoAppTokenCacheDuration time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.ToutiaoAppTokenCacheDuration = toutiaoAppTokenCacheDuration }
}

func WithFengniaoAppTokenCacheDuration(fengniaoAppTokenCacheDuration time.Duration) ConfigOption {
    return func(o *ConfigOptions) { o.FengniaoAppTokenCacheDuration = fengniaoAppTokenCacheDuration }
}

func NewConfig(opts ...ConfigOption) *Config {
    options := defaultConfigOptions
    for _, o := range opts {
        o(&options)
    }
    return &Config{
        Address:                            options.Address,
        WechatAppTokenCacheDuration:        options.WechatAppTokenCacheDuration,
        WechatTpTokenCacheDuration:         options.WechatTpTokenCacheDuration,
        WechatTpAuthTokenCacheDuration:     options.WechatTpAuthTokenCacheDuration,
        WechatCorpTokenCacheDuration:       options.WechatCorpTokenCacheDuration,
        WechatCorpTpAuthTokenCacheDuration: options.WechatCorpTpAuthTokenCacheDuration,
        ToutiaoAppTokenCacheDuration:       options.ToutiaoAppTokenCacheDuration,
        FengniaoAppTokenCacheDuration:      options.FengniaoAppTokenCacheDuration,
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

func httpGet(subpath string) (string, error) {
    if 0 == len(ConfigInstance.Address) {
        return "", errors.New("未配置Varys.Address")
    }
    return gokits.NewHttpReq(ConfigInstance.Path(subpath)).Get()
}
