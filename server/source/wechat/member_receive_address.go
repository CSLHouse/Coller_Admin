package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderReceiveAddress = initOrderWXUser + 1

type initMemberReceiveAddress struct{}

// auto run
func init() {
	system.RegisterInit(initOrderReceiveAddress, &initMemberReceiveAddress{})
}

func (i *initMemberReceiveAddress) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.MemberReceiveAddress{},
	)
}

func (i *initMemberReceiveAddress) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.MemberReceiveAddress{})
}

func (i initMemberReceiveAddress) InitializerName() string {
	return wechatModel.MemberReceiveAddress{}.TableName()
}

func (i *initMemberReceiveAddress) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.MemberReceiveAddress{
		{

			UserId:        1,
			Name:          "小李",
			PhoneNumber:   "15101668083",
			DefaultStatus: 0,
			PostCode:      "",
			Province:      "河南省",
			City:          "商丘市",
			Region:        "梁园区",
			DetailAddress: "雍景台2号楼4层401",
		},
		{

			UserId:        1,
			Name:          "大个",
			PhoneNumber:   "17601627456",
			DefaultStatus: 1,
			PostCode:      "",
			Province:      "河南省",
			City:          "商丘市",
			Region:        "梁园区",
			DetailAddress: "雍景台商业街西段27号楼113商铺",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.MemberReceiveAddress{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, err
}

func (i *initMemberReceiveAddress) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("user_id = ?", 45).First(&wechatModel.MemberReceiveAddress{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
