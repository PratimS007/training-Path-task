// package main

// import "fmt"

// func main(){
// 	x := 7
// 	y := &x
// 	fmt.Println(x,y)
// }



// package main

// import "fmt"

// func main() {
//     //Declare
//     var b *int
//     a := 2
//     b = &a
    
//     //Will print a address. Output will be different everytime.
//     fmt.Println(b)
//     fmt.Println(*b)
//     b = new(int)
//     *b = 10
//     fmt.Println(*b) 
// }


// Example of default zero value of a pointer

package main

import "fmt"

func main() {
    var a *int
    fmt.Print("Default Zero Value of a pointer: ")
    fmt.Println(a)
}