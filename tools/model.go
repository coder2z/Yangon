package tools

import "regexp"

func SqlType2StructType(t string) (string, bool) {
	switch {
	case regexp.MustCompile(`date`).MatchString(t):
		return "time.Time", true
	case regexp.MustCompile(`var`).MatchString(t):
		return "string", false
	case regexp.MustCompile(`int`).MatchString(t):
		return "uint", false
	default:
		return "string", false
	}
}
