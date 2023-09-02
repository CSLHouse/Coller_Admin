package business

import (
	"context"
	businessModel "github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderCombo = system.InitBusinessSystem + 1

type initCombo struct{}

// auto run
func init() {
	//system.RegisterInit(initOrderCombo, &initCombo{})
}

func (i *initCombo) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&businessModel.VIPCombo{},
	)
}

func (i *initCombo) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&businessModel.VIPCombo{})
}

func (i initCombo) InitializerName() string {
	return businessModel.VIPCombo{}.TableName()
}

func (i *initCombo) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []businessModel.VIPCombo{
		{
			ComboName:  "10次卡",
			ComboPrice: 198,
			Times:      1,
			State:      1,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, businessModel.VIPCombo{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initCombo) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("ComboName = ?", "10次卡").First(&businessModel.VIPCombo{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
