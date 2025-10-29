package patterns

import "fmt"

func Pattern15(n int) {
	for i := n; i > 0; i-- {
		var ch rune = 'A'
		for j := i; j > 0; j-- {
			fmt.Printf("%c", ch)
			ch++
		}
		fmt.Println()
	}
}
