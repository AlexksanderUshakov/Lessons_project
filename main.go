package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:

		fmt.Print(res)
	case <-time.After(time.Second * 1):
		fmt.Print("timeout 1\n")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 2"
	}()

	select {
	case res := <-c2:

		fmt.Print(res)
	case <-time.After(time.Second * 3):
		fmt.Print("timeout 2")
	}

}
