package parser

import (
	"testing"
)

var (
	validMinute         = 45
	lessThanZero        = -2
	greaterThanFiveNine = 89

	validHour           = 8
	greaterThanTwoThree = 24

	validDom            = 2
	lessThanOne         = 0
	greaterThanThreeOne = 32

	validMonth        = 2
	greaterThanTwelve = 13

	validWeek        = 2
	greaterThanSeven = 8
)

func Test_isValidMinute(t *testing.T) {
	testSuite := []struct {
		name     string
		minute   int
		expected bool
	}{
		{"Valid Minute", validMinute, true},
		{"Less than zero minute value ", lessThanZero, false},
		{"More than five nine minute value", greaterThanFiveNine, false},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got := isValidMinute(data.minute)
			if got != data.expected {
				t.Errorf("isValidMinute(data.minute) = %t; want %t", got, data.expected)
			}
		})
	}
}

func Test_isValidHour(t *testing.T) {
	testSuite := []struct {
		name     string
		hour     int
		expected bool
	}{
		{"Valid Hour", validHour, true},
		{"Less than zero hour value", lessThanZero, false},
		{"More than two three hour value", greaterThanTwoThree, false},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got := isValidHour(data.hour)
			if got != data.expected {
				t.Errorf("isValidHour(data.hour) = %t; want %t", got, data.expected)
			}
		})
	}
}

func Test_isValidDom(t *testing.T) {
	testSuite := []struct {
		name     string
		dom      int
		expected bool
	}{
		{"Valid day of month", validDom, true},
		{"Less than one for dom value", lessThanOne, false},
		{"More than three one for dom value", greaterThanThreeOne, false},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got := isValidDOM(data.dom)
			if got != data.expected {
				t.Errorf("isValidDom(data.dom) = %t; want %t", got, data.expected)
			}
		})
	}
}

func Test_isValidMonth(t *testing.T) {
	testSuite := []struct {
		name     string
		month    int
		expected bool
	}{
		{"Valid Hour", validMonth, true},
		{"Less than one for month value", lessThanOne, false},
		{"More than twelve for month value", greaterThanTwelve, false},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got := isValidMonth(data.month)
			if got != data.expected {
				t.Errorf("isValidMonth(data.hour) = %t; want %t", got, data.expected)
			}
		})
	}
}

func Test_isValidDow(t *testing.T) {
	testSuite := []struct {
		name     string
		dow      int
		expected bool
	}{
		{"Valid Day of week", validWeek, true},
		{"Less than zero for week value", lessThanZero, false},
		{"More than seven for week value", greaterThanSeven, false},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got := isValidDOW(data.dow)
			if got != data.expected {
				t.Errorf("isValidDOW(data.dow) = %t; want %t", got, data.expected)
			}
		})
	}
}
