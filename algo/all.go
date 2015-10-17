package algo

import(
	"../check"
)

func AllUpTo(n *uint64) []uint64 {
	arr := []uint64{}
	var num uint64 = 2
	for ; num <= *n; num++ {
		if len(arr) == 0 {
			arr = append(arr, num)
			continue
		}
		if !check.CheckAll(&num, &arr) {
			arr = append(arr, num)
			continue
		}
	}
	return arr
}

