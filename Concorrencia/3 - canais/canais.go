package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal //O canal vai esperar receber um valor

	// 	// verifica se o canal ainda está aberto. Canal aberto sem receber dados gera deadlock.
	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	
	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // a variável canal recebe(<-)
		time.Sleep(time.Second)
	}

	close(canal) // fecha o canal para evitar erros deadlock
}
