package main

func ReturnFibonacci(n int) []int {
	var output []int
	t1 := 0
	t2 := 1

	for i := 1; i <= n; i++ {
		if i == 1 {
			output = append(output, 0)
			continue
		}

		if i == 2 {
			output = append(output, 1)
			continue
		}

		next := t1 + t2
		output = append(output, next)

		t1 = t2
		t2 = next
	}

	return output
}

func main() {}
