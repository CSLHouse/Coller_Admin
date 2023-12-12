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
			ProductAttributeCategoryId: 0,
			Name:                       "上市年份",
			SelectType:                 1,
			InputType:                  1,
			InputList:                  "2013年,2014年,2015年,2016年,2017年",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 0,
			Name:                       "适用对象",
			SelectType:                 1,
			InputType:                  1,
			InputList:                  "青年女性,中年女性",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 1,
			Name:                       "商品编号",
			SelectType:                 1,
			InputType:                  0,
			InputList:                  "",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
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
			Name:                       "上市时间",
			SelectType:                 1,
			InputType:                  1,
			InputList:                  "2017年秋,2017年冬,2018年春,2018年夏",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
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
		{
			ProductAttributeCategoryId: 2,
			Name:                       "上市时间",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "2018年春,2018年夏",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 3,
			Name:                       "容量",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "16G,32G,64G,128G,256G,512G",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 3,
			Name:                       "屏幕尺寸",
			SelectType:                 0,
			InputType:                  0,
			InputList:                  "",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 3,
			Name:                       "网络",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "3G,4G,5G,WLAN",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 3,
			Name:                       "系统",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "Android,IOS",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 3,
			Name:                       "电池容量",
			SelectType:                 0,
			InputType:                  0,
			InputList:                  "",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 11,
			Name:                       "颜色",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "红色,蓝色,绿色",
			Sort:                       0,
			FilterType:                 1,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 11,
			Name:                       "尺寸",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "38,39,40",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 11,
			Name:                       "风格",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "夏季,秋季",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 12,
			Name:                       "尺寸",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "50英寸,65英寸,70英寸",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 12,
			Name:                       "内存",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "8G,16G,32G",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       0,
		},
		{
			ProductAttributeCategoryId: 12,
			Name:                       "商品毛重",
			SelectType:                 0,
			InputType:                  0,
			InputList:                  "",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 12,
			Name:                       "商品产地",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "中国大陆,其他",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 12,
			Name:                       "电视类型",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "大屏,教育电视,4K超清",
			Sort:                       0,
			FilterType:                 0,
			SearchType:                 0,
			RelatedStatus:              0,
			HandAddStatus:              0,
			Type:                       1,
		},
		{
			ProductAttributeCategoryId: 13,
			Name:                       "版本",
			SelectType:                 0,
			InputType:                  1,
			InputList:                  "R7 16G 512,R5 16G 512,I5 16G 512,I7 16G 512",
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
