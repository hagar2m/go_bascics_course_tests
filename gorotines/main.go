package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Millisecond * 800)
	}
	channel <- "Done"
}

func main() {
	var chanel chan string
	go printMessage("Hagggar", chanel)
	// printMessage("Start in Goooooo", chanel)
	res := <- chanel
	fmt.Println(res)
}
