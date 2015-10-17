package check

// nがmで割り切れるかどうか
func Check(n *uint64, m *uint64) bool {
	return *n % *m == 0
}

// nがa(スライス)で割り切れるかどうか
func CheckAll(n *uint64, a *[]uint64) bool {
	arr := *a
	ok := false
	for _, i := range arr{
		//fmt.Println(i)
		if Check(n, &i){
			ok = true
			break
		}
	}
	return ok
}
