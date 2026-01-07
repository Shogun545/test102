package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `valid:"required~Title is required"`
	Author string `valid:"required~Author is required"`
	Price float64 `valid:"required~Price is required,range(1|10000)~Price must be between 1 and 10000"`
}