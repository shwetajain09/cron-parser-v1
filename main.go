package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cron-parser/V1/parser"
)

func main() {
	// read input from console in a string format
	fmt.Println("Enter a cron parser expression string")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// split a string around consecutive whitespace characters.
	strSlice := strings.Fields(text)

	// validate and parse string to an expression format
	expression, err := parser.ParseStringArray(strSlice)
	if err != nil {
		fmt.Println(err)
		return
	}

	// build response for each field
	res, err := parser.BuildResponse(expression)
	if err != nil {
		fmt.Println(err)
		return
	}

	// render table output
	parser.RenderOutput(res)
}
