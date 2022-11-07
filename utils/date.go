package utils

const (
	ValidMonthAndDay = iota
	InvalidMonth
	InvalidDay
)

func ValidateMonthAndDay(month int, day int) int {
	if month < 1 || month > 12 {
		return InvalidMonth
	}

	if day < 1 || day > 31 {
		return InvalidDay
	}

	if month == 2 {
		if day >= 1 && day <= 29 {
			return ValidMonthAndDay
		}
	}
	if isThirtyDayMonth(month) {
		if day >= 1 && day <= 30 {
			return ValidMonthAndDay
		}
	}
	if isThirtyOneDayMonth(month) {
		if day >= 1 && day <= 31 {
			return ValidMonthAndDay
		}
	}
	return InvalidDay
}

func isThirtyOneDayMonth(month int) bool {
	return month == 1 ||
		month == 3 ||
		month == 5 ||
		month == 7 ||
		month == 8 ||
		month == 10 ||
		month == 12
}

func isThirtyDayMonth(month int) bool {
	return month == 4 ||
		month == 6 ||
		month == 9 ||
		month == 11
}
