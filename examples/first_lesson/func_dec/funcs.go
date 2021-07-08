package main

func main() {
	a := Fun0("Tom")
	println(a)

	b, c := Fun1("a", 17)
	println(b)
	println(c)

	_, d := Fun2("a", "b")
	println(d)
}

// Fun0 只有一个返回值，不需要括号括起来
func Fun0(name string) string {
	return "Hello, " + name
}

// Fun1 多个参数，多个返回值。参数有名字，但是返回值没有
func Fun1(a string, b int) (int, string) {
	return 0, "你好"
}

// Fun2 的返回值具有名字，可以在内部直接复制，然后返回
// 也可以忽略age, name，直接返回别的。
func Fun2(a string, b string) (age int, name string) {
	age = 19
	name = "Tom"
	return
	//return 19, "Tom" // 这样返回也可以
}

// Fun3 多个参数具有相同类型放在一起，可以只写一次类型
func Fun3(a, b, c string, abc, bcd int, p string) (d, e int, g string) {
	d = 15
	e = 16
	g = "你好"
	return
	//return 0, 0, "你好" // 这样也可以
}