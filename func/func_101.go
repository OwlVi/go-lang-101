package func_101

// func namefunc(parameter) returntype
func Foofoo() int {
	return 0
}
func Foo2(x int, y int) int {
	return x + y
}

func Foo3(x float32, y int) int {
	return int(x) + y
}

func Foo4(x int, y int) (int, int) {
	return x + y, x - y
}
