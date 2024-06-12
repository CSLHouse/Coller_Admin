package business

import (
	"context"
	businessModel "cooller/server/model/business"
	"cooller/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMember = initOrderConsume + 1

type initMember struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMember, &initMember{})
}

func (i *initMember) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&businessModel.Customer{},
		&businessModel.VIPCard{},
		&businessModel.VIPOrder{},
		&businessModel.VIPStatement{},
		&businessModel.VIPStatistics{},
		&businessModel.ConsumeRecord{},
		&businessModel.VIPCertificate{},
		&businessModel.QrCode{},
	)
}

func (i *initMember) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&businessModel.Customer{})
}

func (i initMember) InitializerName() string {
	return businessModel.Customer{}.TableName()
}

func (i *initMember) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []businessModel.QrCode{
		{
			Title: "猪迪克",
			Url:   "https://weixin.qq.com/g/AwYAANUTOI753GzyQ1_P238ACPiea5tuP8F-GQCOclpDCg-GvpQxLppZk0_4Xiwk",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, businessModel.QrCode{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initMember) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("title = ?", "猪迪克").First(&businessModel.QrCode{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
