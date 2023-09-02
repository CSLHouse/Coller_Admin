package business

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessRes "github.com/flipped-aurora/gin-vue-admin/server/model/business/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	date_conversion "github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ConsumeApi struct{}

func (e *ConsumeApi) GetConsumeList(c *gin.Context) {
	var searchInfo request.ConsumeSearchInfo
	err := c.ShouldBindQuery(&searchInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	consumeList, total, err := consumeService.GetVIPConsumeInfoList(utils.GetUserAuthorityId(c), searchInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	recordData := make([]businessRes.VipConsumeResModel, 0)
	recordList := consumeList.([]business.ConsumeRecord)
	for _, value := range recordList {
		consumeRes := businessRes.VipConsumeResModel{}
		consumeRes.ID = value.ID
		consumeRes.MemberName = value.Member.MemberName
		consumeRes.Telephone = value.Telephone
		consumeRes.MemberType = value.Member.Combo.ComboName
		consumeRes.MemberState = value.Member.State
		consumeRes.Deadline = value.Member.Deadline
		consumeRes.RemainTimes = value.RemainTimes
		consumeRes.ConsumeTimes = value.ConsumeTimes
		consumeRes.PunchDate = value.PunchDate
		recordData = append(recordData, consumeRes)
	}

	response.OkWithDetailed(response.PageResult{
		List:     recordData,
		Total:    total,
		Page:     searchInfo.Page,
		PageSize: searchInfo.PageSize,
	}, "获取成功", c)
}

// 消费登记
func (e *ConsumeApi) ConsumeVIPCard(c *gin.Context) {
	var consume business.Consume
	err := c.ShouldBindQuery(&consume)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(consume, utils.ConsumeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	oldData, err := memberService.GetVIPMember(consume.CardID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if oldData.CardID != consume.CardID {
		global.GVA_LOG.Error("卡号错误!", zap.Error(err))
		response.FailWithMessage("划卡失败", c)
		return
	}

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("划卡失败", c)
		return
	}
	if oldData.RemainTimes < consume.Number {
		global.GVA_LOG.Error(fmt.Sprintf("划卡失败，卡中剩余:%d!", oldData.RemainTimes), zap.Error(err))
		response.FailWithMessage("划卡失败", c)
		return
	}
	if date_conversion.CheckOverTime(oldData.Deadline) {
		global.GVA_LOG.Error("划卡失败,卡已逾期", zap.Error(err))
		response.FailWithMessage("划卡失败", c)
		return
	}
	oldData.RemainTimes -= consume.Number
	err = memberService.UpdateVIPMember(&oldData)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	var consumeRecord business.ConsumeRecord
	consumeRecord.Telephone = oldData.Telephone
	consumeRecord.RemainTimes = oldData.RemainTimes
	consumeRecord.ConsumeTimes = consume.Number
	consumeRecord.PunchDate = date_conversion.BuildTheTimeStr()
	consumeRecord.State = 1 // 默认已确认
	consumeRecord.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	//err = consumeService.CreateVIPConsume(consumeRecord)
	//if err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//	return
	//}
	var statement business.VIPStatement
	statement.Date = date_conversion.BuildTheDayStr()
	statement.CardNumber = 1
	statement.ConsumeNumber = consumeRecord.ConsumeTimes
	statement.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	//err = orderService.CreateVIPStatement(&statement)
	//if err != nil {
	//	global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
	//	response.FailWithMessage("创建订单失败", c)
	//	return
	//}

	var statistics business.VIPStatistics
	statistics.TotalConsumer = 1
	statistics.SysUserAuthorityID = consumeRecord.SysUserAuthorityID
	//err = orderService.CreateVIPStatistics(&statistics)
	//if err != nil {
	//	global.GVA_LOG.Error("统计失败!", zap.Error(err))
	//	response.FailWithMessage("统计失败", c)
	//	return
	//}
	err = consumeService.CreateVIPConsumeSynchronous(&consumeRecord, &statement, &statistics)
	if err != nil {
		global.GVA_LOG.Error("统计失败!", zap.Error(err))
		response.FailWithMessage("统计失败", c)
		return
	}
	response.OkWithMessage("登记成功", c)
}
