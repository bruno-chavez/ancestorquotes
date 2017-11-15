//Package slices has functions that return slices populated with the number of seconds in a minute or number of minutes in a day for example.
package slices

import (
	"strconv"
)

//SecondsMinutes populates with the number of seconds in a minute and the number of minutes in an hour.
func SecondsMinutes() []string {
	slice := make([]string, 60)

	for i := 1; i < 60; i++ {
		slice[i] = strconv.Itoa(i)
	}

	return slice
}
