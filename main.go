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
	"fmt"
	"os"
)

type Data struct {
	Name  string
	Age   int
	Score float32
}

func main() {

	arguments := os.Args[1:]

	if len(arguments) < 5 {
		/*var data = []Data{
			Data{Name: "John Smith", Age: 30, Score: 100.0},
			Data{Name: "Jane Smith", Age: 30, Score: 100.0},
		}
		olog.Print(data)*/

		fmt.Println(arguments[1])
		fmt.Println(arguments[2])

	} else {
		fmt.Println("To many arguements. Excepted arguments are -d for delemeter and -h for headers.")
	}

}
