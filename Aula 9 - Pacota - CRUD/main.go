package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Student struct {
	Id          string
	Name        string
	Email       string
	PhoneNumber string
}

func CreateNewStudent() string {
	var name, email, phoneNumber string

	fmt.Println("Digite o nome do aluno: ")
	fmt.Scanln(&name)
	fmt.Println("Digite o email do aluno: ")
	fmt.Scanln(&email)
	fmt.Println("Digite o telefone do aluno: ")
	fmt.Scanln(&phoneNumber)

	newStudent := Student{
		Id:          uuid.New().String(),
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}

	return fmt.Sprint(newStudent.Id, ",", newStudent.Name, ",", newStudent.Email, ",", newStudent.PhoneNumber, "\n")
}

func GetAllStudents() {
	file, err := os.Open("students.csv")

	if err != nil {
		fmt.Printf("Erro ao abrir o arquio: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, ",")

		fmt.Printf("\nId : %v\nName : %v\nEmail : %v\nPhone number: %v\n", columns[0], columns[1], columns[2], columns[3])
		fmt.Println("------------------------------------------------")
	}

}

func GetStudentsById() {
	var idStudents string

	fmt.Println("Digite o Id do aluno: ")
	fmt.Scanln(&idStudents)

	if idStudents != "" {

		file, err := os.Open("students.csv")

		if err != nil {
			fmt.Printf("Erro ao abrir o arquio: %v", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			columns := strings.Split(line, ",")

			fmt.Printf("\nId : %v\nName : %v\nEmail : %v\nPhone number: %v\n", columns[0], columns[1], columns[2], columns[3])
			fmt.Println("------------------------------------------------")
		}
	}

}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func printHeader() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|       SISTEMA DE ALUNOS        |")
	fmt.Println("+--------------------------------+")
}

func showMenu() {
	printHeader()
	fmt.Println("|                                |")
	fmt.Println("| 1 - Criar aluno                |")
	fmt.Println("| 2 - Editar aluno               |")
	fmt.Println("| 3 - Buscar aluno               |")
	fmt.Println("| 4 - Excluir aluno              |")
	fmt.Println("|                                |")
	fmt.Println("| 9 - Sair                       |")
	fmt.Println("|                                |")
	fmt.Println("+--------------------------------+")
	fmt.Println("Sua Opção :")
}

func main() {

	file, err := os.OpenFile("students.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0655)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var optionUser int

	for optionUser != 9 {
		clearScreen()
		showMenu()
		fmt.Scanln(&optionUser)
		switch optionUser {
		case 1:
			clearScreen()
			printHeader()
			res := CreateNewStudent()
			_, err := file.WriteString(res)
			if err != nil {
				fmt.Printf("Erro ao salvar arquivo: %v", err)
				return
			}
			fmt.Printf("Registro incluido!")
			time.Sleep(3 * time.Second)

		case 2:
			clearScreen()
			printHeader()
			GetAllStudents()
			fmt.Printf("Pressione qualquer tecla para sair!")
			fmt.Scanln()
		}
	}

}
