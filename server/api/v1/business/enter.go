package business

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ComboApi
	MemberApi
	ConsumeApi
	OrderApi
	QrCodeApi
}

var (
	comboService                 = service.ServiceGroupApp.BusinessServiceGroup.VIPComboService
	memberService                = service.ServiceGroupApp.BusinessServiceGroup.VIPMemberService
	consumeService               = service.ServiceGroupApp.BusinessServiceGroup.VIPConsumeService
	orderService                 = service.ServiceGroupApp.BusinessServiceGroup.VIPOrderService
	qrcodeService                = service.ServiceGroupApp.BusinessServiceGroup.QrCodeService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
