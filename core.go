package main

import (
	"./algo"
	"./util"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	//処理時間開始
	timerStart := time.Now()

	// 調べる数(uint)
	checkNum, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("ParseError: 引数のパースに失敗しました")
		fmt.Println(err)
		fmt.Println()
		checkNum = 13
	}
	fmt.Println("引数: " + strconv.FormatUint(checkNum, 10) + "\n")

	arr := algo.AllUpTo(&checkNum)

	arr_s := util.Uint64SliceToString(arr)
	fmt.Println("素数: " + arr_s)
	fmt.Print("計: ")
	fmt.Print(len(arr))
	fmt.Println("個")

	//処理時間終了
	timerEnd := time.Now()
	fmt.Printf("%f秒\n", (timerEnd.Sub(timerStart)).Seconds())
	/*
		アルゴリズムなし結果
		100000まで 1.035948s
		50000まで 0.268564s
		30000まで 0.100462s
		10000まで 0.026150s
		1000まで 0.001996s
		500まで 0.002002s
		100まで 0.01002s
		10まで 0.001001s
	*/
}
