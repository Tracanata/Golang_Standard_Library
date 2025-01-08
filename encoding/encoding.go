package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/*<-- Base 64 dengan menggunakan Encoding -->*/
	value := "Fatra Candra Ardiwinata"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	/*<-- Membuat CSV Reader -->*/
	csvString := "Ayunda, Risu\n" +
		"Moona, Hoshinova\n" +
		"Airani, Iofifteen"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	/*<-- Membuat CSV Writer -->*/
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Fatra", "Candra", "Ardiwinata"})
	_ = writer.Write([]string{"Ayunda", "Risu"})
	_ = writer.Write([]string{"Airani", "Iofifteen"})

	writer.Flush()
}
