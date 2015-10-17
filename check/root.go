package check

import(
	"math"
)

// nがa(スライス)で割り切れるかどうか(ルート利用)
func RootCheckAll(n *uint64, a *[]uint64) bool {
	arr := *a
	ok := false
	rootN := uint64(math.Sqrt(float64(*n)))
	for _, i := range arr{
		if rootN < i {
			break
		}
		//fmt.Println(i)
		if Check(n, &i){
			ok = true
			break
		}
	}
	return ok
}
