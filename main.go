package main

import "fmt" //pacote que tem a função Println

const frances = "francês"
const prefixHelloFrances = "Bonjour, "
const espanhol = "espanhol"
const prefixHelloEspanhol = "Hola, "
const prefixHelloPortugues = "Hello, "

func Hello(nome string, idioma string) string {

	if nome == "" {
		nome = "Word"
	}

	return prefixodeSaudacao(idioma) + nome

}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixHelloEspanhol
	case frances:
		prefixo = prefixHelloFrances
	default:
		prefixo = prefixHelloPortugues
	}

	return
}
func main() {
	fmt.Println(Hello("junior", "francês"))
}
