package algo

import(
	"../check"
)

// nまでのすべての素数を返す(ルートを使って計算量を減らす)
func RootUpTo(n *uint64) []uint64 {
	arr := []uint64{}
	var num uint64 = 2
	for ; num <= *n; num++ {
		if len(arr) == 0 {
			arr = append(arr, num)
			continue
		}
		if !check.RootCheckAll(&num, &arr) {
			arr = append(arr, num)
			continue
		}
	}
	return arr
}

// n個のすべての素数を返す
func RootToNum(n *uint64) []uint64 {
	arr := []uint64{}
	var num uint64 = 2
	for ; uint64(len(arr)) < *n; num++ {
		if len(arr) == 0 {
			arr = append(arr, num)
			continue
		}
		if !check.RootCheckAll(&num, &arr) {
			arr = append(arr, num)
			continue
		}
	}
	return arr
}
