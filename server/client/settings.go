package pay

import (
	"fmt"
	auth2 "github.com/flipped-aurora/gin-vue-admin/server/client/auth"
	"github.com/flipped-aurora/gin-vue-admin/server/client/cipher"
	"net/http"
)

// DialSettings 微信支付 API v3 Go SDK core.Client 需要的配置信息
type DialSettings struct {
	HTTPClient *http.Client    // 自定义所使用的 HTTPClient 实例
	Signer     auth2.Signer    // 签名器
	Validator  auth2.Validator // 应答包签名校验器
	Cipher     cipher.Cipher   // 敏感字段加解密套件
}

// Validate 校验请求配置是否有效
func (ds *DialSettings) Validate() error {
	if ds.Validator == nil {
		return fmt.Errorf("validator is required for Client")
	}
	if ds.Signer == nil {
		return fmt.Errorf("signer is required for Client")
	}
	return nil
}
