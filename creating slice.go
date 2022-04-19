package main

import "fmt"

func main(){
	fmt.Println("slices")

	a := []int{1,2,3,4,5,6,7,8,9,10}

	//creating slice from another slice

	b := a[:]
	c := a[3:]
	// d := a[:6]
	e := a[3:6]

	// modifying the value in array and it will reflect change in slice too and 0 and 1 is the index position

	// d[0] = 100
	// b[1] = 200

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// fmt.Println(d)
	fmt.Println(e)



// creating slice using make function

// a = make([]int, 3 , 10 )
// fmt.Println(a,len(a),cap(a))

// append a value to the slice and also multiple values

// a = append(a, 100,200,300)
// fmt.Println(a,len(a),cap(a))

// appending a value to a empty slice

// a = make([]int, 0 , 0 )
// a = append(a, 100,200,300)
// fmt.Println(a,len(a),cap(a))

}