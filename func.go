package main 

import "fmt"


func mult(x int, y int) int {
	return x * y
}


func nome(nome string, sobrenome string) string{
	fmt.Println('My name is', nome, 'and sobrenome is', sobrenome)

}

func main(){
	nome("Anderson", "Cezar")
}