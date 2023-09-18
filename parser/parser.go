package parser

import "errors"

// this would parse the string array to a body object that represents the expression recieved through the string input
func ParseStringArray(str []string) (*body, error) {
	// the expression should have exactly 6 fields
	if len(str) != 6 {
		return nil, errors.New("invalid expression")
	}
	b := &body{
		minutes:    str[0],
		hour:       str[1],
		dayOfMonth: str[2],
		month:      str[3],
		dayOfWeek:  str[4],
		command:    str[5],
	}
	return b, nil
}
