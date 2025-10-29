package patterns

import "fmt"

func Pattern14(n int) {
	for i := 1; i <= n; i++ {
		var ch rune = 'A'
		for j := 1; j <= i; j++ {
			fmt.Printf("%c ", ch)
			ch++
		}
		fmt.Println()
	}
}
