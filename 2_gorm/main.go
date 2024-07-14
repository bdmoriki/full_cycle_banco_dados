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
	db.AutoMigrate(&Produto{})

	// db.Create(&Produto{
	// 	Nome:  "Microfone HyperX",
	// 	Preco: 500.00})

	// create batch
	// produtos := []Produto{
	// 	{Nome: "Teclado", Preco: 200.50},
	// 	{Nome: "Monitor", Preco: 3000.00},
	// 	{Nome: "Mouse", Preco: 150.50}}

	// db.Create(produtos)

	// var produto Produto
	// db.First(&produto, 1)
	// fmt.Println(produto)

	// db.First(&produto, "nome = ?", "Teclado")
	// fmt.Println(produto)

	var produtos []Produto
	db.Find(&produtos)

	for _, p := range produtos {
		fmt.Println(p)
	}
}
