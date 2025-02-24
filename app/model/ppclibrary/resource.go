package ppclibrary

import (
	"gorm.io/gorm"
)

// User 定义了一个用户模型
type Resource struct {
	ResourceId int    `gorm:"primaryKey" json:"resource_id"`
	Name       string `gorm:"size:100;not null" json:"name"`
	CreatedAt  int
	UpdatedAt  int
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
