package patterns

import "fmt"

func Pattern16(n int) {
	var ch rune = 'A'
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
		ch++
	}
}
