package main

import "fmt"

func main()  {

	var numero int
	numero = 25

	fmt.Println(numero)
	numero = 40
	fmt.Println(numero)

	nombre := "Alejandro"
	fmt.Println(nombre)

	var num = uintptr(numero)

	fmt.Println(num)

	var entero32 int32
	var entero64 int64

	entero32 = 45
	entero64 = 89

	fmt.Println(entero32+int32(entero64))

}
