package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	appMap := map[rune]string{
		'\u2757': "recommendation",
		'\U0001F50D': "search",
		'\u2600': "weather",

	}
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
 swap := func(currentRune rune)rune {
	switch {
	case currentRune == oldRune:
		return newRune
	default:
		return currentRune
	}
 }
	return strings.Map(swap, log)
	 
}


// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
