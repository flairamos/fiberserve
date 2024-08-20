package tables

import (
	"sunny/pkg/custom"
)

// User 模型
type User struct {
	custom.Id
	Name string `gorm:"type:varchar(20);unique;comment:账户" json:"name,omitempty"`
	custom.At
}

// TableName 表名
func (User) TableName() string {
	return "user"
}

// TableComment 方法返回表的注释
func (User) TableComment() string {
	return "账户"
}
