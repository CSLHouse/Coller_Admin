package system

import (
	"context"
	. "cooller/server/model/system"
	"cooller/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderAuthority + 1

type initMenu struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysBaseMenu{},
		&SysBaseMenuParameter{},
		&SysBaseMenuBtn{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{}) &&
		m.HasTable(&SysBaseMenuParameter{}) &&
		m.HasTable(&SysBaseMenuBtn{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	// ParentId是menu在数据库中的id顺序，0-表示在根目录
	entities := []SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 9, Meta: Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 2, Meta: Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		//{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "员工管理", Icon: "coordinate"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: Meta{Title: "字典管理", Icon: "notebook"}},
		//{MenuLevel: 0, Hidden: true, ParentId: "2", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: Meta{Title: "字典详情-${id}", Icon: "list", ActiveName: "dictionary"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: Meta{Title: "操作历史", Icon: "pie-chart"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 7, Meta: Meta{Title: "示例文件", Icon: "management"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: Meta{Title: "断点续传", Icon: "upload-filled"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: Meta{Title: "客户列表（资源示例）", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 4, Meta: Meta{Title: "系统工具", Icon: "tools"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "16", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		//{MenuLevel: 0, Hidden: false, ParentId: "16", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "8", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 1, Meta: Meta{Title: "系统管理", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: "8", Path: "setting", Name: "setting", Component: "view/systemTools/setting/setting.vue", Sort: 2, Meta: Meta{Title: "系统配置", Icon: "setting"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "16", Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 1, Meta: Meta{Title: "自动化代码管理", Icon: "magic-stick"}},
		//{MenuLevel: 0, Hidden: true, ParentId: "16", Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: Meta{Title: "自动化代码-${id}", Icon: "magic-stick"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "16", Path: "autoPkg", Name: "autoPkg", Component: "view/systemTools/autoPkg/autoPkg.vue", Sort: 0, Meta: Meta{Title: "自动化package", Icon: "folder"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: Meta{Title: "官方网站", Icon: "home-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 8, Meta: Meta{Title: "服务器状态", Icon: "cloudy"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 6, Meta: Meta{Title: "插件系统", Icon: "cherry"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "https://plugin.gin-vue-admin.com/", Name: "https://plugin.gin-vue-admin.com/", Component: "https://plugin.gin-vue-admin.com/", Sort: 0, Meta: Meta{Title: "插件市场", Icon: "shop"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "installPlugin", Name: "installPlugin", Component: "view/systemTools/installPlugin/index.vue", Sort: 1, Meta: Meta{Title: "插件安装", Icon: "box"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "autoPlug", Name: "autoPlug", Component: "view/systemTools/autoPlug/autoPlug.vue", Sort: 2, Meta: Meta{Title: "插件模板", Icon: "folder"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "pubPlug", Name: "pubPlug", Component: "view/systemTools/pubPlug/pubPlug.vue", Sort: 3, Meta: Meta{Title: "打包插件", Icon: "files"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "plugin-email", Name: "plugin-email", Component: "plugin/email/view/index.vue", Sort: 4, Meta: Meta{Title: "邮件插件", Icon: "message"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "16", Path: "chatTable", Name: "chatTable", Component: "view/chatgpt/chatTable.vue", Sort: 6, Meta: Meta{Title: "万用表格", Icon: "chat-dot-square"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "members", Name: "members", Component: "view/members/index.vue", Sort: 2, Meta: Meta{Title: "会员管理", Icon: "chat-dot-square"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "newMember", Name: "newMember", Component: "view/members/newMember/newMember.vue", Sort: 1, Meta: Meta{Title: "新会员办理", Icon: "shop"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "oldMember", Name: "oldMember", Component: "view/members/oldMember/oldMember.vue", Sort: 2, Meta: Meta{Title: "老会员续卡", Icon: "shopping-cart"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "showMembers", Name: "showMembers", Component: "view/members/showMembers/showMembers.vue", Sort: 3, Meta: Meta{Title: "会员查看", Icon: "shopping-cart-full"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "vipCombo", Name: "vipCombo", Component: "view/members/vipCombo/vipCombo.vue", Sort: 4, Meta: Meta{Title: "Vip套餐管理", Icon: "set-up"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "statistics", Name: "Statistics", Component: "view/members/Statistics/Statistics.vue", Sort: 5, Meta: Meta{Title: "数据统计", Icon: "Collection"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "orderView", Name: "orderView", Component: "view/members/orderView/orderView.vue", Sort: 6, Meta: Meta{Title: "会员订单", Icon: "Reading"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "shop", Name: "shop", Component: "view/shop/index.vue", Sort: 4, Meta: Meta{Title: "门店管理", Icon: "Shop"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "consumeRegister", Name: "consumeRegister", Component: "view/shop/consumeRegister/consumeRegister.vue", Sort: 1, Meta: Meta{Title: "消费登记", Icon: "ShoppingCart"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "qrcode", Name: "qrcode", Component: "view/shop/qrcode/qrcode.vue", Sort: 2, Meta: Meta{Title: "二维码", Icon: "Reading"}},
		{MenuLevel: 0, Hidden: false, ParentId: "19", Path: "consumeList", Name: "consumeList", Component: "view/shop/consumeList/consumeList.vue", Sort: 3, Meta: Meta{Title: "消费列表", Icon: "Reading"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "orderManage", Name: "orderManage", Component: "view/orderManage/index.vue", Sort: 5, Meta: Meta{Title: "订单管理", Icon: "Management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "23", Path: "orderList", Name: "orderList", Component: "view/orderManage/orderList/orderList.vue", Sort: 1, Meta: Meta{Title: "订单列表", Icon: "Collection"}},
		{MenuLevel: 0, Hidden: false, ParentId: "23", Path: "orderSetting", Name: "orderSetting", Component: "view/orderManage/orderSetting/orderSetting.vue", Sort: 2, Meta: Meta{Title: "订单设置", Icon: "Reading"}},
		{MenuLevel: 0, Hidden: true, ParentId: "23", Path: "orderDetail", Name: "orderDetail", Component: "view/orderManage/orderDetail/orderDetail.vue", Sort: 3, Meta: Meta{Title: "订单详情", Icon: "Reading"}},

		//{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "product", Name: "Product", Component: "view/product/index.vue", Sort: 6, Meta: Meta{Title: "商品", Icon: "Goods"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "productList", Name: "ProductList", Component: "view/product/productList/ProductList.vue", Sort: 1, Meta: Meta{Title: "商品列表", Icon: "Box"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "newProduct", Name: "NewProduct", Component: "view/product/newProduct/NewProduct.vue", Sort: 2, Meta: Meta{Title: "添加商品", Icon: "Sell"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "productCategories", Name: "ProductCategories", Component: "view/product/productCategories/ProductCategories.vue", Sort: 3, Meta: Meta{Title: "商品分类", Icon: "Paperclip"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "productType", Name: "ProductType", Component: "view/product/productType/ProductType.vue", Sort: 4, Meta: Meta{Title: "商品类型", Icon: "MagicStick"}},
		//{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "brand", Name: "Brand", Component: "view/product/brand/Brand.vue", Sort: 5, Meta: Meta{Title: "品牌管理", Icon: "Medal"}},
		//{MenuLevel: 0, Hidden: true, ParentId: "27", Path: "attribute", Name: "Attribute", Component: "view/product/attribute/Attribute.vue", Sort: 6, Meta: Meta{Title: "商品属性参数", Icon: "Medal"}},
		//{MenuLevel: 0, Hidden: true, ParentId: "27", Path: "updateProduct", Name: "updateProduct", Component: "view/product/updateProduct/updateProduct.vue", Sort: 7, Meta: Meta{Title: "修改商品", Icon: "Sell"}},.

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "product", Name: "Product", Component: "view/product/index.vue", Sort: 6, Meta: Meta{Title: "商品", Icon: "Goods"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "list", Name: "list", Component: "view/product/list/index.vue", Sort: 1, Meta: Meta{Title: "商品列表", Icon: "Box"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "add", Name: "add", Component: "view/product/add/add.vue", Sort: 2, Meta: Meta{Title: "添加商品", Icon: "Sell"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "productCategories", Name: "ProductCategories", Component: "view/product/productCategories/ProductCategories.vue", Sort: 3, Meta: Meta{Title: "商品分类", Icon: "Paperclip"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "productType", Name: "ProductType", Component: "view/product/productType/ProductType.vue", Sort: 4, Meta: Meta{Title: "商品类型", Icon: "MagicStick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "brand", Name: "Brand", Component: "view/product/brand/Brand.vue", Sort: 5, Meta: Meta{Title: "品牌管理", Icon: "Medal"}},
		{MenuLevel: 0, Hidden: true, ParentId: "27", Path: "attribute", Name: "Attribute", Component: "view/product/attribute/Attribute.vue", Sort: 6, Meta: Meta{Title: "商品属性参数", Icon: "Medal"}},
		{MenuLevel: 0, Hidden: true, ParentId: "27", Path: "update", Name: "updateProduct", Component: "view/product/update/update.vue", Sort: 7, Meta: Meta{Title: "修改商品", Icon: "Sell"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "marketing", Name: "Marketing", Component: "view/marketing/index.vue", Sort: 7, Meta: Meta{Title: "营销", Icon: "Van"}},
		{MenuLevel: 0, Hidden: false, ParentId: "35", Path: "flashPromotion", Name: "FlashPromotion", Component: "view/marketing/flashPromotion/index.vue", Sort: 1, Meta: Meta{Title: "秒杀活动列表", Icon: "Clock"}},
		{MenuLevel: 0, Hidden: true, ParentId: "35", Path: "productRelation", Name: "productRelation", Component: "view/marketing/flashPromotion/productRelationList.vue", Sort: 2, Meta: Meta{Title: "秒杀商品列表", Icon: "Present"}},
		{MenuLevel: 0, Hidden: true, ParentId: "35", Path: "selectSession", Name: "selectSession", Component: "view/marketing/flashPromotion/selectSessionList.vue", Sort: 3, Meta: Meta{Title: "秒杀时间段选择", Icon: "Box"}},
		{MenuLevel: 0, Hidden: true, ParentId: "35", Path: "session", Name: "session", Component: "view/marketing/flashPromotion/sessionList.vue", Sort: 4, Meta: Meta{Title: "秒杀时间段列表", Icon: "Box"}},
		{MenuLevel: 0, Hidden: false, ParentId: "35", Path: "recommend", Name: "RecommendProduct", Component: "view/marketing/recommend/RecommendProduct.vue", Sort: 5, Meta: Meta{Title: "人气推荐", Icon: "Medal"}},
		{MenuLevel: 0, Hidden: false, ParentId: "35", Path: "brandRecommend", Name: "BrandRecommend", Component: "view/marketing/brandRecommend/brandRecommend.vue", Sort: 6, Meta: Meta{Title: "品牌推荐", Icon: "Sell"}},
		{MenuLevel: 0, Hidden: false, ParentId: "35", Path: "advertise", Name: "Advertise", Component: "view/marketing/advertise/Advertise.vue", Sort: 7, Meta: Meta{Title: "广告列表", Icon: "MagicStick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "35", Path: "newRecommend", Name: "NewRecommend", Component: "view/marketing/newRecommend/NewRecommend.vue", Sort: 8, Meta: Meta{Title: "新品推荐", Icon: "Medal"}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "person").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
