package main

func main() {
	var stack []string

	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = append(stack, "c")

	for len(stack) > 0 {
		n := len(stack) - 1
		println(stack[n])

		stack = stack[:n]
	}
}
