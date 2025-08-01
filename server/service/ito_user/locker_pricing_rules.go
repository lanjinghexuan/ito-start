
package ito_user

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ito_user"
    ito_userReq "github.com/flipped-aurora/gin-vue-admin/server/model/ito_user/request"
    "gorm.io/gorm"
)

type LockerPricingRulesService struct {}
// CreateLockerPricingRules 创建lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService) CreateLockerPricingRules(ctx context.Context, lockerPricingRules *ito_user.LockerPricingRules) (err error) {
	err = global.GVA_DB.Create(lockerPricingRules).Error
	return err
}

// DeleteLockerPricingRules 删除lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)DeleteLockerPricingRules(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ito_user.LockerPricingRules{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&ito_user.LockerPricingRules{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteLockerPricingRulesByIds 批量删除lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)DeleteLockerPricingRulesByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ito_user.LockerPricingRules{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&ito_user.LockerPricingRules{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateLockerPricingRules 更新lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)UpdateLockerPricingRules(ctx context.Context, lockerPricingRules ito_user.LockerPricingRules) (err error) {
	err = global.GVA_DB.Model(&ito_user.LockerPricingRules{}).Where("id = ?",lockerPricingRules.ID).Updates(&lockerPricingRules).Error
	return err
}

// GetLockerPricingRules 根据ID获取lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)GetLockerPricingRules(ctx context.Context, ID string) (lockerPricingRules ito_user.LockerPricingRules, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&lockerPricingRules).Error
	return
}
// GetLockerPricingRulesInfoList 分页获取lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)GetLockerPricingRulesInfoList(ctx context.Context, info ito_userReq.LockerPricingRulesSearch) (list []ito_user.LockerPricingRules, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ito_user.LockerPricingRules{})
    var lockerPricingRuless []ito_user.LockerPricingRules
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.NetworkId != nil {
        db = db.Where("network_id = ?", *info.NetworkId)
    }
    if info.RuleName != nil && *info.RuleName != "" {
        db = db.Where("rule_name = ?", *info.RuleName)
    }
    if info.FeeType != nil && *info.FeeType != "" {
        db = db.Where("fee_type LIKE ?", "%"+ *info.FeeType+"%")
    }
    if info.LockerType != nil && *info.LockerType != "" {
        db = db.Where("locker_type LIKE ?", "%"+ *info.LockerType+"%")
    }
    if info.FreeDuration != nil {
        db = db.Where("free_duration = ?", *info.FreeDuration)
    }
    if info.IsDepositEnabled != nil && *info.IsDepositEnabled != "" {
        db = db.Where("is_deposit_enabled LIKE ?", "%"+ *info.IsDepositEnabled+"%")
    }
    if info.IsAdvancePay != nil && *info.IsAdvancePay != "" {
        db = db.Where("is_advance_pay LIKE ?", "%"+ *info.IsAdvancePay+"%")
    }
    if info.HourlyRate != nil {
        db = db.Where("hourly_rate = ?", *info.HourlyRate)
    }
    if info.DailyCap != nil {
        db = db.Where("daily_cap = ?", *info.DailyCap)
    }
    if info.DailyRate != nil {
        db = db.Where("daily_rate = ?", *info.DailyRate)
    }
    if info.AdvanceAmount != nil {
        db = db.Where("advance_amount = ?", *info.AdvanceAmount)
    }
    if info.DepositAmount != nil {
        db = db.Where("deposit_amount = ?", *info.DepositAmount)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status LIKE ?", "%"+ *info.Status+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&lockerPricingRuless).Error
	return  lockerPricingRuless, total, err
}
func (lockerPricingRulesService *LockerPricingRulesService)GetLockerPricingRulesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
