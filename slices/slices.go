package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Fatra", "Candra", "Ardiwinata"}
	values := []int{100, 95, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Cahya"))
	fmt.Println(slices.Index(names, "Cahya"))
	fmt.Println(slices.Index(names, "Fatra"))
}
