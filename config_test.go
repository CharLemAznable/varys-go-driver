package varys

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestConfigInstance(t *testing.T) {
    a := assert.New(t)

    ConfigInstance = NewConfig()
    a.Equal("http://127.0.0.1:4236", ConfigInstance.Address)
    a.Equal(time.Minute*10, ConfigInstance.WechatAppTokenCacheDuration)
    a.Equal(time.Minute*10, ConfigInstance.WechatTpTokenCacheDuration)
    a.Equal(time.Minute*10, ConfigInstance.WechatTpAuthTokenCacheDuration)
    a.Equal(time.Minute*10, ConfigInstance.WechatCorpTokenCacheDuration)
    a.Equal(time.Minute*10, ConfigInstance.WechatCorpTpAuthTokenCacheDuration)
    a.Equal(time.Minute*10, ConfigInstance.ToutiaoAppTokenCacheDuration)

    ConfigInstance = NewConfig(
        WithAddress("1"),
        WithWechatAppTokenCacheDuration(time.Minute*5),
        WithWechatTpTokenCacheDuration(time.Minute*5),
        WithWechatTpAuthTokenCacheDuration(time.Minute*5),
        WithWechatCorpTokenCacheDuration(time.Minute*5),
        WithWechatCorpTpAuthTokenCacheDuration(time.Minute*5),
        WithToutiaoAppTokenCacheDuration(time.Minute*5),
    )
    a.Equal("1", ConfigInstance.Address)
    a.Equal(time.Minute*5, ConfigInstance.WechatAppTokenCacheDuration)
    a.Equal(time.Minute*5, ConfigInstance.WechatTpTokenCacheDuration)
    a.Equal(time.Minute*5, ConfigInstance.WechatTpAuthTokenCacheDuration)
    a.Equal(time.Minute*5, ConfigInstance.WechatCorpTokenCacheDuration)
    a.Equal(time.Minute*5, ConfigInstance.WechatCorpTpAuthTokenCacheDuration)
    a.Equal(time.Minute*5, ConfigInstance.ToutiaoAppTokenCacheDuration)
}

func TestConfigPath(t *testing.T) {
    a := assert.New(t)

    ConfigInstance = NewConfig()
    a.Equal("http://127.0.0.1:4236/abc", ConfigInstance.Path("/abc"))
    a.Equal("http://127.0.0.1:4236/abc", ConfigInstance.Path("abc"))

    ConfigInstance = NewConfig(WithAddress("http://127.0.0.1:4236/"))
    a.Equal("http://127.0.0.1:4236/abc", ConfigInstance.Path("/abc"))
    a.Equal("http://127.0.0.1:4236/abc", ConfigInstance.Path("abc"))
}
