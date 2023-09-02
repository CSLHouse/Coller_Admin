package business

import (
	"context"
	businessModel "github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMember = system.InitBusinessSystem + 1

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
		&businessModel.VIPMember{},
		&businessModel.VIPCombo{},
		&businessModel.VIPOrder{},
		&businessModel.VIPStatement{},
		&businessModel.VIPStatistics{},
		&businessModel.ConsumeRecord{},
	)
}

func (i *initMember) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&businessModel.VIPMember{})
}

func (i initMember) InitializerName() string {
	return businessModel.VIPMember{}.TableName()
}

func (i *initMember) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []businessModel.VIPMember{
		{
			CardID:      123,
			MemberName:  "艾米",
			ComboId:     0,
			RemainTimes: 1,
			Deadline:    "",
			State:       1,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, businessModel.VIPMember{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initMember) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("ComboName = ?", "10次卡").First(&businessModel.VIPMember{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
