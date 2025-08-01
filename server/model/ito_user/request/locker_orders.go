
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type LockerOrdersSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      OrderNumber  *string `json:"orderNumber" form:"orderNumber"` 
      UserId  *int `json:"userId" form:"userId"` 
      StartTime  *time.Time `json:"startTime" form:"startTime"` 
      ScheduledDuration  *int `json:"scheduledDuration" form:"scheduledDuration"` 
      ActualDuration  *int `json:"actualDuration" form:"actualDuration"` 
      Price  *float64 `json:"price" form:"price"` 
      Discount  *float64 `json:"discount" form:"discount"` 
      AmountPaid  *float64 `json:"amountPaid" form:"amountPaid"` 
      StorageLocationName  *string `json:"storageLocationName" form:"storageLocationName"` 
      CabinetId  *int `json:"cabinetId" form:"cabinetId"` 
      Status  *string `json:"status" form:"status"` 
      CreateTime  *time.Time `json:"createTime" form:"createTime"` 
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime"` 
      DepositStatus  *string `json:"depositStatus" form:"depositStatus"` 
      LockerPointId  *int `json:"lockerPointId" form:"lockerPointId"` 
      Title  *string `json:"title" form:"title"` 
    request.PageInfo
}
