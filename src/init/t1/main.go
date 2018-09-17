package main

var (
	a string
	b int
)

func init() {
	a = "hello world"
}

func init() {
	b = 110
}

func main() {
	println(a)
	println(b)
}