package main

import (
	"fmt"
)

// nがmで割り切れるかどうか
func check(n int, m int) bool {
	return n % m == 0
}

// nがa(スライスのポインタ)で割り切れるかどうか
func checkAll(n int, a *[]int) bool {
	arr := *a
	ok := false
	for _, i := range arr{
		fmt.Println(i)
		if check(n, i){
			ok = true
			break
		}
	}
	return ok
}

func main() {
	arr := []int{2, 3, 5, 7, 9}
	fmt.Println(checkAll(8, &arr))
}


