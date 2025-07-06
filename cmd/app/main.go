package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/JamesCookDev/Treino-Go/greetings"
)
func main() {
	fmt.Println("--- Saída da biblioteca local 'greetings' ---")
	fmt.Println(greetings.Hello("James"))
	fmt.Println("--- Saída da biblioteca externa 'quote' ---")
	fmt.Println(quote.Go())
	fmt.Println("--- Fim da execução ---")
}
