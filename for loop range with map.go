package main

import "fmt"

func main(){
	var mp map[string]int = map[string]int{
		"apple":5,
		"orange":9,
		"pear":6,
	}
	fmt.Println(mp["apple"]) //to access the value
	fmt.Println(mp)

	delete(mp,"apple") //to delete the key from the map

    mp["tim"] = 900 // to update the key value
}