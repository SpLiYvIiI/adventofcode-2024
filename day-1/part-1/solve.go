package part_1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve() {
	var firstArr []int
	var secondArr []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var input []string = strings.Split(scanner.Text(), "   ")
		firstVal, _ := strconv.Atoi(input[0])
		secondVal, _ := strconv.Atoi(input[1])
		firstArr = append(firstArr, firstVal)
		secondArr = append(secondArr, secondVal)
	}
	slices.Sort(firstArr)
	slices.Sort(secondArr)
	var ans int = 0
	for i := 0; i < len(firstArr); i++ {
		if firstArr[i] > secondArr[i] {
			ans += firstArr[i] - secondArr[i]
		} else {
			ans += secondArr[i] - firstArr[i]
		}
	}
	fmt.Println(ans)
}
