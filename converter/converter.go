//Receives an interval of time and returns a string with a correct format for the robfig/cron functions.
package converter

import "strconv"

//time is needed to correctly format the string either for seconds or minutes.
//if s is 1 for example, ancestorquotes will run every 1 second or minute depending of time.
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
