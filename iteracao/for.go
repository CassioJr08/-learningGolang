package iteracao

const qtdeRepeticoes = 5

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < qtdeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
