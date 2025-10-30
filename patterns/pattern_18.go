package patterns

import "fmt"

func Pattern18(n int) {
	for i := 1; i <= n; i++ {
		var ch rune = 'A'
		ch = ch + int32(n-i)
		for j := 1; j <= i; j++ {
			fmt.Printf("%c", ch)
			ch++
		}
		fmt.Println()
	}
}
