package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderHome = system.InitWechatInternal + 1

type initHome struct{}

// auto run
func init() {
	system.RegisterInit(initOrderHome, &initHome{})
}

func (i *initHome) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.HomeAdvertise{},
	)
}

func (i *initHome) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.HomeAdvertise{})
}

func (i initHome) InitializerName() string {
	return wechatModel.HomeAdvertise{}.TableName()
}

func (i *initHome) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.HomeAdvertise{
		{
			Name:       "小米推荐广告",
			Type:       1,
			Pic:        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/xiaomi_banner_01.png",
			StartTime:  "2022-11-08 17:04:03",
			EndTime:    "2023-11-08 17:04:05",
			State:      1,
			ClickCount: 0,
			OrderCount: 0,
			Url:        "/pages/brand/brandDetail?id=6",
			Note:       "夏季大热促销",
			Sort:       0,
		},
		{
			Name:       "华为推荐广告",
			Type:       1,
			Pic:        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png",
			StartTime:  "2022-11-08 17:12:54",
			EndTime:    "2023-11-08 17:12:55",
			State:      1,
			ClickCount: 0,
			OrderCount: 0,
			Url:        "/pages/brand/brandDetail?id=51",
			Note:       "",
			Sort:       0,
		},
		{
			Name:       "电影推荐广告",
			Type:       1,
			Pic:        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20181113/movie_ad.jpg",
			StartTime:  "2018-11-01 00:00:00",
			EndTime:    "2018-11-24 00:00:00",
			State:      0,
			ClickCount: 0,
			OrderCount: 0,
			Url:        "/pages/brand/brandDetail?id=51",
			Note:       "www.baidu.com', '电影推荐广告",
			Sort:       100,
		},
		{
			Name:       "汽车促销广告",
			Type:       1,
			Pic:        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20181113/car_ad.jpg",
			StartTime:  "2018-11-13 00:00:00",
			EndTime:    "2018-11-24 00:00:00",
			State:      0,
			ClickCount: 0,
			OrderCount: 0,
			Url:        "xxx",
			Note:       "www.baidu.com', '电影推荐广告",
			Sort:       99,
		},
		{
			Name:       "夏季大热促销",
			Type:       1,
			Pic:        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190525/ad1.jpg",
			StartTime:  "2018-11-01 14:01:37",
			EndTime:    "2018-11-15 14:01:37",
			State:      0,
			ClickCount: 0,
			OrderCount: 0,
			Url:        "",
			Note:       "夏季大热促销",
			Sort:       0,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.HomeAdvertise{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initHome) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "夏季大热促销").First(&wechatModel.HomeAdvertise{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
