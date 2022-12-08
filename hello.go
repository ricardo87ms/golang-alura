package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Ricardo"
	var idade = 35
	var versao = 2.3

	fmt.Println("Olá mundo", nome, "sua idade é", idade)
	fmt.Println("A versão deste programa é", versao)

	fmt.Println("Tipo variável nome", reflect.TypeOf(nome))
	fmt.Println("Tipo variável idade", reflect.TypeOf(idade))
	fmt.Println("Tipo variável versao", reflect.TypeOf(versao))
}
