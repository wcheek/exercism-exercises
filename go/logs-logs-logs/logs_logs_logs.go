package logs

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	logType := "default"
	for _, char := range log {
		switch fmt.Sprintf("%U", char) {
		case "U+2757":
			logType = "recommendation"
		case "U+1F50D":
			logType = "search"
		case "U+2600":
			logType = "weather"
		}
		if logType != "default" {
			break
		}
	}
	return logType
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var buffer bytes.Buffer
	for _, char := range log {
		if char == oldRune {
			buffer.WriteRune(newRune)
		} else {
			buffer.WriteRune(char)
		}
	}
	return buffer.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
