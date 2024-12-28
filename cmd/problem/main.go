package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 1}
	m := make(map[int]string)

	for _, val := range nums {
		if _, exists := m[val]; exists {
			fmt.Printf("\n num: %v, exists: %v", val, exists)
		}
		m[val] = ""
	}
}
