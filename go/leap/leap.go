package leap

// IsLeapYear is the year given a leap year
func IsLeapYear(year int) bool {
	// Not sure of idiomatic go - is it a one line conditional
	// or nested statements?

//  if year%4 == 0 && (year%400 == 0 || year%100 != 0) {
//    return true
//  }

	if year%4 == 0 {
		if year%400 == 0 || year%100 != 0 {
			return true
		}
	}

	return false
}
