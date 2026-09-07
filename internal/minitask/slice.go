package minitask

import (
	"fmt"
	"slices"
)

var numSlice = []int{50, 75, 66, 20, 32, 90}

func PrintNum() {
	n1 := slices.Grow(numSlice[0:3], 4)
	n2 := numSlice[3:6]

	n1 = append(n1, 88)

	result := slices.Concat(n1, n2)

	for i := range result {
		fmt.Println(result[i])
	}
}
