package part_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	var arr []int
	var secondArrCount map[int]int = map[int]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var input []string = strings.Split(scanner.Text(), "   ")
		firstVal, _ := strconv.Atoi(input[0])
		secondVal, _ := strconv.Atoi(input[1])
		arr = append(arr, firstVal)
		secondArrCount[secondVal]++
	}
	var ans int = 0
	for i := 0; i < len(arr); i++ {
		ans += arr[i] * secondArrCount[arr[i]]
	}
	fmt.Println(ans)
}
