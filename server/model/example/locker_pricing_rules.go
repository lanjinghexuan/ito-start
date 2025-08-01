
// 自动生成模板LockerPricingRules
package example
import (
	"time"
)

// lockerPricingRules表 结构体  LockerPricingRules
type LockerPricingRules struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:19;"`  //id字段
  NetworkId  *int `json:"networkId" form:"networkId" gorm:"comment:网点ID;column:network_id;size:19;"`  //网点ID
  RuleName  *string `json:"ruleName" form:"ruleName" gorm:"comment:规则名称;column:rule_name;size:50;"`  //规则名称
  FeeType  *bool `json:"feeType" form:"feeType" gorm:"comment:1-计时收费 2-按日收费;column:fee_type;"`  //1-计时收费 2-按日收费
  LockerType  *bool `json:"lockerType" form:"lockerType" gorm:"comment:1-小柜子 2-大柜子;column:locker_type;"`  //1-小柜子 2-大柜子
  FreeDuration  *float64 `json:"freeDuration" form:"freeDuration" gorm:"comment:免费时长(小时);column:free_duration;size:5;"`  //免费时长(小时)
  IsDepositEnabled  *bool `json:"isDepositEnabled" form:"isDepositEnabled" gorm:"comment:是否启用押金;column:is_deposit_enabled;"`  //是否启用押金
  IsAdvancePay  *bool `json:"isAdvancePay" form:"isAdvancePay" gorm:"comment:是否启用预付;column:is_advance_pay;"`  //是否启用预付
  HourlyRate  *float64 `json:"hourlyRate" form:"hourlyRate" gorm:"comment:每小时费用;column:hourly_rate;size:10;"`  //每小时费用
  DailyCap  *float64 `json:"dailyCap" form:"dailyCap" gorm:"comment:24小时封顶价;column:daily_cap;size:10;"`  //24小时封顶价
  DailyRate  *float64 `json:"dailyRate" form:"dailyRate" gorm:"comment:每日费用;column:daily_rate;size:10;"`  //每日费用
  AdvanceAmount  *float64 `json:"advanceAmount" form:"advanceAmount" gorm:"comment:预付金额;column:advance_amount;size:10;"`  //预付金额
  DepositAmount  *float64 `json:"depositAmount" form:"depositAmount" gorm:"comment:押金金额;column:deposit_amount;size:10;"`  //押金金额
  Status  *bool `json:"status" form:"status" gorm:"comment:1-生效 0-失效;column:status;"`  //1-生效 0-失效
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
}


// TableName lockerPricingRules表 LockerPricingRules自定义表名 locker_pricing_rules
func (LockerPricingRules) TableName() string {
    return "locker_pricing_rules"
}





