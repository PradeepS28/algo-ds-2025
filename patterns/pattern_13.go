package patterns

import "fmt"

func Pattern13(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num)
			fmt.Print(" ")
			num++
		}
		fmt.Println()
	}
}
