package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println()
		return
	}

	ags := os.Args[1:]

	for _, n := range ags {
		for _,rn := range n {
			if rn < '0' && rn > '9' {
				fmt.Println("Error")
				return
			}

		}
	}

	sliceNum := []int{}

	for _, nb := range ags {
		num, err := strconv.Atoi(nb)
		if err != nil {
			fmt.Println("Error converting string to number")
			return
		}
		sliceNum = append(sliceNum, num)
	}
	maxNum := getMax(sliceNum)
	smallNum := getSmall(sliceNum)

	fmt.Printf("Min:%3v    Max:%3v\n", smallNum, maxNum)
}

// Function getMax() returns the max number
func getMax(s []int) int {
	m := s[0]
	for i, n := range s {
		if n > m {
			m = s[i]
		}
	}

	return m
}

// Function getSmall() returns the smallest number
func getSmall(s []int) int {
	m := s[0]
	for i, n := range s {
		if n < m {
			m = s[i]
		}
	}

	return m
}
