package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sortInts(a)

	l, r := 0, n-1
	for l < r {
		sum := a[l] + a[r]
		if sum == x {
			fmt.Println("POSSIBLE")
			return
		}
		if sum < x {
			l++
		} else {
			r--
		}
	}

	fmt.Println("IMPOSSIBLE")
}

func sortInts(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}