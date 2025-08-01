
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type LockerPricingRulesSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      NetworkId  *int `json:"networkId" form:"networkId"` 
      RuleName  *string `json:"ruleName" form:"ruleName"` 
      FeeType  *string `json:"feeType" form:"feeType"` 
      LockerType  *string `json:"lockerType" form:"lockerType"` 
      FreeDuration  *float64 `json:"freeDuration" form:"freeDuration"` 
      IsDepositEnabled  *string `json:"isDepositEnabled" form:"isDepositEnabled"` 
      IsAdvancePay  *string `json:"isAdvancePay" form:"isAdvancePay"` 
      HourlyRate  *float64 `json:"hourlyRate" form:"hourlyRate"` 
      DailyCap  *float64 `json:"dailyCap" form:"dailyCap"` 
      DailyRate  *float64 `json:"dailyRate" form:"dailyRate"` 
      AdvanceAmount  *float64 `json:"advanceAmount" form:"advanceAmount"` 
      DepositAmount  *float64 `json:"depositAmount" form:"depositAmount"` 
      Status  *string `json:"status" form:"status"` 
    request.PageInfo
}
