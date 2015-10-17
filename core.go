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
	ok := false
	for _, i := range arr{
		fmt.Println(i)
		if check(n, &i){
			ok = true
			break
		}
	}
	return ok
}

func main() {
	// 調べる数(uint)
	checkNum, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("ParseError: 引数のパースに失敗しました")
		fmt.Println(err)
		fmt.Println()
		checkNum = 13
	}
	fmt.Println("引数: " + strconv.FormatUint(checkNum, 10) + "\n")

	arr := []uint64{}
	var num uint64 = 2
	for ; num <= checkNum; num++ {
		if len(arr) == 0 {
			arr = append(arr, num)
			continue
		}
		if !checkAll(&num, &arr) {
			arr = append(arr, num)
			continue
		}
	}

	arr_s := ""
	for i, v := range arr {
		arr_s += strconv.FormatUint(v, 10)
		if i != len(arr)-1 {
			arr_s += ", "
		}
	}
	fmt.Println("素数: " + arr_s)
	fmt.Print("計: ")
	fmt.Print(len(arr))
	fmt.Println("個")
}


