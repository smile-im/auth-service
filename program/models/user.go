package models

import (
	"github.com/jinzhu/gorm"
	"github.com/micro-kit/micro-common/db"
	"github.com/smile-im/microkit-client/proto/authpb"
)

// User 表名-无前缀
type User struct {
	UserId        int64  `json:"user_id" gorm:"column:user_id"`
	Username      string `json:"username" gorm:"column:username"`
	Nickname      string `json:"nickname" gorm:"column:nickname"`
	Password      string `json:"password,omitempty" gorm:"column:password"`
	HeadSculpture string `json:"head_sculpture" gorm:"column:head_sculpture"`
	Desc          string `json:"desc" gorm:"column:desc"`
	RegisterAt    int64  `json:"register_at" gorm:"column:register_at"`
	CreatedAt     int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     int64  `json:"updated_at" gorm:"column:updated_at"`
}

// TableName 表名
func (User) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "user")
}

// UserLogin 根据用户名密码查询数据库 - login使用
func (m *User) UserLogin(username, password string) error {
	return db.GetSlaveDB().Model(m).Where("username = ? and password = ?", username, password).Find(m).Error
}

// 转pb对象
func (m *User) ToPb() *authpb.UserInfo {
	return &authpb.UserInfo{
		UserId:        m.UserId,
		Username:      m.Username,
		Nickname:      m.Nickname,
		HeadSculpture: m.HeadSculpture,
		Desc:          m.Desc,
		RegisterAt:    m.RegisterAt,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}
