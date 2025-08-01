// 自动生成模板LockerPoint
package example

// 网点管理 结构体  LockerPoint
type LockerPoint struct {
	Id              *int     `json:"id" form:"id" gorm:"primarykey;comment:寄存点ID;column:id;size:10;"`                               //寄存点ID
	LocationId      *int     `json:"locationId" form:"locationId" gorm:"comment:所属地点ID;column:location_id;size:10;"`                //所属地点ID
	Name            *string  `json:"name" form:"name" gorm:"comment:寄存点名称;column:name;size:30;"`                                    //寄存点名称
	Address         *string  `json:"address" form:"address" gorm:"comment:详细地址;column:address;size:50;"`                            //详细地址
	Latitude        *float64 `json:"latitude" form:"latitude" gorm:"comment:纬度;column:latitude;size:11;"`                           //纬度
	Longitude       *float64 `json:"longitude" form:"longitude" gorm:"comment:经度;column:longitude;size:11;"`                        //经度
	AvailableLarge  *int     `json:"availableLarge" form:"availableLarge" gorm:"comment:可用大柜数量;column:available_large;size:10;"`    //可用大柜数量
	AvailableMedium *int     `json:"availableMedium" form:"availableMedium" gorm:"comment:可用中柜数量;column:available_medium;size:10;"` //可用中柜数量
	AvailableSmall  *int     `json:"availableSmall" form:"availableSmall" gorm:"comment:可用小柜数量;column:available_small;size:10;"`    //可用小柜数量
	OpenTime        *string  `json:"openTime" form:"openTime" gorm:"comment:营业时间;column:open_time;size:30;"`                        //营业时间
	Mobile          *string  `json:"mobile" form:"mobile" gorm:"comment:联系电话;column:mobile;size:20;"`                               //联系电话
	AdminId         *int     `json:"adminId" form:"adminId" gorm:"comment:管理员id;column:admin_id;size:10;"`                          //管理员id
	PointImage      string   `json:"pointImage" form:"pointImage" gorm:"comment:网点图片;column:point_image;size:200;"`                 //网点图片
	Status          *string  `json:"status" form:"status" gorm:"comment:状态;column:status;size:30;"`                                 //状态
	PointType       *string  `json:"pointType" form:"pointType" gorm:"comment:网点类型;column:point_type;size:20;"`                     //网点类型

	MonthSum float64 `gorm:"-" json:"monthSum"` // gorm:"-" 表示不映射数据库字段
	DaySum   float64 `gorm:"-" json:"daySum"`   // gorm:"-" 表示不映射数据库字段
}

// TableName 网点管理 LockerPoint自定义表名 locker_point
func (LockerPoint) TableName() string {
	return "locker_point"
}
