package ito_user

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ito_user"
	ito_userReq "github.com/flipped-aurora/gin-vue-admin/server/model/ito_user/request"
	"gorm.io/gorm"
)

type LockerOrdersService struct{}

// CreateLockerOrders 创建lockerOrders表记录
// Author [yourname](https://github.com/yourname)
func (lockerOrdersService *LockerOrdersService) CreateLockerOrders(ctx context.Context, lockerOrders *ito_user.LockerOrders) (err error) {
	err = global.GVA_DB.Create(lockerOrders).Error
	return err
}

// DeleteLockerOrders 删除lockerOrders表记录
// Author [yourname](https://github.com/yourname)
func (lockerOrdersService *LockerOrdersService) DeleteLockerOrders(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ito_user.LockerOrders{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&ito_user.LockerOrders{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteLockerOrdersByIds 批量删除lockerOrders表记录
// Author [yourname](https://github.com/yourname)
func (lockerOrdersService *LockerOrdersService) DeleteLockerOrdersByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&ito_user.LockerOrders{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&ito_user.LockerOrders{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateLockerOrders 更新lockerOrders表记录
// Author [yourname](https://github.com/yourname)
func (lockerOrdersService *LockerOrdersService) UpdateLockerOrders(ctx context.Context, lockerOrders ito_user.LockerOrders) (err error) {
	err = global.GVA_DB.Model(&ito_user.LockerOrders{}).Where("id = ?", lockerOrders.ID).Updates(&lockerOrders).Error
	return err
}

// GetLockerOrders 根据ID获取lockerOrders表记录
// Author [yourname](https://github.com/yourname)
func (lockerOrdersService *LockerOrdersService) GetLockerOrders(ctx context.Context, ID string) (lockerOrders ito_user.LockerOrders, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&lockerOrders).Error
	return
}

// GetLockerOrdersInfoList 分页获取lockerOrders表记录
// Author [yourname](https://github.com/yourname)
func (lockerOrdersService *LockerOrdersService) GetLockerOrdersInfoList(ctx context.Context, info ito_userReq.LockerOrdersSearch) (list []ito_user.LockerOrders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ito_user.LockerOrders{})
	var lockerOrderss []ito_user.LockerOrders
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.OrderNumber != nil && *info.OrderNumber != "" {
		db = db.Where("order_number = ?", *info.OrderNumber)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
	}
	if info.StartTime != nil {
		db = db.Where("start_time = ?", *info.StartTime)
	}
	if info.ScheduledDuration != nil {
		db = db.Where("scheduled_duration = ?", *info.ScheduledDuration)
	}
	if info.ActualDuration != nil {
		db = db.Where("actual_duration = ?", *info.ActualDuration)
	}
	if info.Price != nil {
		db = db.Where("price = ?", *info.Price)
	}
	if info.Discount != nil {
		db = db.Where("discount = ?", *info.Discount)
	}
	if info.AmountPaid != nil {
		db = db.Where("amount_paid = ?", *info.AmountPaid)
	}
	if info.StorageLocationName != nil && *info.StorageLocationName != "" {
		db = db.Where("storage_location_name = ?", *info.StorageLocationName)
	}
	if info.CabinetId != nil {
		db = db.Where("cabinet_id = ?", *info.CabinetId)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status LIKE ?", "%"+*info.Status+"%")
	}
	if info.CreateTime != nil {
		db = db.Where("create_time = ?", *info.CreateTime)
	}
	if info.UpdateTime != nil {
		db = db.Where("update_time = ?", *info.UpdateTime)
	}
	if info.DepositStatus != nil && *info.DepositStatus != "" {
		db = db.Where("deposit_status LIKE ?", "%"+*info.DepositStatus+"%")
	}
	if info.LockerPointId != nil {
		db = db.Where("locker_point_id = ?", *info.LockerPointId)
	}
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title = ?", *info.Title)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&lockerOrderss).Error
	return lockerOrderss, total, err
}
func (lockerOrdersService *LockerOrdersService) GetLockerOrdersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
