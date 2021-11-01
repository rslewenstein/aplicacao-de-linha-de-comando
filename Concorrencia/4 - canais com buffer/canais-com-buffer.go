package main

import "fmt"

func main() {
	canal := make(chan string, 2) // 2 = Capacidade total do buffer. Se passar da capacidade, dá dead lock.
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
