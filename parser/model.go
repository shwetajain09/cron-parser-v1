package parser

var allowedEnglishMonthNames = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
var allowedEnglishWeekNames = []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}

type body struct {
	minutes    string
	hour       string
	dayOfMonth string
	month      string
	dayOfWeek  string
	command    string
}

type Validate func(int) bool

type response struct {
	minutes    []int
	hour       []int
	dayOfMonth []int
	month      []int
	dayOfWeek  []int
	command    string
}
