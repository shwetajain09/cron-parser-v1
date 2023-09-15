package parser

func isValidMinute(minute int) bool {
	return (minute >= 0 && minute < 60)
}
func isValidHour(hour int) bool {
	return hour >= 0 && hour < 24
}
func isValidDOM(dom int) bool {
	return dom >= 1 && dom < 32
}
func isValidDOW(dow int) bool {
	return dow >= 0 && dow < 8
}
func isValidMonth(month int) bool {
	return month >= 1 && month <= 12
}
