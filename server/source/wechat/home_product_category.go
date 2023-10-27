package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
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
		&wechatModel.HomeProductCategory{},
	)
}

func (i *initHomeProductCategory) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.HomeProductCategory{})
}

func (i initHomeProductCategory) InitializerName() string {
	return wechatModel.HomeProductCategory{}.TableName()
}

func (i *initHomeProductCategory) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.HomeProductCategory{
		{
			ParentId:     0,
			Name:         "服装",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "服装",
			Description:  "服装分类",
		},
		{
			ParentId:     0,
			Name:         "手机数码",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "手机数码",
			Description:  "手机数码",
		},
		{
			ParentId:     0,
			Name:         "家用电器",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "家用电器",
			Description:  "家用电器",
		},
		{
			ParentId:     0,
			Name:         "家具家装",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "家具家装",
			Description:  "家具家装",
		},
		{
			ParentId:     0,
			Name:         "汽车用品",
			Level:        0,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "汽车用品",
			Description:  "汽车用品",
		},
		{
			ParentId:     1,
			Name:         "外套",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac4780cN6087feb5.jpg",
			Keywords:     "外套",
			Description:  "外套",
		},
		{
			ParentId:     1,
			Name:         "T恤",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47ffaN8a7b2e14.png",
			Keywords:     "T恤",
			Description:  "T恤",
		},
		{
			ParentId:     1,
			Name:         "休闲裤",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47845N7374a31d.jpg",
			Keywords:     "休闲裤",
			Description:  "休闲裤",
		},
		{
			ParentId:     1,
			Name:         "牛仔裤",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac47841Nff658599.jpg",
			Keywords:     "牛仔裤",
			Description:  "牛仔裤",
		},
		{
			ParentId:     1,
			Name:         "衬衫",
			Level:        1,
			ProductCount: 100,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac48007Nb30b2118.jpg",
			Keywords:     "衬衫",
			Description:  "衬衫",
		},
		{
			ParentId:     2,
			Name:         "手机通讯",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac48d27N3f5bb821.jpg",
			Keywords:     "手机通讯",
			Description:  "手机通讯",
		},
		{
			ParentId:     1,
			Name:         "男鞋",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   0,
			Sort:         0,
			Icon:         "",
			Keywords:     "服装",
			Description:  "服装分类",
		},
		{
			ParentId:     2,
			Name:         "手机配件",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5ac48672N11cf61fe.jpg",
			Keywords:     "手机配件",
			Description:  "手机配件",
		},
		{
			ParentId:     2,
			Name:         "摄影摄像",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5a1679f2Nc2f659b6.jpg",
			Keywords:     "摄影摄像",
			Description:  "摄影摄像",
		},
		{
			ParentId:     2,
			Name:         "影音娱乐",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5a167859N01d8198b.jpg",
			Keywords:     "影音娱乐",
			Description:  "影音娱乐",
		},
		{
			ParentId:     2,
			Name:         "数码配件",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20190519/5a1676e9N1ba70a81.jpg",
			Keywords:     "数码配件",
			Description:  "数码配件",
		},
		{
			ParentId:     2,
			Name:         "智能设备",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "",
			Keywords:     "",
			Description:  "",
		},
		{
			ParentId:     3,
			Name:         "电视",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5a17f71eN25360979.jpg",
			Keywords:     "电视",
			Description:  "电视",
		},
		{
			ParentId:     3,
			Name:         "空调",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5a17f6f6Ndfe746aa.jpg",
			Keywords:     "空调",
			Description:  "空调",
		},
		{
			ParentId:     3,
			Name:         "洗衣机",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5a17f6eaN9ec936de.jpg",
			Keywords:     "洗衣机",
			Description:  "洗衣机",
		},
		{
			ParentId:     3,
			Name:         "冰箱",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5a17f6c5Ne56d7e26.jpg",
			Keywords:     "冰箱",
			Description:  "冰箱",
		},
		{
			ParentId:     4,
			Name:         "五金工具",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5a28e743Nf6d99998.jpg",
			Keywords:     "五金工具",
			Description:  "五金工具",
		},
		{
			ParentId:     0,
			Name:         "电脑办公",
			Level:        0,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    1,
			ShowStatus:   1,
			Sort:         1,
			Icon:         "",
			Keywords:     "电脑办公",
			Description:  "电脑办公",
		},
		{
			ParentId:     52,
			Name:         "平板电脑",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/pad_category_01.jpg",
			Keywords:     "平板电脑",
			Description:  "平板电脑",
		},
		{
			ParentId:     52,
			Name:         "笔记本",
			Level:        1,
			ProductCount: 0,
			ProductUnit:  "件",
			NavStatus:    0,
			ShowStatus:   1,
			Sort:         0,
			Icon:         "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/computer_category_01.jpg",
			Keywords:     "笔记本",
			Description:  "笔记本",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.HomeProductCategory{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initHomeProductCategory) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "笔记本").First(&wechatModel.HomeProductCategory{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
