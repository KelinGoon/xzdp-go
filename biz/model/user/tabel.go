package user

import (
	"time"

	"gorm.io/gorm"
)

func (User) TableName() string {
	return "tb_user"
}

// BeforeCreate 是一个 GORM 钩子函数，在插入数据前执行
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = time.Now().String()
	u.UpdateTime = time.Now().String()
	return
}
