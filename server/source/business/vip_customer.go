package business

import (
	"context"
	businessModel "github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
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
		&businessModel.Customer{},
		&businessModel.VIPCard{},
		&businessModel.VIPCombo{},
		&businessModel.VIPOrder{},
		&businessModel.VIPStatement{},
		&businessModel.VIPStatistics{},
		&businessModel.ConsumeRecord{},
		&businessModel.VIPCertificate{},
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

	return next, err
}

func (i *initMember) DataInserted(ctx context.Context) bool {

	return true
}
