package main

import "fmt"

func receiver(ch chan int){
	fmt.Println(<-ch) // this go routime will be blocked until it reads the value from the channel ch
	for val := range ch{
		fmt.Println(val)
	}
}
func main(){
	fmt.Println("Channels") // it prints channels

	ch := make(chan int) // then creates a channel

	go receiver(ch) // starts another go routine to receive the value

	ch <- 42
	ch <- 43
	ch <- 44
	ch <- 45
	close(ch) // to close the channel  
}