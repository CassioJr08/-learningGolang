package main

import "fmt" //pacote que tem a função Println

func Ola(nome string) string {
	return "Olá, " + nome
}
func main() {
	fmt.Println(Ola("mundo"))
}
