
// 自动生成模板LockerOrders
package ito_user
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// lockerOrders表 结构体  LockerOrders
type LockerOrders struct {
    global.GVA_MODEL
  OrderNumber  *string `json:"orderNumber" form:"orderNumber" gorm:"comment:业务订单号（唯一标识）;column:order_number;size:200;" binding:"required"`  //业务订单号（唯一标识）
  UserId  *int `json:"userId" form:"userId" gorm:"comment:用户ID（关联用户表）;column:user_id;size:20;" binding:"required"`  //用户ID（关联用户表）
  StartTime  *time.Time `json:"startTime" form:"startTime" gorm:"comment:寄存开始时间;column:start_time;"`  //寄存开始时间
  ScheduledDuration  *int `json:"scheduledDuration" form:"scheduledDuration" gorm:"comment:计划寄存时长（小时）;column:scheduled_duration;size:10;" binding:"required"`  //计划寄存时长（小时）
  ActualDuration  *int `json:"actualDuration" form:"actualDuration" gorm:"comment:实际寄存时长（小时）;column:actual_duration;size:10;"`  //实际寄存时长（小时）
  Price  *float64 `json:"price" form:"price" gorm:"comment:基础费用;column:price;size:10;"`  //基础费用
  Discount  *float64 `json:"discount" form:"discount" gorm:"comment:优惠金额;column:discount;size:10;"`  //优惠金额
  AmountPaid  *float64 `json:"amountPaid" form:"amountPaid" gorm:"comment:实付金额;column:amount_paid;size:10;" binding:"required"`  //实付金额
  StorageLocationName  *string `json:"storageLocationName" form:"storageLocationName" gorm:"comment:寄存网点名称;column:storage_location_name;size:40;" binding:"required"`  //寄存网点名称
  CabinetId  *int `json:"cabinetId" form:"cabinetId" gorm:"comment:柜子ID;column:cabinet_id;size:10;" binding:"required"`  //柜子ID
  Status  *string `json:"status" form:"status" gorm:"comment:订单状态：1-待支付、2-寄存中、3-已完成、4-已取消、5-超时、6-异常;column:status;" binding:"required"`  //订单状态：1-待支付、2-寄存中、3-已完成、4-已取消、5-超时、6-异常
  CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"comment:创建时间;column:create_time;" binding:"required"`  //创建时间
  UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"comment:更新时间;column:update_time;" binding:"required"`  //更新时间
  DepositStatus  *string `json:"depositStatus" form:"depositStatus" gorm:"comment:押金状态：1-已支付、2-已退还、3-已扣除;column:deposit_status;" binding:"required"`  //押金状态：1-已支付、2-已退还、3-已扣除
  LockerPointId  *int `json:"lockerPointId" form:"lockerPointId" gorm:"comment:寄存点id;column:locker_point_id;size:10;"`  //寄存点id
  Title  *string `json:"title" form:"title" gorm:"column:title;size:50;" binding:"required"`  //title字段
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName lockerOrders表 LockerOrders自定义表名 locker_orders
func (LockerOrders) TableName() string {
    return "locker_orders"
}





