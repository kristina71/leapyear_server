package year

func IsLeapYear(year uint64) bool {
	if year == 0 || year > 9999 {
		return false
	}

	if year%4 != 0 {
		return false
	}

	if year%100 != 0 {
		return true
	}

	if year%400 == 0 {
		return true
	}

	return false
}
