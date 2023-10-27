package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderBrand = initOrderHome + 1

type initBrand struct{}

// auto run
func init() {
	system.RegisterInit(initOrderBrand, &initBrand{})
}

func (i *initBrand) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.HomeBrand{},
	)
}

func (i *initBrand) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.HomeBrand{})
}

func (i initBrand) InitializerName() string {
	return wechatModel.HomeBrand{}.TableName()
}

func (i *initBrand) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.HomeBrand{
		{
			Name:                "万和",
			FirstLetter:         "W",
			Sort:                0,
			FactoryStatus:       1,
			ShowStatus:          1,
			ProductCount:        100,
			ProductCommentCount: 100,
			Logo:                "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5b07ca8aN4e127d2f.jpg",
			BigPic:              "http://img13.360buyimg.com/cms/jfs/t1/121860/35/2430/187800/5ec4e294E22f3ffcc/1e233b65b94ba192.jpg",
			BrandStory:          "万和成立于1993年8月，总部位于广东顺德国家级高新技术开发区内，是国内生产规模最大的燃气具专业制造企业，也是中国燃气具发展战略的首倡者和推动者、中国五金制品协会燃气用具分会第三届理事长单位。",
		},
		{
			Name:                "三星",
			FirstLetter:         "S",
			Sort:                100,
			FactoryStatus:       1,
			ShowStatus:          1,
			ProductCount:        100,
			ProductCommentCount: 100,
			Logo:                "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/57201b47N7bf15715.jpg",
			BigPic:              "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/sanxing_banner_01.png",
			BrandStory:          "三星集团（英文：SAMSUNG、韩文：삼성）是韩国最大的跨国企业集团，三星集团包括众多的国际下属企业，旗下子公司有：三星电子、三星物产、三星人寿保险等，业务涉及电子、金融、机械、化学等众多领域。",
		},
		{
			Name:                "华为",
			FirstLetter:         "H",
			Sort:                100,
			FactoryStatus:       1,
			ShowStatus:          1,
			ProductCount:        100,
			ProductCommentCount: 100,
			Logo:                "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5abf6f26N31658aa2.jpg",
			BigPic:              "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png",
			BrandStory:          "荣耀品牌成立于2013年,是华为旗下手机双品牌之一。荣耀以“创新、品质、服务”为核心战略,为全球年轻人提供潮酷的全场景智能化体验,打造年轻人向往的先锋文化和潮流生活方式",
		},
		{
			Name:                "格力",
			FirstLetter:         "G",
			Sort:                30,
			FactoryStatus:       1,
			ShowStatus:          1,
			ProductCount:        100,
			ProductCommentCount: 100,
			Logo:                "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg (3).jpg",
			BigPic:              "",
			BrandStory:          "Victoria\\'s Secret的故事",
		},
		{
			Name:                "小米",
			FirstLetter:         "M",
			Sort:                500,
			FactoryStatus:       1,
			ShowStatus:          1,
			ProductCount:        100,
			ProductCommentCount: 100,
			Logo:                "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5565f5a2N0b8169ae.jpg",
			BigPic:              "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/xiaomi_banner_01.png",
			BrandStory:          "小米公司正式成立于2010年4月，是一家专注于高端智能手机、互联网电视自主研发的创新型科技企业。主要由前谷歌、微软、摩托、金山等知名公司的顶尖人才组建。",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.HomeBrand{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initBrand) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "三星").First(&wechatModel.HomeBrand{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
