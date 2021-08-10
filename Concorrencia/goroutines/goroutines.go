package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine é utilizado para iniciar um método/função
	// e não precisa esperar o término da execução para iniciar um outro método/função

	go escrever("Olá Mundo!") //goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
