package business

import (
	"context"
	businessModel "github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
)

const initOrderConsume = initOrderMember + 1

type initConsume struct{}

// auto run
func init() {
	//system.RegisterInit(initOrderConsume, &initConsume{})
}

func (i *initConsume) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&businessModel.ConsumeRecord{},
	)
}

func (i *initConsume) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&businessModel.ConsumeRecord{})
}

func (i initConsume) InitializerName() string {
	return businessModel.ConsumeRecord{}.TableName()
}

func (i *initConsume) InitializeData(ctx context.Context) (next context.Context, err error) {
	return next, err
}

func (i *initConsume) DataInserted(ctx context.Context) bool {
	return true
}
