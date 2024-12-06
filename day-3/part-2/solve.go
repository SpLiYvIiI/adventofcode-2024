package part_2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var input string = ""

func Solve() {
	regexp := regexp.MustCompile(`(?m)mul\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\)|don't\(\)|do\(\)`)
	var matches []string = regexp.FindAllString(input, -1)
	var ans int = 0
	var wasLastDo bool = true
	for i := 0; i < len(matches); i++ {
		if matches[i] == "don't()" {
			wasLastDo = false
			continue
		}
		if matches[i] == "do()" {
			wasLastDo = true
			continue
		}
		var commaIndex int = strings.Index(matches[i], ",")
		var lastBraceIndex int = strings.LastIndex(matches[i], ")")
		var firstVal, _ = strconv.Atoi(matches[i][4:commaIndex])
		var secondVal, _ = strconv.Atoi(matches[i][commaIndex+1 : lastBraceIndex])
		if wasLastDo {
			ans += firstVal * secondVal
		}
	}
	fmt.Println(ans)
}
