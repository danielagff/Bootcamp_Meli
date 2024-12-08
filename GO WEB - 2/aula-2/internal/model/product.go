package model

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p *Product) Validate() error {
    if utf8.RuneCountInString(p.Name) == 0 {
        return fmt.Errorf("nome não pode ser vazio")
    }
    if p.Quantity < 0 {
        return fmt.Errorf("quantidade não pode ser negativa")
    }
    if p.Price < 0 {
        return fmt.Errorf("preço não pode ser negativo")
    }
    if _, err := time.Parse("02/01/2006", p.Expiration); err != nil {
        return fmt.Errorf("data de expiração inválida")
    }
    return nil
}
