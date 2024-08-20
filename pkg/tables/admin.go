package tables

import (
	"sunny/pkg/custom"
)

// Admin 管理员模型
type Admin struct {
	custom.Id
	Name string `gorm:"type:varchar(20);index;comment:姓名" json:"name,omitempty"`
	Pwd  string `gorm:"type:varchar(60);comment:密码" json:"pwd,omitempty"`
	custom.TenantId
	custom.At
	custom.DeletedAt
}

// TableName 表名
func (Admin) TableName() string {
	return "admin"
}

// TableComment 方法返回表的注释
func (Admin) TableComment() string {
	return "管理员"
}
