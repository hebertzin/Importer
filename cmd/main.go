package main

import (
	"enube-challenge/packages/database"
	"fmt"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Hello from go!!")
}
