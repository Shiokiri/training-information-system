package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Register) TableName() string {
	return "register"
}

type Register struct {
	gorm.Model
	courseId  string `json:"courseId"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Company   string `json:"company"`
	Post      string `json:"Post"`
	Level     string `json:"Level"`
	Contact   string `json:"contact"`
	Status    string `json:"Status"`
}
