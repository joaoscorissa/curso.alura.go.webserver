package models

import "curso.alura.web_app/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaProdutos() []Produto {
	db := db.ConnectDb()
	selectProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConnectDb()

	removeProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	removeProduto.Exec(id)
	defer db.Close()
}

func GetProduto(id string) Produto {
	db := db.ConnectDb()

	produto, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for produto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}

	defer db.Close()
	return p
}

func UpdateProduto(id, quantidade int, nome, descricao string, preco float64) {
	db := db.ConnectDb()

	removeProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, quantidade=$3, preco=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	removeProduto.Exec(nome, descricao, quantidade, preco, id)
	defer db.Close()
}
