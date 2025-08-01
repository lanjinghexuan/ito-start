package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
	AttachmentCategoryApi
	LockerPricingRulesApi
	LockerPointApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	attachmentCategoryService    = service.ServiceGroupApp.ExampleServiceGroup.AttachmentCategoryService
	lockerPricingRulesService    = service.ServiceGroupApp.ExampleServiceGroup.LockerPricingRulesService
	lockerPointService           = service.ServiceGroupApp.ExampleServiceGroup.LockerPointService
)
