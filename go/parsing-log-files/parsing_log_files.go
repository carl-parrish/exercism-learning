package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[TRC\] |^\[DBG\] |^\[INF\] |^WRN |^\[ERR\] |^\[FTL\]`)
		return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[\*\~\-\=]*>`)
		return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[\w\s]*password[\w\s]*"`)
	count :=0
		for _, line := range lines {
			count += len(re.FindAllString(line, -1))
		}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
		return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)
		for idx, line := range lines {
			match := re.FindStringSubmatch(line)
			if  len(match) > 1  {
				lines[idx] = fmt.Sprintf("[USR] %s %s", match[1], line)
			}else {
				continue
			}
		}
	return lines
}
