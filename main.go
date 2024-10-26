package main

import (
	"fmt"
	"time"
)

//TimeOut

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

	//Non-Blocking Channel Operations

	messages := make(chan string)
	signals := make(chan bool)

	// Здесь показан неблокирующий приём.
	// Если значение доступно
	// в `messages`, то `select` выберет вариант
	// `<-messages` с этим значением. Если нет,
	// то будет сразу выбран вариант `default`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Неблокирующая отправка работает аналогично.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Можно использовать несколько `case` перед
	// оператором `default` для реализации многовариантного
	// неблокирующего выбора. Здесь показана попытка
	// неблокирующим способом получить
	// как `messages` так и `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
