//Package slices has functions that return slices populated with the number of seconds in a minute
//or number of minutes in a day for example.
//The syntax for the name of the functions is, first letter "number of x", second letter "in y"
package slices

import "strconv"

func MinutesDay() []string {
	slice := make([]string, 1440)

	for i := 0; i < 1440; i++ {
		slice[i] = strconv.Itoa(i)
	}

	return slice
}
