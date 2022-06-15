package main

import (
	"fmt"
	"time"
)

func main() {
	go sleeprint(1)
	go sleeprint(2)
	go sleeprint(3)
	go sleeprint(4)
	go sleeprint(5)

	select {}
}

func sleeprint(n int) {
	for {
		fmt.Print(n)
		time.Sleep(time.Second * time.Duration(n))
	}
}
