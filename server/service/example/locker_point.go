package example

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
	"time"
)

type LockerPointService struct{}

// CreateLockerPoint 创建网点管理记录
// Author [yourname](https://github.com/yourname)
func (lockerPointService *LockerPointService) CreateLockerPoint(ctx context.Context, lockerPoint *example.LockerPoint) (err error) {
	err = global.GVA_ITO_DB.Create(lockerPoint).Error
	return err
}

// DeleteLockerPoint 删除网点管理记录
// Author [yourname](https://github.com/yourname)
func (lockerPointService *LockerPointService) DeleteLockerPoint(ctx context.Context, id string) (err error) {
	err = global.GVA_ITO_DB.Delete(&example.LockerPoint{}, "id = ?", id).Error
	return err
}

// DeleteLockerPointByIds 批量删除网点管理记录
// Author [yourname](https://github.com/yourname)
func (lockerPointService *LockerPointService) DeleteLockerPointByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_ITO_DB.Delete(&[]example.LockerPoint{}, "id in ?", ids).Error
	return err
}

// UpdateLockerPoint 更新网点管理记录
// Author [yourname](https://github.com/yourname)
func (lockerPointService *LockerPointService) UpdateLockerPoint(ctx context.Context, lockerPoint example.LockerPoint) (err error) {
	err = global.GVA_ITO_DB.Model(&example.LockerPoint{}).Where("id = ?", lockerPoint.Id).Updates(&lockerPoint).Error
	return err
}

// GetLockerPoint 根据id获取网点管理记录
// Author [yourname](https://github.com/yourname)
func (lockerPointService *LockerPointService) GetLockerPoint(ctx context.Context, id string) (lockerPoint example.LockerPoint, err error) {
	err = global.GVA_ITO_DB.Where("id = ?", id).First(&lockerPoint).Error
	return
}

// GetLockerPointInfoList 分页获取网点管理记录
// Author [yourname](https://github.com/yourname)
func (lockerPointService *LockerPointService) GetLockerPointInfoList(ctx context.Context, info exampleReq.LockerPointSearch) (list []example.LockerPoint, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_ITO_DB.Model(&example.LockerPoint{})
	var lockerPoints []example.LockerPoint
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&lockerPoints).Error
	var ids []int32
	for _, lockerPoint := range lockerPoints {
		ids = append(ids, int32(*lockerPoint.Id))
	}
	type MonthPrice struct {
		LockerPointId int32   `gorm:"column:locker_point_id"`
		SumNum        float64 `gorm:"column:sum_num"`
	}
	var monthPrice []MonthPrice
	err = global.GVA_ITO_DB.Table("locker_orders").
		Where("locker_point_id in (?)", ids).
		Select("locker_point_id,sum(amount_paid) as sum_num").
		Where("status = ?", 3).Debug().
		Where("start_time > ?", time.Now().Format("2006-01")+"-01 00:00:00").
		Group("locker_point_id").Find(&monthPrice).Error

	if err != nil {
		fmt.Println(err)
	}
	monthPriceMap := make(map[int32]float64, len(monthPrice))
	for _, m := range monthPrice {
		monthPriceMap[m.LockerPointId] = m.SumNum
	}

	var dayPrice []MonthPrice
	err = global.GVA_ITO_DB.Table("locker_orders").
		Where("locker_point_id in (?)", ids).
		Select("locker_point_id,sum(amount_paid) as sum_num").
		Where("status = ?", 3).Debug().
		Where("start_time > ?", time.Now().Format("2006-01-02")+"00:00:00").
		Group("locker_point_id").Find(&dayPrice).Error

	if err != nil {
		fmt.Println(err)
	}
	dayPriceMap := make(map[int32]float64, len(dayPrice))
	for _, m := range dayPrice {
		dayPriceMap[m.LockerPointId] = m.SumNum
	}

	for i, m := range lockerPoints {
		lockerPoints[i].MonthSum = monthPriceMap[int32(*m.Id)]
		lockerPoints[i].DaySum = dayPriceMap[int32(*m.Id)]
	}

	return lockerPoints, total, err
}
func (lockerPointService *LockerPointService) GetLockerPointPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
