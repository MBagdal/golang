package main

import (
	"fmt"
	"os"
)

func main() {
	showHeader()

	opc, erro := readUserCommand()

	if erro == nil {
		//ChooseOpcIfStatment(opc)
		ChooseOpcSweetStatment(opc)
	}
}

func ChooseOpcIfStatment(opc int) {
	if opc == 1 {
		fmt.Println("Você escolheu a opção MONITORAR SITES")
	} else if opc == 2 {
		fmt.Println("Você escolheu a opção VER LOGS")
	} else {
		fmt.Println("Opção inválida")
	}
}

func ChooseOpcSweetStatment(opc int) {
	switch opc {
	case 1:
		fmt.Println("Você escolheu a opção MONITORAR SITES")
	case 2:
		fmt.Println("Você escolheu a opção VER LOGS")
	default:
		fmt.Println("Opção inválida")
		os.Exit(-1)
	}
}

func showHeader() {
	fmt.Println("Olá, bem vindo ao rastreador de sites")
	fmt.Println("Digite: 1 - Monitorar sites, 2 - Ver logs, 3 - Sair")
}

func readUserCommand() (opc int, err error) {
	_, erro := fmt.Scan(&opc)

	return opc, erro
}
