package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Produto struct {
	ID    int `gorm:"primaryKey"`
	Nome  string
	Preco float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8,b4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Produto{})
	db.Create(&Produto{
		Nome:  "Fone earbud",
		Preco: 200,
	})

	var p Produto
	db.First(&p, 1)

	p.Nome = "Alexa"
	db.Save(&p)

	db.Delete(&p)
}
