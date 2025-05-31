package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:100;not null"`
	Price     float64 `gorm:"not null"`
	Stock     int     `gorm:"default:0"`
	CreatedAt string  `gorm:"autoCreateTime"`
	UpdatedAt string  `gorm:"autoUpdateTime"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	var product Product
	db.First(&product)

	fmt.Println(product.Name)
}
