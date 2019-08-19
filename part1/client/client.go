package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("rcp.Dial failed. [%v]", err)
	}
	var reply string
	err = client.Call("EchoService.Echo", "golang rpc", &reply)
	if err != nil {
		log.Fatalf("client.Call failed. [%v]", err)
	}
	fmt.Println(reply)
}
