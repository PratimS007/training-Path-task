package main

import "fmt"

// struct is created

type employee struct {
    name   string
    age    int
    salary int
}

func main() {
    emp1 := employee{} // empyty struct as default value is zero
    fmt.Printf("Emp1: %+v\n", emp1)

    emp2 := employee{name: "Sam", age: 31, salary: 2000}
    fmt.Printf("Emp2: %+v\n", emp2)

    emp3 := employee{
        name:   "Sam",
        age:    31,
        salary: 2000,
    }
    fmt.Printf("Emp3: %+v\n", emp3)
	fmt.Printf("Emp3: %+v\n", emp3.name) // to access the struct field using dot operator

    emp4 := employee{ // all the field are not initialized so itw ill take zero default value
        name: "Sam",
        age:  31,
    }
    fmt.Printf("Emp4: %+v\n", emp4)
}