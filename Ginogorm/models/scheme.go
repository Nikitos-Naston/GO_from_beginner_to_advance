package models

import "github.com/jinzhu/gorm"

type Arcticle struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"contnent"`
}

func (a *Arcticle) TableName() string {
	return "arcticle"
}
