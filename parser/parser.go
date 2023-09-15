package parser

import "errors"

func ParseStringArray(str []string) (*body, error) {
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
