package system

import (
	"context"

	"cooller/server/service/system"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderCasbin = initOrderApi + 1

type initCasbin struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbin, &initCasbin{})
}

func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i initCasbin) InitializerName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "888", V1: "/user/admin_register", V2: "POST"},

		//{Ptype: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/api/deleteApisByIds", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuthorities", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/resetPassword", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/findFile", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/breakpointContinueFinish", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/breakpointContinue", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/removeChunk", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/system/getServerInfo", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/customer/customerList", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getMeta", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/preview", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getColumn", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/autoCode/rollback", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/delSysHistory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getSysHistory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/createPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/getPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/delPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/createPlug", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/installPlugin", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/autoCode/pubPlug", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/sysDictionary/findSysDictionary", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/updateSysDictionary", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/getSysDictionaryList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/createSysDictionary", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysDictionary/deleteSysDictionary", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/email/emailTest", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/simpleUploader/upload", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/simpleUploader/checkFileMd5", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/simpleUploader/mergeFileMd5", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/authorityBtn/setAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authorityBtn/getAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authorityBtn/canRemoveAuthorityBtn", V2: "POST"},

		//{Ptype: "p", V0: "888", V1: "/chatGpt/getTable", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/chatGpt/createSK", V2: "POST"},
		//{Ptype: "p", V0: "888", V1: "/chatGpt/getSK", V2: "GET"},
		//{Ptype: "p", V0: "888", V1: "/chatGpt/deleteSK", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: "/business/combo", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/business/combo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/business/combo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/business/combo", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/business/comboList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/business/allCombo", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/business/member", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/business/member", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/business/member", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/business/member", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/business/memberList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/business/memberSearch", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/business/searchCard", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/business/renew", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/business/consume", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/business/consumeList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/business/orderList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/business/statistics", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/business/statement", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/product/advertise", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/content", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/advertise", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/advertise", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/advertiseState", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/updateKeyword", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/advertiseList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/recommendProductList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/brand", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/brand", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/brand", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/brand", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/brandState", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/recommendProduct", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/recommendProduct", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/recommendProduct", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/recommendProduct", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/brandDetail", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/productDetail", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/attributeCategory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/attributeCategory", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/attributeCategory", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/attributeCategory", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/attribute", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/attribute", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/attribute", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/attribute", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/productCategory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/productCategory", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/productCategory", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/productCategory", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/allCategory", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/cart", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/product/cart", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/cart/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/product/sku", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/product/cart/clear", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/deletes", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/product/updateRecommendSort", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/flash/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/flash/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/flash/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/flash/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/flash/updateStatus", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/flash/createFlashProductRelation", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/flash/deleteFlashProductRelation", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/flash/flashProductRelationList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/flash/updateFlashProductRelation", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/flash/createFlashSession", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/flash/deleteFlashSession", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/flash/flashSessionList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/flash/updateFlashSession", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/flash/updateFlashSessionStatus", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/flash/flashSessionSelectList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/qrcode/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/qrcode/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/qrcode/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/qrcode/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/qrcode/download", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/qrcode/scan", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/order/detail", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/order/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/order/cancelOrder", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/order/closeOrders", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/order/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/order/setting", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/order/settingUpdate", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/order/update/receiverInfo", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/order/update/moneyInfo", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/order/update/note", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/order/update/complete", V2: "POST"},

		{Ptype: "p", V0: "8881", V1: "/user/admin_register", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/api/getAllApis", V2: "POST"},
		//{Ptype: "p", V0: "8881", V1: "/authority/createAuthority", V2: "POST"},
		//{Ptype: "p", V0: "8881", V1: "/authority/deleteAuthority", V2: "POST"},
		//{Ptype: "p", V0: "8881", V1: "/authority/getAuthorityList", V2: "POST"},
		//{Ptype: "p", V0: "8881", V1: "/authority/setDataAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getBaseMenuTree", V2: "POST"},
		//{Ptype: "p", V0: "8881", V1: "/menu/addMenuAuthority", V2: "POST"},
		//{Ptype: "p", V0: "8881", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/customer/customerList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/user/getUserInfo", V2: "GET"},

		{Ptype: "p", V0: "8881", V1: "/business/combo", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/business/combo", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/business/combo", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/business/combo", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/business/comboList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/business/allCombo", V2: "GET"},

		{Ptype: "p", V0: "8881", V1: "/business/member", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/business/member", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/business/member", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/business/member", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/business/memberList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/business/memberSearch", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/business/searchCard", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/business/renew", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/business/consume", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/business/consumeList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/business/orderList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/business/statistics", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/business/statement", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/product/deletes", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/product/allCategory", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/product/create", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/product/update", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/product/attributeCategory", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/product/attributeCategory", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/product/attributeCategory", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/product/attributeCategory", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/product/recommendProduct", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/product/updateRecommendSort", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/flash/create", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/flash/delete", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/flash/list", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/flash/update", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/flash/updateStatus", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/flash/createFlashProductRelation", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/flash/deleteFlashProductRelation", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/flash/flashProductRelationList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/flash/updateFlashProductRelation", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/flash/createFlashSession", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/flash/deleteFlashSession", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/flash/flashSessionList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/flash/updateFlashSession", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/flash/updateFlashSessionStatus", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/flash/flashSessionSelectList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/qrcode/create", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/qrcode/delete", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/qrcode/list", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/qrcode/update", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: "/order/detail", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/order/list", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: "/order/cancelOrder", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/order/closeOrders", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/order/delete", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: "/order/update/receiverInfo", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/order/update/moneyInfo", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/order/update/note", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: "/order/update/complete", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: "/user/admin_register", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},

		//{Ptype: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
		//{Ptype: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/customer/customerList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/autoCode/createTemp", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/user/resetNickName", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: "/business/combo", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/business/combo", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/combo", V2: "PUT"},
		{Ptype: "p", V0: "9528", V1: "/business/combo", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/business/comboList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/allCombo", V2: "GET"},

		{Ptype: "p", V0: "9528", V1: "/business/member", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/business/member", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/member", V2: "PUT"},
		{Ptype: "p", V0: "9528", V1: "/business/member", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/business/memberList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/memberSearch", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/business/searchCard", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/business/cardList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/certificateList", V2: "GET"},

		{Ptype: "p", V0: "9528", V1: "/business/renew", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/business/consume", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/business/consumeList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/orderList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/statistics", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/business/statement", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/order/cart", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/order/cart", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/order/carts", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/order/cart/list", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/order/cart/clear", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/order/cart", V2: "PUT"},
		{Ptype: "p", V0: "9528", V1: "/order/generateConfirmOrder", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/pay/generateOrder", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/pay/detail", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/order/generateOrder", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/order/detail", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/order/list", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/order/paySuccess", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/order/cancelOrder", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/order/tmpCart", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: "/base/address", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/base/address", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/base/address", V2: "PUT"},
		{Ptype: "p", V0: "9528", V1: "/base/address", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: "/base/addressList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/base/phoneNumber", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: "/base/checkPhone", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: "/base/recordShare", V2: "POST"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitializerName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
