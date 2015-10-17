package util

import(
	"strconv"
)

//uint64のスライスを文字列に変換
func Uint64SliceToString(arr []uint64) string {
	arr_s := ""
	for i, v := range arr {
		arr_s += strconv.FormatUint(v, 10)
		if i != len(arr)-1 {
			arr_s += ", "
		}
	}
	return arr_s
}

