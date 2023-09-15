package parser

import (
	"errors"
	"reflect"
	"testing"
)

var mockBodyResponse = &body{
	minutes:    "1",
	hour:       "1",
	month:      "1",
	dayOfMonth: "1",
	dayOfWeek:  "1",
	command:    "/usr",
}

func Test_ParseStringArray(t *testing.T) {
	testSuite := []struct {
		name     string
		strArray []string
		expected *body
		err      error
	}{
		{"length of expression is greater than 6", []string{"*", "*", "*", "*", "*", "*", "*", "/usr"}, nil, errors.New("invalid expression")},
		{"missing any field", []string{"*", "*", "/usr"}, nil, errors.New("invalid expression")},
		{"missing any field", []string{"1", "1", "1", "1", "1", "/usr"}, mockBodyResponse, nil},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got, err := ParseStringArray(data.strArray)
			if !reflect.DeepEqual(got, data.expected) || !reflect.DeepEqual(err, data.err) {
				t.Errorf("got = %v; want = %v", got, data.expected)
				t.Errorf("got = %v; want = %v", err, data.err)
			}
		})
	}
}
