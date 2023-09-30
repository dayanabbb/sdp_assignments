package main

import (
	"fmt"
	"sort"
)

type Sorting interface {
	Sort([]int)
}

// bubble sort
type BubbleSort struct{}

func (bs *BubbleSort) Sort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// quick sort
type QuickSort struct{}

func (qs *QuickSort) Sort(array []int) {
	sort.Ints(array)
}

type Context struct {
	strategy Sorting
}

func NewContext(strategy Sorting) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (c *Context) SetStrategy(strategy Sorting) {
	c.strategy = strategy
}

func (c *Context) Sort(array []int) {
	c.strategy.Sort(array)
}

// input array
func main() {
	array := []int{7, 5, 1, 4, 9}

	//output

	context := NewContext(&BubbleSort{})
	fmt.Println("Bubble sort:")
	fmt.Println("Before sorting:", array)
	context.Sort(array)
	fmt.Println("After sorting:", array)

	context.SetStrategy(&QuickSort{})
	fmt.Println("Quick sort:")
	fmt.Println("Before sorting:", array)
	context.Sort(array)
	fmt.Println("After sorting:", array)

}
