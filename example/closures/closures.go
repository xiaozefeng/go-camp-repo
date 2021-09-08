package closures


func intSeq() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
