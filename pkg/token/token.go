package token

import (
	"errors"
	"sync"
)

// Config 包括 token 包的配置选项
type Config struct {
	key         string
	identityKey string
}

// ErrMissingHeader 表示 `Authorization` 请求头为空
var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

var (
	config = Config{"Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5", "identityKey"}
	once   sync.Once
)

// Init 设置包级别的配置 config, config 会用于本包后面的 token 签发和解析
func Init(key string, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

// Parse 使用指定的密钥 key 解析 token，解析成功返回 token 上下文，否则报错.
func Parse(tokenString string, key string) {

}
