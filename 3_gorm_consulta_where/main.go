package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Produto struct {
	ID    int `gorm:"primaryKey"`
	Nome  string
	Preco float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// var produtos []Produto
	// db.Limit(2).Offset(2).Find(&produtos)

	// for _, p := range produtos {
	// 	fmt.Println(p)
	// }

	// var produtos []Produto
	// db.Where("preco > ?", 2000).Find(&produtos)
	// for _, p := range produtos {
	// 	fmt.Println(p)
	// }

	var produtos []Produto
	db.Where("nome LIKE ?", "%Mon%").Find(&produtos)
	for _, p := range produtos {
		fmt.Println(p)
	}
}
