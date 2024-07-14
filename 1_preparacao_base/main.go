package main

import (
	"context"
	"database/sql"
	"fmt"
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

	p, err := selecionarUmProduto(db, produto.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("O produto %s, possui o preço de %.2f \n", p.Nome, p.Preco)

	produtos, err := selecionarTodosProdutos(db)
	if err != nil {
		panic(err)
	}
	for _, p := range produtos {
		fmt.Printf("O produto %s, possui o preço de %.2f \n", p.Nome, p.Preco)
	}
}

func selecionarUmProduto(db *sql.DB, id string) (*Produto, error) {
	stmt, err := db.Prepare("select id, name, price from products where id =?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Produto

	//err = stmt.QueryRow(id).Scan(&p.ID, &p.Nome, &p.Preco)
	err = stmt.QueryRowContext(context.Background(), id).Scan(&p.ID, &p.Nome, &p.Preco)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selecionarTodosProdutos(db *sql.DB) ([]Produto, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var produtos []Produto

	for rows.Next() {
		var p Produto
		err = rows.Scan(&p.ID, &p.Nome, &p.Preco)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}
	return produtos, nil
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
