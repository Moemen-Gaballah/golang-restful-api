package main

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);" json:"title" binding:"required"`
	Description string `gorm:"type:varchar(500);" json:"description" binding:"required"`
	Status      string `json:"status"`
}
