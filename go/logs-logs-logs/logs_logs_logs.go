package logs

import (
	"strings"
	"unicode/utf8"
)

var appMap = map[rune]string{
		'\u2757': "recommendation",
		'\U0001F50D': "search",
		'\u2600': "weather",

	}

// Application identifies the application emitting the given log.
func Application(log string) string {
	
	for _, currentRune := range log {
		if appName, ok := appMap[currentRune]; ok{
			return appName
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}


// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
