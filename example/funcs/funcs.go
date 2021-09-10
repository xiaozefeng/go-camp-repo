package funcs

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Contains(vs []string, t string) bool {
	for _, v := range vs {
		if v == t {
			return true
		}
	}
	return false
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func Filter(vs []string, f func(string) bool) []string {
	var res []string
	for _, v := range vs {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Map(vs[]string ,f func(string) string) []string {
	var res []string
	for _, v := range vs {
		converted :=f(v)
		res = append(res, converted)
	}
	return res
}


