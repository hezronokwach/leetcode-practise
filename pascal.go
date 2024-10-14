package main

import (
	"fmt"
	"strconv"
)

func pascalTriangle(n int) [][]int {
	// Implementation here
	var result [][]int
	count := 2
	// var last []int
	result = append(result, []int{1})
	if n == 1 {
		return result
	}
	result = append(result, []int{1, 1})
	if n == 2 {
		return result
	}

	for count < n {
		last := result[len(result)-1]
		prev := 0
		var store []int

		for _, num := range last {
			store = append(store, num+prev)
			prev = num
		}
		store = append(store, last[len(last)-1])
		result = append(result, store)
		count += 1

	}
	return result
}

func printTriangle(triangle [][]int) {
	for _, row := range triangle {
		fmt.Printf("[%s]\n", joinInts(row))
	}
}

func joinInts(arr []int) string {
	// Join integers as strings for printing
	var result string
	for _, num := range arr {
		result += strconv.Itoa(num) + ","
	}
	return result[:len(result)-1]
}

func main() {
	printTriangle(pascalTriangle(5))
}
