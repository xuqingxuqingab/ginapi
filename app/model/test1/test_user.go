package test1

import "gorm.io/gorm"

type TestUser struct {
	UserId    int    `gorm:"primaryKey" json:"user_id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	Account   string `json:"account"`
	CreatedAt int
	UpdatedAt int
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
