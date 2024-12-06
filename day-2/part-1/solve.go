package part_1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkArray(arr []int) bool {
	var copyArr []int = make([]int, len(arr))
	copy(copyArr, arr)
	slices.Sort(copyArr)
	if slices.Equal(arr, copyArr) {
		for i := 1; i < len(arr); i++ {
			if arr[i]-arr[i-1] < 1 || arr[i]-arr[i-1] > 3 {
				return false
			}
		}
		return true
	}
	slices.Reverse(copyArr)
	if slices.Equal(arr, copyArr) {
		for i := 1; i < len(arr); i++ {
			if arr[i-1]-arr[i] < 1 || arr[i-1]-arr[i] > 3 {
				return false
			}
		}
		return true
	}
	return false
}

func Solve() {
	var arr []int
	scanner := bufio.NewScanner(os.Stdin)
	var ans int = 0
	for scanner.Scan() {
		var input []string = strings.Split(scanner.Text(), " ")
		arr = []int{}
		for i := 0; i < len(input); i++ {
			val, _ := strconv.Atoi(input[i])
			arr = append(arr, val)
		}
		if checkArray(arr) {
			ans++
		}
	}
	fmt.Println(ans)
}
