package main

import (
	"fmt"
	"os"
	"strconv"
)

// nがmで割り切れるかどうか
func check(n *uint64, m *uint64) bool {
	return *n % *m == 0
}

// nがa(スライス)で割り切れるかどうか
func checkAll(n *uint64, a *[]uint64) bool {
	arr := *a
	ok := true
	for _, i := range arr{
		fmt.Println(i)
		if check(n, &i){
			ok = false
			break
		}
	}
	return ok
}

func main() {
	// 調べる数(uint)
	checkNum, err := strconv.ParseUint(os.Args[0], 10, 64)
	if err != nil {
		fmt.Println("ParseError: 引数のパースに失敗しました")
		fmt.Println(err)
		fmt.Println()
		checkNum = 13
	}
	fmt.Println("引数: " + strconv.FormatUint(checkNum, 10) + "\n")

	arr := []uint64{2, 3, 5, 7, 9}
	fmt.Println(checkAll(&checkNum, &arr))
}


