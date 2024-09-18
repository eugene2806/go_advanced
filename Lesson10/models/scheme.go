package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	// ID        uint `gorm:"primary_key"`
	//    CreatedAt time.Time
	//    UpdatedAt time.Time
	//    DeletedAt *time.Time `sql:"index"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

// Требование gorm - наличие метода возвращающего имя таблицы
func (a *Article) TableName() string {
	return "article"
}
