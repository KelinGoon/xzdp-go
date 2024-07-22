package mysql

import (
	"context"
	"time"
	model "xzdp/biz/model/user"
	"xzdp/biz/utils"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	Phone      string    `thrift:"Phone,1" json:"phone" gorm:"phone"`
	Password   string    `thrift:"password,3" form:"password" json:"password" gorm:"password"`
	ID         uint      `gorm:"primarykey"`
	NickName   string    `thrift:"NickName,5" form:"NickName" json:"NickName" gorm:"nick_name"`
	Icon       string    `thrift:"icon,6" form:"icon" json:"icon" gorm:"icon"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `thrift:"updateTime,8" form:"updateTime" json:"updateTime" gorm:"update_time"`
}

func GetById(ctx context.Context, id int64) (*model.User, error) {

	var user model.User
	db := DB.Model(&model.User{})

	// Perform the query
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle case where no user is found
			return nil, utils.ErrNotFound
		}
		// Handle other potential errors
		return nil, err
	}

	// Return the user
	return &user, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	return
}

func (User) TableName() string {
	return "tb_user"
}
