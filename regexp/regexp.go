package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("efo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eko edo egi ego eto e1o eFo", 10))
}
