package algo

import(
	"../check"
)

// nまでのすべての素数を返す(2乗を記憶して計算量を減らす&メモリを最初に設定)
func ReadyMemRootUpTo(n *uint64) []uint64 {
	arr := make([]uint64, 0, *n)
	var num uint64 = 2
	square := make([]Square, 0, *n)
	for ; num <= *n; num++ {
		if len(arr) == 0 {
			arrAppend(&arr, &num, &square)
			continue
		}
		if num <= square[0].i2 {
			smalledArr := arr[:square[0].i]
			if !check.CheckAll(&num, &smalledArr) {
				arrAppend(&arr, &num, &square)
				continue
			}
		} else {
			square = square[1:len(square)]
			if num <= square[0].i2 {
				smalledArr := arr[:square[0].i]
				if !check.CheckAll(&num, &smalledArr) {
					arrAppend(&arr, &num, &square)
					continue
				}
			}
		}
	}
	return arr
}

// n個のすべての素数を返す
func ReadyMemRootToNum(n *uint64) []uint64 {
	arr := make([]uint64, 0, *n)
	var num uint64 = 2
	square := make([]Square, 0, *n)
	for ; uint64(len(arr)) < *n; num++ {
		if len(arr) == 0 {
			arrAppend(&arr, &num, &square)
			continue
		}
		if num <= square[0].i2 {
			smalledArr := arr[:square[0].i]
			if !check.CheckAll(&num, &smalledArr) {
				arrAppend(&arr, &num, &square)
				continue
			}
		} else {
			square = square[1:len(square)]
			if num <= square[0].i2 {
				smalledArr := arr[:square[0].i]
				if !check.CheckAll(&num, &smalledArr) {
					arrAppend(&arr, &num, &square)
					continue
				}
			}
		}
	}
	return arr
}
