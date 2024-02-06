package main

import "testing"

func TestHellow(t *testing.T) {

	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper() //Ao fazer isso, quando o teste falhar, o número da linha relatada estará em nossa chamada de função, e não dentro do nosso auxiliar de teste.
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("diz hello para as pessoas", func(t *testing.T) {
		resultado := Hello("Word", "")
		esperado := "Hello, Word"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Hello, Word' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Hello("", "")
		esperado := "Hello, Word"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em, espanhol", func(t *testing.T) {
		resultado := Hello("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Hello("Princeso", "francês")
		esperado := "Bonjour, Princeso"
		verificarMensagemCorreta(t, resultado, esperado)
	})
}
