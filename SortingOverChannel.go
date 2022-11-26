package main

import (
	"fmt"
	"sort"
)

func sortList(c chan []int, list []int) {
	sort.Ints(list)
	c <- list
}

func SortingOverChannel() {
	// define the size of the number we want to sort
	// put the number list into a slice
	list := []int{-9, -11, 12, 13, 9}
	fmt.Println("Your digits input : ", list)

	// sort in different go routines
	// c := make(chan []int, 4)
	c := make(chan []int)
	go sortList(c, list)

	sortedS1 := <-c

	fmt.Println(sortedS1)
}
