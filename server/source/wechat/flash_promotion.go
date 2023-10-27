package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderFlashPromotion = initOrderBrand + 1

type initFlashPromotion struct{}

// auto run
func init() {
	system.RegisterInit(initOrderFlashPromotion, &initFlashPromotion{})
}

func (i *initFlashPromotion) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.HomeFlashPromotion{},
	)
}

func (i *initFlashPromotion) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.HomeFlashPromotion{})
}

func (i initFlashPromotion) InitializerName() string {
	return wechatModel.HomeFlashPromotion{}.TableName()
}

func (i *initFlashPromotion) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.HomeFlashPromotion{
		{
			Title:     "双11特卖活动",
			StartDate: "2022-11-09",
			EndDate:   "2023-12-31",
			Status:    1,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.HomeFlashPromotion{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initFlashPromotion) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("title = ?", "双11特卖活动").First(&wechatModel.HomeFlashPromotion{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
