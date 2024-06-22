/*
CSV viewer takes in a json file to view first few lines of file in the terminal.
This way you can see the way the file looks without opening the file. There are
other tools that do this but I am also using it to learn go.

Usage:
	csvfiew [flags] [path]

The flags are:
	-d
		This is the delemeter for the file. Ex -d ; or -d ,
	-h
		If the CSV has headers. Ex -h if headers exist
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/da0x/golang/olog"
)

type Data struct {
	Username  string
	ID        string
	FirstName string
	LastName  string
}

func readFile(file string) string {
	body, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("Unable to read file: : %v", err)
	}

	return string(body)
}

func parseCsvContents(contents string) []Data {
	var displayData = []Data{}
	scanner := bufio.NewScanner(strings.NewReader(contents))

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {

			data := strings.Split(scanner.Text(), ";")
			fmt.Println(data)
			displayData = append(displayData, Data{Username: data[0], ID: data[1], FirstName: data[2], LastName: data[3]})
		}
	}

	return displayData

}

func main() {

	arguments := os.Args[1:]

	if len(arguments) < 5 {
		/*var data = []Data{
			Data{Name: "John Smith", Age: 30, Score: 100.0},
			Data{Name: "Jane Smith", Age: 30, Score: 100.0},
		}
		olog.Print(data)*/

		fmt.Println(arguments[0])
		contents := readFile(arguments[0])
		result := parseCsvContents(contents)
		olog.Print(result)

	} else {
		fmt.Println("To many arguements. Excepted arguments are -d for delemeter and -h for headers.")
	}

}
