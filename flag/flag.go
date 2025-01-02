package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("hots", "localhost", "Put Your Database Host")
	username := flag.String("username", "root", "Put Your Database Username")
	password := flag.String("password", "root", "Put Your Database Password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
