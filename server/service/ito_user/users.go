
package ito_user

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ito_user"
    ito_userReq "github.com/flipped-aurora/gin-vue-admin/server/model/ito_user/request"
    "gorm.io/gorm"
)

type UsersService struct {}
// CreateUsers 创建users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) CreateUsers(ctx context.Context, users *ito_user.Users) (err error) {
	err = global.GVA_DB.Create(users).Error
	return err
}

// DeleteUsers 删除users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)DeleteUsers(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ito_user.Users{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&ito_user.Users{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteUsersByIds 批量删除users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)DeleteUsersByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&ito_user.Users{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&ito_user.Users{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateUsers 更新users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)UpdateUsers(ctx context.Context, users ito_user.Users) (err error) {
	err = global.GVA_DB.Model(&ito_user.Users{}).Where("id = ?",users.ID).Updates(&users).Error
	return err
}

// GetUsers 根据ID获取users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)GetUsers(ctx context.Context, ID string) (users ito_user.Users, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&users).Error
	return
}
// GetUsersInfoList 分页获取users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)GetUsersInfoList(ctx context.Context, info ito_userReq.UsersSearch) (list []ito_user.Users, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ito_user.Users{})
    var userss []ito_user.Users
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Username != nil && *info.Username != "" {
        db = db.Where("username = ?", *info.Username)
    }
    if info.Mobile != nil && *info.Mobile != "" {
        db = db.Where("mobile = ?", *info.Mobile)
    }
    if info.Image != nil && *info.Image != "" {
        db = db.Where("image = ?", *info.Image)
    }
    if info.Password != nil && *info.Password != "" {
        db = db.Where("password = ?", *info.Password)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userss).Error
	return  userss, total, err
}
func (usersService *UsersService)GetUsersPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
