package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categoria struct {
	ID   int `gorm:"primaryKey"`
	Nome string
}

type Produto struct {
	ID          int `gorm:"primaryKey"`
	Nome        string
	Preco       float64
	CategoriaID int
	Categoria   Categoria
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8,b4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Produto{}, &Categoria{})

	// criar categoria
	categoria := Categoria{Nome: "Eletronicos"}
	db.Save(&categoria)

	// criar produto
	produto := Produto{Nome: "Monitor", Preco: 2000, CategoriaID: categoria.ID}
	db.Save(&produto)

	produto2 := Produto{Nome: "Mouse", Preco: 2000, CategoriaID: categoria.ID}
	db.Save(&produto2)

	var produtos []Produto
	db.Preload("Categoria").Find(&produtos)
	for _, p := range produtos {
		fmt.Println(p.Nome, p.Categoria.Nome)
	}
}
