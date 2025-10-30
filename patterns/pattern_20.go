package patterns

import "fmt"

/*
Given an integer n. You need to recreate the pattern given below for any value of N. Let's say for N = 5, the pattern should look like as below:

*        * i=1, 8 space
**      ** i=2. 6 space
***    *** i=3, 4 space
****  **** i=4, 2 space
********** i=5, 0 space
****  ****
***    ***
**      **
*        *

Print the pattern in the function given to you.
*/
func Pattern20(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= 2*(n-i); j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print("*")
		}

		for j := 1; j < (i*2)+1; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
