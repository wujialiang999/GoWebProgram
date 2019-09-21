package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}
func RPCServer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("Server received:", data)
		ch <- "reger"
	}
}
func main() {
	ch := make(chan string)
	go RPCServer(ch)
	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Client received", recv)
	}
}
