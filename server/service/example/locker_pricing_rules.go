
package example

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
)

type LockerPricingRulesService struct {}
// CreateLockerPricingRules 创建lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService) CreateLockerPricingRules(ctx context.Context, lockerPricingRules *example.LockerPricingRules) (err error) {
	err = global.GVA_DB.Create(lockerPricingRules).Error
	return err
}

// DeleteLockerPricingRules 删除lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)DeleteLockerPricingRules(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&example.LockerPricingRules{},"id = ?",id).Error
	return err
}

// DeleteLockerPricingRulesByIds 批量删除lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)DeleteLockerPricingRulesByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]example.LockerPricingRules{},"id in ?",ids).Error
	return err
}

// UpdateLockerPricingRules 更新lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)UpdateLockerPricingRules(ctx context.Context, lockerPricingRules example.LockerPricingRules) (err error) {
	err = global.GVA_DB.Model(&example.LockerPricingRules{}).Where("id = ?",lockerPricingRules.Id).Updates(&lockerPricingRules).Error
	return err
}

// GetLockerPricingRules 根据id获取lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)GetLockerPricingRules(ctx context.Context, id string) (lockerPricingRules example.LockerPricingRules, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lockerPricingRules).Error
	return
}
// GetLockerPricingRulesInfoList 分页获取lockerPricingRules表记录
// Author [yourname](https://github.com/yourname)
func (lockerPricingRulesService *LockerPricingRulesService)GetLockerPricingRulesInfoList(ctx context.Context, info exampleReq.LockerPricingRulesSearch) (list []example.LockerPricingRules, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&example.LockerPricingRules{})
    var lockerPricingRuless []example.LockerPricingRules
    // 如果有条件搜索 下方会自动创建搜索语句
    
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
