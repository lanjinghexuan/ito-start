package example

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	CustomerRouter
	FileUploadAndDownloadRouter
	AttachmentCategoryRouter
	LockerPricingRulesRouter
	LockerPointRouter
}

var (
	exaCustomerApi              = api.ApiGroupApp.ExampleApiGroup.CustomerApi
	exaFileUploadAndDownloadApi = api.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	attachmentCategoryApi       = api.ApiGroupApp.ExampleApiGroup.AttachmentCategoryApi
	lockerPricingRulesApi       = api.ApiGroupApp.ExampleApiGroup.LockerPricingRulesApi
	lockerPointApi              = api.ApiGroupApp.ExampleApiGroup.LockerPointApi
)
