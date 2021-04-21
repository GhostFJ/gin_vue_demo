package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId uint `json:"user_id" gorm:"not null"`
	Category *Category
	CategoryId uint `json:"category_id" gorm:"not null"`
	Title string `json:"title" gorm:"type:varchar(50); not null"` 
	HeadImg string `json:"head_img"`
	Content string `json:"content" gorm:"tyep:text;not null"`
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"`
}

// 添加Hook来修改ID为uuid格式达到
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}