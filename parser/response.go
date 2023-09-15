package parser

import (
	"errors"
	"strconv"
	"strings"
)

// func BuildResponse(body *body) (*response, error) {
// 	res := &response{
// 		command: body.command,
// 	}
// 	minutes, minutesErr := prepare(body.minutes, "minute", isValidMinute, 59)
// 	hour, hourErr := prepare(body.hour, "hour", isValidHour, 23)
// 	dom, domErr := prepare(body.dayOfMonth, "day of month", isValidDOM, 31)
// 	month, monthErr := prepare(body.month, "month", isValidMonth, 12)
// 	dow, dowErr := prepare(body.dayOfWeek, "day of week", isValidDOW, 7)
// 	err := errors.Join(minutesErr, hourErr, domErr, monthErr, dowErr)
// 	sort.Ints(minutes)
// 	res.minutes = minutes
// 	sort.Ints(hour)
// 	res.hour = hour
// 	sort.Ints(dom)
// 	res.dayOfMonth = dom
// 	sort.Ints(month)
// 	res.month = month
// 	sort.Ints(dow)
// 	res.dayOfWeek = dow

// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func prepareHyphenString(str []string, expField string, fn Validate) ([]int, error) {
	arr := []int{}
	str[0], str[1] = processStrings(expField, str[0], str[1])

	initial, iErr := strconv.Atoi(str[0])
	last, lErr := strconv.Atoi(str[1])
	if iErr != nil || lErr != nil {
		return []int{}, errors.New("invalid " + expField + " value")
	}
	if !fn(initial) || !fn(last) {
		return []int{}, errors.New("invalid " + expField + " value")
	}
	for j := initial; j <= last; j++ {
		arr = append(arr, j)
	}
	return arr, nil
}

func processStrings(expField, str1, str2 string) (string, string) {
	if expField == "month" {
		s0 := formatFieldValues(str1, allowedEnglishMonthNames)
		s1 := formatFieldValues(str2, allowedEnglishMonthNames)
		if s0 != -1 {
			str1 = strconv.Itoa(s0)
		}
		if s1 != -1 {
			str2 = strconv.Itoa(s1)
		}
	}
	if expField == "day of week" {
		s0 := formatFieldValues(str1, allowedEnglishWeekNames)
		s1 := formatFieldValues(str2, allowedEnglishWeekNames)
		if s0 != -1 {
			str1 = strconv.Itoa(s0 - 1)
		}
		if s1 != -1 {
			str2 = strconv.Itoa(s1 - 1)
		}
	}
	return str1, str2
}

func prepareSlashString(str []string, expField string, fn Validate, max int) ([]int, error) {
	arr := []int{}
	str[0], str[1] = processStrings(expField, str[0], str[1])
	// value is *
	start := 0

	subStringsHyphen := strings.Split(str[0], "-")
	if len(subStringsHyphen) == 2 {
		// value is a range
		vals, err := prepareHyphenString(subStringsHyphen, expField, fn)
		if err != nil {
			return []int{}, errors.New("invalid " + expField + " value")
		}
		start = vals[0]
		max = vals[len(vals)-1]
	} else {
		// value is a single numeric digit
		start, _ = strconv.Atoi(str[0])
	}
	// TODO: error check
	interval, _ := strconv.Atoi(str[1])
	if !fn(start) || !fn(max) {
		return []int{}, errors.New("invalid " + expField + " value")
	}
	for start <= max {
		if !fn(start) {
			return []int{}, errors.New("invalid " + expField + " value")
		}
		arr = append(arr, start)
		start = start + interval
	}
	return arr, nil
}

func prepare(str string, expField string, fn Validate, max int) ([]int, error) {
	arr := []int{}
	subStringsComma := strings.Split(str, ",")
	for _, i := range subStringsComma {
		// split the string with a slash to get a step value field value
		subStringsSlash := strings.Split(i, "/")
		if len(subStringsSlash) == 2 {
			val, err := prepareSlashString(subStringsSlash, expField, fn, max)
			if err != nil {
				return []int{}, err
			}
			arr = append(arr, val...)
		} else {
			// split the string with a hyphen to get a continuation field value
			subStringsHyphen := strings.Split(i, "-")
			if len(subStringsHyphen) == 2 {
				val, err := prepareHyphenString(subStringsHyphen, expField, fn)
				if err != nil {
					return []int{}, err
				}
				arr = append(arr, val...)
			} else {
				if expField == "month" {
					e1 := formatFieldValues(i, allowedEnglishMonthNames)
					if e1 != -1 {
						i = strconv.Itoa(e1)
					}
				}
				if expField == "day of week" {
					e1 := formatFieldValues(i, allowedEnglishWeekNames)
					if e1 != -1 {
						i = strconv.Itoa(e1 - 1)
					}
				}
				// if the string is a *, then add possible values
				if i == "*" {
					arr = append(arr, addAll(expField)...)
				} else {
					ele, err := strconv.Atoi(i)
					if err != nil {
						return []int{}, errors.New("invalid " + expField + " value")
					}
					if !fn(ele) {
						return []int{}, errors.New("invalid " + expField + " value")
					}
					arr = append(arr, ele)
				}
			}
		}
	}
	return arr, nil
}

func formatFieldValues(str string, allowedVals []string) int {
	for i, j := range allowedVals {
		if j == strings.ToLower(str) {
			return i + 1
		}
	}
	return -1
}

func addAll(expField string) []int {
	switch expField {
	case "minute":
		return allValues(0, 59)
	case "hour":
		return allValues(0, 23)
	case "day of month":
		return allValues(1, 31)
	case "month":
		return allValues(1, 12)
	case "day of week":
		return allValues(0, 6)
	}
	return []int{}
}

func allValues(min, max int) []int {
	arr := []int{}
	for i := min; i <= max; i++ {
		arr = append(arr, i)
	}
	return arr
}
