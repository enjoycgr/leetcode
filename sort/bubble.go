package main

import "fmt"

func bubble(a []int) {
	if len(a) <= 1 {
		return
	}

	for i := 0; i < len(a); i++ {
		flag := false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}

func main() {
	a := []int{5, 9, 1, 4, 3}
	bubble(a)
	fmt.Println(a)
}
