package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Categoria struct {
	ID       int `gorm:"primaryKey"`
	Nome     string
	Produtos []Produto
}

type Produto struct {
	ID           int `gorm:"primaryKey"`
	Nome         string
	Preco        float64
	CategoriaID  int
	Categoria    Categoria
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProdutoID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8,b4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Produto{}, &Categoria{}, &SerialNumber{})

	// criar categoria
	categoria := Categoria{Nome: "Eletronicos"}
	db.Save(&categoria)

	categoria2 := Categoria{Nome: "Cozinha"}
	db.Save(&categoria2)

	// criar produto
	produtos := []Produto{
		{Nome: "Monitor", Preco: 2000, CategoriaID: categoria.ID},
		{Nome: "Garfo", Preco: 2000, CategoriaID: categoria2.ID},
	}
	db.Save(&produtos)

	// criar serial number
	db.Create(&SerialNumber{Number: "S123", ProdutoID: 1})

	var categorias []Categoria
	//err = db.Model(&Categoria{}).Preload("Produtos").Preload("Produtos.SerialNumber").Find(&categorias).Error
	err = db.Model(&Categoria{}).Preload("Produtos.SerialNumber").Find(&categorias).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categorias {
		fmt.Println("Categoria: ", c.Nome)
		for _, p := range c.Produtos {
			fmt.Println("- Produto: ", p.Nome, ", Serial number:", p.SerialNumber.Number)
		}
	}
}
