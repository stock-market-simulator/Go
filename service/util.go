package service

import "strconv"

func convert(year int, month int, day int, separator string) string {
	if day < 1 {
		month -= 1
		day = 30
	}
	if month < 1 {
		year -= 1
		month = 12 - month
	}
	return toString(year) + separator + toString(month) + separator + toString(day)
}

func toString(num int) string {
	if num < 10 {
		return "0" + strconv.Itoa(num)
	} else {
		return strconv.Itoa(num)
	}
}
