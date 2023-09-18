package parser

import (
	"errors"
	"reflect"
	"testing"
)

func Test_allValues(t *testing.T) {
	testSuite := []struct {
		name     string
		min      int
		max      int
		expected []int
	}{
		{"Min and Max are different values", 1, 2, []int{1, 2}},
		{"Min and Max are same values ", 1, 1, []int{1}},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got := allValues(data.min, data.max)
			if !reflect.DeepEqual(got, data.expected) {
				t.Errorf("allValues(data.min,data.max) = %d; want %d", got, data.expected)
			}
		})
	}
}

func Test_addAll(t *testing.T) {
	testSuite := []struct {
		name     string
		expField string
		expected []int
	}{
		{"Minute value", "minute", allValues(0, MinuteLimit)},
		{"Hour value", "hour", allValues(0, 23)},
		{"day of month value", "day of month", allValues(1, 31)},
		{"month value", "month", allValues(1, 12)},
		{"day of month value", "day of week", allValues(0, 6)},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got := addAll(data.expField)
			if !reflect.DeepEqual(got, data.expected) {
				t.Errorf("addAll(data.expField) = %d; want %d", got, data.expected)
			}
		})
	}
}

func Test_prepare_minute(t *testing.T) {
	testSuite := []struct {
		name     string
		str      string
		expField string
		fn       Validate
		expected []int
		err      error
	}{
		{"asterisk", "*", "minute", isValidMinute, allValues(0, MinuteLimit), nil},
		{"empty field", "", "minute", isValidMinute, []int{}, errors.New("invalid minute value")},
		{"step value", "2-8/2", "minute", isValidMinute, []int{2, 4, 6, 8}, nil},
		{"continuation", "2-5", "minute", isValidMinute, []int{2, 3, 4, 5}, nil},
		{"incompatible minute value", "u", "minute", isValidMinute, []int{}, errors.New("invalid minute value")},
		{"out of bounds minute value", "90", "minute", isValidMinute, []int{}, errors.New("invalid minute value")},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got, err := prepare(data.str, data.expField, data.fn, MinuteLimit)
			if !reflect.DeepEqual(got, data.expected) || !reflect.DeepEqual(err, data.err) {
				t.Errorf("got = %d; want = %d", got, data.expected)
				t.Errorf("got = %d; want = %d", err, data.err)
			}
		})
	}
}

func Test_prepare_hour(t *testing.T) {
	testSuite := []struct {
		name     string
		str      string
		expField string
		fn       Validate
		max      int
		expected []int
		err      error
	}{
		{"asterisk", "*", "hour", isValidHour, 23, allValues(0, 23), nil},
		{"empty field", "", "hour", isValidHour, 23, []int{}, errors.New("invalid hour value")},
		{"step value", "2-8/2", "hour", isValidHour, 23, []int{2, 4, 6, 8}, nil},
		{"continuation", "2-5", "hour", isValidHour, 23, []int{2, 3, 4, 5}, nil},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got, err := prepare(data.str, data.expField, data.fn, data.max)
			if !reflect.DeepEqual(got, data.expected) || !reflect.DeepEqual(err, data.err) {
				t.Errorf("got = %d; want = %d", got, data.expected)
				t.Errorf("got = %d; want = %d", err, data.err)
			}
		})
	}
}

func Test_prepare_dom(t *testing.T) {
	testSuite := []struct {
		name     string
		str      string
		expField string
		fn       Validate
		max      int
		expected []int
		err      error
	}{
		{"asterisk", "*", "day of month", isValidDOM, 31, allValues(1, 31), nil},
		{"empty field", "", "day of month", isValidDOM, 31, []int{}, errors.New("invalid day of month value")},
		{"step value", "2-8/2", "day of month", isValidDOM, 31, []int{2, 4, 6, 8}, nil},
		{"continuation", "2-5", "day of month", isValidDOM, 31, []int{2, 3, 4, 5}, nil},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got, err := prepare(data.str, data.expField, data.fn, data.max)
			if !reflect.DeepEqual(got, data.expected) || !reflect.DeepEqual(err, data.err) {
				t.Errorf("got = %d; want = %d", got, data.expected)
				t.Errorf("got = %d; want = %d", err, data.err)
			}
		})
	}
}

func Test_prepare_month(t *testing.T) {
	testSuite := []struct {
		name     string
		str      string
		expField string
		fn       Validate
		max      int
		expected []int
		err      error
	}{
		{"asterisk", "*", "month", isValidMonth, 12, allValues(1, 12), nil},
		{"empty field", "", "month", isValidMonth, 12, []int{}, errors.New("invalid month value")},
		{"step value", "2-8/2", "month", isValidMonth, 12, []int{2, 4, 6, 8}, nil},
		{"continuation", "Feb-May", "month", isValidMonth, 12, []int{2, 3, 4, 5}, nil},
		{"english word", "jAn", "month", isValidMonth, 12, []int{1}, nil},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got, err := prepare(data.str, data.expField, data.fn, data.max)
			if !reflect.DeepEqual(got, data.expected) || !reflect.DeepEqual(err, data.err) {
				t.Errorf("got = %d; want = %d", got, data.expected)
				t.Errorf("got = %d; want = %d", err, data.err)
			}
		})
	}
}

func Test_prepare_dow(t *testing.T) {
	testSuite := []struct {
		name     string
		str      string
		expField string
		fn       Validate
		max      int
		expected []int
		err      error
	}{
		{"asterisk", "*", "day of week", isValidDOW, 7, allValues(0, 6), nil},
		{"empty field", "", "day of week", isValidDOW, 7, []int{}, errors.New("invalid day of week value")},
		{"step value", "2-6/2", "day of week", isValidDOW, 7, []int{2, 4, 6}, nil},
		{"continuation", "tue-Fri", "day of week", isValidDOW, 7, []int{2, 3, 4, 5}, nil},
		{"english word", "mOn", "day of week", isValidDOW, 7, []int{1}, nil},
	}

	for _, data := range testSuite {
		t.Run(data.name, func(t *testing.T) {
			got, err := prepare(data.str, data.expField, data.fn, data.max)
			if !reflect.DeepEqual(got, data.expected) || !reflect.DeepEqual(err, data.err) {
				t.Errorf("got = %d; want = %d", got, data.expected)
				t.Errorf("got = %d; want = %d", err, data.err)
			}
		})
	}
}
