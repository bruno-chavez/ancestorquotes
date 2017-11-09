//Package convert receives an interval of time and returns
package converter

import "strconv"

func TimeConverter(s string, time string) string {
	integer, _ := strconv.Atoi(s)

	if time == "minutes" {
		intervalString := strconv.Itoa(integer)
		return "@every 0h" + intervalString + "m00s"

	} else {
		intervalString := strconv.Itoa(integer)
		return "@every 0h00" + intervalString + "s"
	}
}
