/*
Uma universidade precisa registrar os alunos e gerar uma funcionalidade para imprimir
os detalhes dos dados de cada aluno, como segue:
Nome: [Primeiro nome do aluno]
Sobrenome: [Sobrenome do aluno]
ID: [ID do aluno]
Data: [Data de admissão do aluno]
Os valores entre colchetes devem ser substituídos pelos
dados fornecidos pelos alunos. Para isso, é necessário gerar
uma estrutura Students com as variáveis Name, Surname, DNI, Date e que tenha um método de detalhamento
*/
package main

import "fmt"

type Students struct {
	ID      int
	Name    string
	Surname string
	DNI     string
	Date    string
}

func (a *Students) obterDetalhesAluno() {
	fmt.Printf("ID : %v\nName: %vSurname: %v\nData:%v\nDNI:%v\n ", a.ID, a.Name, a.Surname, a.Date, a.DNI)
}

func main() {
	newStudent := Students {
		ID: 1,
		Name: "Daniel",
		Surname: "Filho",
		DNI: "seilaoqeisso",
		Date: "12/09/2024",
	}

	newStudent.obterDetalhesAluno()
}
