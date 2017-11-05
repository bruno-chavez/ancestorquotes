package converter

import "strconv"

func TimeConverter(s string, time string) string {
	intervalInteger, _ := strconv.Atoi(s)

	if intervalInteger < 60 {
		return strconv.Itoa(intervalInteger)
	} else {

	}
}
