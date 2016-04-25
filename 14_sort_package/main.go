package main

import (
	"fmt"
	"sort"
)

type people []string

func (a people) Len() int { return len(a) }

func (a people) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a people) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	// Exercise: sort the fallowing slices

	// (1)
	studyGroup := people{"Zeno", "John", "All", "Jenny"}

	// (2)
	s := []string{"Zeno", "John", "All", "Jenny"}

	// (3)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	// Also sort the above in reverse order

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	fmt.Println("------")

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)

	fmt.Println("------")

	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	fmt.Println("------")
	fmt.Println("------")

	sort.Sort(sort.Reverse(studyGroup))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)
}
