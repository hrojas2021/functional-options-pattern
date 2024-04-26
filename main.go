package main

import (
	"fmt"

	"github.com/hrojas2021/functional-options-pattern/server"
)

func main() {
	defaultServer := server.NewServer()
	fmt.Printf("default server: %+v\n", defaultServer.GetConfig())

	s := server.NewServer(server.WithTLS, server.WithMaxConn(45), server.WithServerID("local-server"))
	fmt.Printf("server: %+v\n", s.GetConfig())
}
