package utils

import (
	"time"
)

func ToTimestamp(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05")
}

func JoinStrings(strings []string) string {
	result := ""
	for i, s := range strings {
		if i > 0 {
			result += ", "
		}
		result += s
	}
	return result
}
