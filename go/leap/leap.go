package leap

// IsLeapYear is the year argument a leap year
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%400 == 0 || year%100 != 0) {
		return true
	}

	return false
}
