/*Algumas lojas de comércio eletrônico precisam criar uma funcionalidade no Go para gerenciar
produtos e retornar o valor do preço total. A empresa tem três tipos de produtos: 
Pequeno, Médio e Grande (muitos outros são esperados).
E os custos adicionais são:
Pequeno: apenas o custo do produto
Médio: o preço do produto + 3% do produto + 3% de mantê-lo na loja
Grande: o preço do produto + 6% de mantê-lo na loja e, além disso, o custo de envio de US$ 2.500.
O custo de manter o produto em estoque na loja é uma porcentagem do preço do produto.
É necessária uma função factory que receba o tipo de produto e o preço e retorne uma interface Product que tenha o método Price.
Deve ser possível executar o método Price e fazer com que o método retorne o preço total com base no custo do produto e em quaisquer custos adicionais.
*/
package main

import (
	"fmt"
)

type Produto interface {
	Preco() float64
}

type ProdutoPequeno struct {
	preco float64
}

func (p ProdutoPequeno) Preco() float64 {
	return p.preco
}

type ProdutoMedio struct {
	preco float64
}

func (p ProdutoMedio) Preco() float64 {
	return p.preco + (p.preco * 0.03) + (p.preco * 0.03)
}

type ProdutoGrande struct {
	preco float64
}

func (p ProdutoGrande) Preco() float64 {
	return p.preco + (p.preco * 0.06) + 2500
}

func FabricaDeProdutos(tipoProduto string, preco float64) (Produto, error) {
	switch tipoProduto {
	case "pequeno":
		return ProdutoPequeno{preco: preco}, nil
	case "medio":
		return ProdutoMedio{preco: preco}, nil
	case "grande":
		return ProdutoGrande{preco: preco}, nil
	default:
		return nil, fmt.Errorf("tipo de produto desconhecido: %s", tipoProduto)
	}
}

func main() {
	produtos := []struct {
		tipoProduto string
		preco       float64
	}{
		{"pequeno", 100.0},
		{"medio", 100.0},
		{"grande", 100.0},
		{"desconhecido", 100.0},
	}

	for _, p := range produtos {
		produto, err := FabricaDeProdutos(p.tipoProduto, p.preco)
		if err != nil {
			fmt.Printf("Erro ao criar produto do tipo %s: %v\n", p.tipoProduto, err)
			continue
		}
		fmt.Printf("Produto do tipo %s, Preço total: %.2f\n", p.tipoProduto, produto.Preco())
	}
}