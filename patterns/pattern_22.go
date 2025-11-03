package patterns

import "fmt"

/*
Given an integer n. You need to recreate the pattern given below for any value of N. Let's say for N = 5, the pattern should look like as below:

5 5 5 5 5 5 5 5 5. i=0, n-1,n,n-1, 7
5 4 4 4 4 4 4 4 5, i=1, n-1,n-1,n-1, 5
5 4 3 3 3 3 3 4 5, i=2, n-1,n-1,n-1, 3
5 4 3 2 2 2 3 4 5, i=3, 1
5 4 3 2 1 2 3 4 5
5 4 3 2 2 2 3 4 5
5 4 3 3 3 3 3 4 5
5 4 4 4 4 4 4 4 5
5 5 5 5 5 5 5 5 5

Print the pattern in the function given to you.
*/
func Pattern22(n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(n-j, " ")
		}
		for j := 0; j < (2*(n-1))-i-2; j++ {
			fmt.Print(n-i, " ")
		}

		for j := i; j >= 0; j-- {
			fmt.Print(n-j, " ")
		}
		fmt.Println()
	}
}
