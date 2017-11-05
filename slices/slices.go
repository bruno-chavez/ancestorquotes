package slices

import "strconv"

func SecondsMinutes() []string {
	slice := make([]string, 1440)

	for i := 0;i<1440;i++{
		slice[i] = strconv.Itoa(i)
	}

	return slice
}