package part_1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var input string = ""

func Solve() {
	regexp := regexp.MustCompile(`(?m)mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)`)
	var matches []string = regexp.FindAllString(input, -1)
	var ans int = 0
	for i := 0; i < len(matches); i++ {
		var commaIndex int = strings.Index(matches[i], ",")
		var lastBraceIndex int = strings.LastIndex(matches[i], ")")
		var firstVal, _ = strconv.Atoi(matches[i][4:commaIndex])
		var secondVal, _ = strconv.Atoi(matches[i][commaIndex+1 : lastBraceIndex])
		ans += firstVal * secondVal
	}
	fmt.Println(ans)
}
