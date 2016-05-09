package main

import (
	"github.com/S--Minecraft/GoPrimeNumber/algo"
	"github.com/S--Minecraft/GoPrimeNumber/util"
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

	//arr := algo.AllToNum(&checkNum)
	//arr := algo.AllUpTo(&checkNum)
	arr := algo.ReadyMemRootUpTo(&checkNum)

	arr_s := util.Uint64SliceToString(arr)
	fmt.Println("素数: " + arr_s)
	fmt.Printf("計: %d個\n", len(arr))

	//処理時間終了
	timerEnd := time.Now()
	fmt.Printf("%fs\n", (timerEnd.Sub(timerStart)).Seconds())
	/*
		アルゴリズムなし結果
		1000000まで 80.591192s
		100000まで   1.035948s
		50000まで    0.268564s
		30000まで    0.100462s
		10000まで    0.026150s
		1000まで     0.001996s
		500まで      0.002002s
		100まで      0.001002s
		10まで       0.001001s

		Root結果
		1000000まで 54.633719s
		100000まで   0.597540s
		50000まで    0.121492s
		30000まで    0.057508s
		10000まで    0.014011s
		1000まで     0.001009s
		500まで      0.001002s
		100まで      0.001011s
		10まで       0.001001s

		MemRoot結果
		1000000まで 52.867165s
		100000まで   0.600758s
		50000まで    0.137934s
		30000まで    0.060644s
		10000まで    0.014040s
		1000まで     0.002002s
		500まで      0.001001s
		100まで      0.001000s
		10まで       0.001001s

		ReadyMemRoot結果
		1000000まで 30.654910s
		100000まで   0.585412s
		50000まで    0.153005s
		30000まで    0.062100s
		10000まで    0.015233s
		1000まで     0.002002s
		500まで      0.001001s
		100まで      0.001000s
		10まで       0.001001s
	*/
}
