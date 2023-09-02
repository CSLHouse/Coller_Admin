package business

import (
	"context"
	"fmt"
	businessModel "github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
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
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []businessModel.ConsumeRecord{
		{
			RemainTimes:  1,
			ConsumeTimes: 1,
			State:        0,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, businessModel.ConsumeRecord{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initConsume) DataInserted(ctx context.Context) bool {
	fmt.Println("---------initConsume------DataInserted---")
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("MemberId = ?", 123).First(&businessModel.ConsumeRecord{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
