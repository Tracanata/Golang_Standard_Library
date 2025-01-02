package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fatra Ardiwinata", "Ardi"))
	fmt.Println(strings.Split("Fatra Ardiwinata", " "))
	fmt.Println(strings.ToLower("Fatra Ardiwinata"))
	fmt.Println(strings.ToUpper("Fatra Ardiwinata"))
	fmt.Println(strings.Trim("    Fatra     ", " "))
	fmt.Println(strings.ReplaceAll("Fatra Fatra Fatra", "Fatra", "Ardiwinata"))
}
