package main

import "fmt"

func main(){
	fmt.Println("pointers")

	var a int = 42
	fmt.Printf("Value = %v , Type = %T\n",a , a)

	var b *int = &a
	fmt.Printf("Value = %v , Type = %T\n", b , b)
	fmt.Printf("Value = %v , Type = %T\n", *b , *b)

}
