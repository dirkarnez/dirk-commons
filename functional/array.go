package functional

func Filter(vs []interface{}, f func(interface{}) bool) []interface{} {
	var vsf []interface{} = nil
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}