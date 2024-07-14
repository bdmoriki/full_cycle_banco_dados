package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Produto struct {
	ID    string
	Nome  string
	Preco float64
}

func NewProduto(nome string, preco float64) *Produto {
	return &Produto{
		ID:    uuid.New().String(),
		Nome:  nome,
		Preco: preco}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	produto := NewProduto("Notebook", 12000)
	err = inserirProduto(db, produto)
	if err != nil {
		panic(err)
	}

	produto.Nome = "Macbook"

	err = atualizarProduto(db, produto)
	if err != nil {
		panic(err)
	}
}

func inserirProduto(db *sql.DB, produto *Produto) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.ID, produto.Nome, produto.Preco)
	if err != nil {
		return err
	}
	return nil
}

func atualizarProduto(db *sql.DB, produto *Produto) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	log.Printf("Atualizando o produto com o id %s \n", produto.ID)
	_, err = stmt.Exec(produto.Nome, produto.Preco, produto.ID)
	if err != nil {
		return err
	}
	return nil
}
