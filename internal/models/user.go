package models

import "time"

type User struct {
	ID          int        `gorm:"primaryKey;autoIncrement;comment:ID" json:"id"`
	Username    string     `gorm:"size:255;not null;unique;comment:用户名" json:"username"`
	Password    string     `gorm:"size:255;not null;comment:密码" json:"-"`
	Status      int8       `gorm:"not null;default:1;comment:状态" json:"status"`
	Role        string     `gorm:"size:255;not null;default:'user';comment:角色" json:"role"`
	Avatar      string     `gorm:"size:255;not null;default:'https://cdn-static.xxcheng.cn/static/uploads/23d61d24c8.png';comment:头像" json:"avatar"`
	Version     int        `gorm:"not null;default:1;comment:版本号" json:"version"`
	CreatedTime *time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdTime"`
	UpdatedTime *time.Time `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updatedTime"`
}

// TableName 设置表名
func (u *User) TableName() string {
	return "users"
}

func (u *User) StatusOk() bool {
	return u.Status == 1
}

func (u *User) SuperAdmin() bool {
	return u.Role == "super"
}
