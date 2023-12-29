package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderHomeProductAttribute = initOrderHomeProductAttributeCategory + 1

type initHomeProductAttribute struct{}

// auto run
func init() {
	system.RegisterInit(initOrderHomeProductAttribute, &initHomeProductAttribute{})
}

func (i *initHomeProductAttribute) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.ProductAttribute{},
		&wechatModel.ProductAttributeValue{},
	)
}

func (i *initHomeProductAttribute) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.ProductAttribute{})
}

func (i initHomeProductAttribute) InitializerName() string {
	return wechatModel.ProductAttribute{}.TableName()
}

func (i *initHomeProductAttribute) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.ProductAttribute{
		{
			ProductAttributeCategoryId: 1,
			Name:                       "尺寸",
			SelectType:                 2,
			InputType:                  1,
			InputList:                  "M,X,XL,2XL,3XL,4XL",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 1,
			Name:                       "颜色",
			SelectType:                 2,
			InputType:                  1,
			InputList:                  "黑色,红色,白色,粉色",
			Sort:                       100,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              1,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 1,
			Name:                       "适用季节",
			SelectType:                 1,
			InputType:                  1,
			InputList:                  "春季,夏季,秋季,冬季",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 2,
			Name:                       "适用人群",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "老年,青年,中年",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 2,
			Name:                       "风格",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "嘻哈风格,基础大众,商务正装",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 2,
			Name:                       "颜色",
			SelectType:                 0,
			InputType:                  0,
			InputList:                  "",
			Sort:                       100,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              1,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 1,
			Name:                       "袖长",
			SelectType:                 1,
			InputType:                  1,
			InputList:                  "短袖,长袖,中袖",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 2,
			Name:                       "尺码",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "29,30,31,32,33,34",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 2,
			Name:                       "适用场景",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "居家,运动,正装",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.ProductAttribute{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initHomeProductAttribute) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "版本").First(&wechatModel.ProductAttribute{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
