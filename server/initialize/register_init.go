package initialize

import (
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/business"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/example"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/system"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/wechat"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
