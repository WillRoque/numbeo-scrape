package numbeo

import (
	"regexp"
	"strings"
)

var float = regexp.MustCompile(`[0-9]*\.*[0-9]*`)

// GetFloatString cleans str so it can be converted into a float.
func getFloatString(str string) string {
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, " ", "", -1)
	return float.FindString(str)
}
