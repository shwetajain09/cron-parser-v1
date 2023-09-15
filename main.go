package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cron-parser/V1/parser"
)

func main() {
	// read input from console in a long string format
	fmt.Println("Enter a cron parser expression string")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.Fields(text)

	// validate and parse string to an expression format
	expression, err := parser.ParseStringArray(str)
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
