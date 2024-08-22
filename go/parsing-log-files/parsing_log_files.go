package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	validityTest, err := regexp.Compile(`^\[((TRC)|(DBG)|(INF)|(WRN)|(ERR)|(FTL))\]`)
	if err != nil {
		panic("regex failed to compile")
	}
	return validityTest.MatchString(text)
}

func SplitLogLine(text string) []string {
	splitRegexp, err := regexp.Compile(`<(\*|\~|\=|\-)*>`)
	if err != nil {
		panic("regex failed to compile")
	}
	return splitRegexp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	passwordRegexp, err := regexp.Compile("(?i)" + `".*password.*"`)
	if err != nil {
		panic("regex failed to compile")
	}
	totalCount := 0
	for _, line := range lines {
		if passwordRegexp.MatchString(line) {
			totalCount++
		}
	}
	return totalCount
}

func RemoveEndOfLineText(text string) string {
	endOfLineRegexp, err := regexp.Compile(`end-of-line\d*`)
	if err != nil {
		panic("regex failed to compile")
	}
	return endOfLineRegexp.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	tagRegexp, err := regexp.Compile(`User\s+(\w*)`)
	if err != nil {
		panic("regex failed to compile")
	}
	linesWithTags := make([]string, len(lines))
	for ind, line := range lines {
		if tagRegexp.MatchString(line) {
			userName := tagRegexp.FindStringSubmatch(line)[1]
			linesWithTags[ind] = fmt.Sprintf("[USR] %s %s", userName, line)
		} else {
			linesWithTags[ind] = line
		}
	}
	return linesWithTags
}
