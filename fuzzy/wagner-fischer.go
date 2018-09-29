package fuzzy

func min(values []int) int {
	min := values[0]
	for _, val := range values {
		if val < min {
			min = val
		}
	}
	return min
}

func make2DArray(m int, n int) [][]int {
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	return arr
}

// func display2DArr(arr [][]int) {
// 	for _, row := range arr {
// 		fmt.Printf("%v\n", row)
// 	}
// }

func EditDistance(s string, t string) int {
	arr := make2DArray(len(s)+1, len(t)+1)
	for i := range s {
		arr[i+1][0] = i + 1
	}

	for j := range t {
		arr[0][j+1] = j + 1
	}

	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {

			if s[i-1] == t[j-1] {
				arr[i][j] = arr[i-1][j-1]
			} else {
				arr[i][j] = min([]int{arr[i-1][j] + 1,
					arr[i][j-1] + 1,
					arr[i-1][j-1] + 1})
			}
		}
	}

	return arr[len(s)][len(t)]
}
