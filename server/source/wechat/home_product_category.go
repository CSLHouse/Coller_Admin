package wechat

import (
	"context"
	wechatModel "cooller/server/model/wechat"
	"cooller/server/service/system"
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
			Name:         "门票",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "张",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "门票",
			Description:  "门票",
		},
		{
			ParentId:     0,
			Name:         "玩具",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "玩具",
			Description:  "玩具",
		},
		{
			ParentId:     0,
			Name:         "饮料",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "瓶",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "饮料",
			Description:  "饮料",
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
	if errors.Is(db.Where("name = ?", "饮料").First(&wechatModel.ProductCategory{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
