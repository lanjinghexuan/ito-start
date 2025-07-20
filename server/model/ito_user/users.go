
// 自动生成模板Users
package ito_user
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// users表 结构体  Users
type Users struct {
    global.GVA_MODEL
  Username  *string `json:"username" form:"username" gorm:"comment:用户名;column:username;size:64;" binding:"required"`  //用户名
  Mobile  *string `json:"mobile" form:"mobile" gorm:"comment:手机号;column:mobile;" binding:"required"`  //手机号
  Image  *string `json:"image" form:"image" gorm:"comment:头像;column:image;size:255;"`  //头像
  Password  *string `json:"password" form:"password" gorm:"comment:密码;column:password;size:128;" binding:"required"`  //密码
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName users表 Users自定义表名 users
func (Users) TableName() string {
    return "users"
}





