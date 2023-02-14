package main

//go:noinline
func foo(i int) *int {
	var f_v1 = 10
	var f_v2 = 11
	var f_v3 = 12
	var f_v4 = 13
	var f_v5 = 14

	//防止内联优化
	for i := 0; i < 5; i++ {
		println(&i, &f_v1, &f_v2, &f_v3, &f_v4, &f_v5)
	}

	return &f_v3
}

//go:noinline
func add(a, b int) int {
	return a + b
}

func main() {
	m_v := foo(8)
	println(*m_v, m_v)

	/* var a, b = 5, 6*/

	//c := add(a, b)
	//println(c)

}
