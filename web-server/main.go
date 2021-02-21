package main

import (
	"ecohanim/echo-server-ex/web-server/router"
	"fmt"
)

func main() {
	fmt.Println("\nWelcome to CHQ Web Server")

	e := router.New()

	e.Start(":9090")
}
