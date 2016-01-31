package algo

import(
	"github.com/S--Minecraft/GoPrimeNumber/check"
)

type Square struct{
	i uint64
	i2 uint64  //i番目の素数の2乗
}

// 素数の配列に追加するとともに2乗の配列にも追加する
func arrAppend(arr *[]uint64, num *uint64, square *[]Square) {
	*arr = append(*arr, *num)
	a := *arr
	last := a[len(a)-1]
	*square = append(*square, Square{uint64(len(a)-1), last*last})
}

// nまでのすべての素数を返す(2乗を記憶して計算量を減らす)
func MemRootUpTo(n *uint64) []uint64 {
	arr := []uint64{}
	var num uint64 = 2
	square := []Square{}
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
func MemRootToNum(n *uint64) []uint64 {
	arr := []uint64{}
	var num uint64 = 2
	square := []Square{}
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
