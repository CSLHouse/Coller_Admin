package wechat

import (
	"context"
	wechatModel "cooller/server/model/wechat"
	"cooller/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderRecommendProduct = initOrderBrand + 1

type initRecommendProduct struct{}

// auto run
func init() {
	system.RegisterInit(initOrderRecommendProduct, &initRecommendProduct{})
}

func (i *initRecommendProduct) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.RecommendProduct{},
	)
}

func (i *initRecommendProduct) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.RecommendProduct{})
}

func (i initRecommendProduct) InitializerName() string {
	return wechatModel.RecommendProduct{}.TableName()
}

func (i *initRecommendProduct) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.RecommendProduct{
		{
			ProductId:       51,
			ProductName:     "Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）",
			RecommendStatus: 1,
			Sort:            0,
		},
		{
			ProductId:       6,
			ProductName:     "小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑(新R5-6600H标压 16G 512G win11)",
			RecommendStatus: 1,
			Sort:            0,
		},
		{
			ProductId:       49,
			ProductName:     "三星（SAMSUNG）500GB SSD固态硬盘 M.2接口(NVMe协议) 980（MZ-V8V500BW）",
			RecommendStatus: 1,
			Sort:            0,
		},
		{
			ProductId:       1,
			ProductName:     "万和（Vanward)燃气热水器天然气家用四重防冻直流变频节能全新升级增压水伺服恒温高抗风 JSQ30-565W16【16升】【恒温旗舰款】",
			RecommendStatus: 1,
			Sort:            0,
		},
		{
			ProductId:       6,
			ProductName:     "OPPO Reno8 8GB+128GB 鸢尾紫 新配色上市 80W超级闪充 5000万水光人像三摄 3200万前置索尼镜头 5G手机",
			RecommendStatus: 1,
			Sort:            0,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.RecommendProduct{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initRecommendProduct) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("product_id = ?", 51).First(&wechatModel.RecommendProduct{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
