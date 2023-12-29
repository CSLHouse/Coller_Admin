package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderHomeProductCategory = initOrderHomeProduct + 1

type initHomeProductCategory struct{}

// auto run
func init() {
	system.RegisterInit(initOrderHomeProductCategory, &initHomeProductCategory{})
}

func (i *initHomeProductCategory) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.ProductCategory{},
	)
}

func (i *initHomeProductCategory) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.ProductCategory{})
}

func (i initHomeProductCategory) InitializerName() string {
	return wechatModel.ProductCategory{}.TableName()
}

func (i *initHomeProductCategory) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.ProductCategory{
		{
			ParentId:     0,
			Name:         "服装",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "服装",
			Description:  "服装分类",
		},
		{
			ParentId:     0,
			Name:         "休闲玩乐",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "休闲玩乐",
			Description:  "休闲玩乐",
		},
		{
			ParentId:     0,
			Name:         "超时便利",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "超时便利",
			Description:  "超时便利",
		},
		{
			ParentId:     0,
			Name:         "美食",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "美食",
			Description:  "美食",
		},
		{
			ParentId:     1,
			Name:         "外套",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac4780cN6087feb5.jpg",
			Keywords:     "外套",
			Description:  "外套",
		},
		{
			ParentId:     1,
			Name:         "T恤",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47ffaN8a7b2e14.png",
			Keywords:     "T恤",
			Description:  "T恤",
		},
		{
			ParentId:     1,
			Name:         "休闲裤",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47845N7374a31d.jpg",
			Keywords:     "休闲裤",
			Description:  "休闲裤",
		},
		{
			ParentId:     1,
			Name:         "牛仔裤",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47841Nff658599.jpg",
			Keywords:     "牛仔裤",
			Description:  "牛仔裤",
		},
		{
			ParentId:     1,
			Name:         "衬衫",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac48007Nb30b2118.jpg",
			Keywords:     "衬衫",
			Description:  "衬衫",
		},
		{
			ParentId:     1,
			Name:         "男鞋",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   0,
			Sort:         0,
			Icon:         "",
			Keywords:     "服装",
			Description:  "服装分类",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.ProductCategory{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initHomeProductCategory) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "笔记本").First(&wechatModel.ProductCategory{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
