package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(arrayRandomNumbers *[100000]int) {
	for i := range arrayRandomNumbers {
		for j := range arrayRandomNumbers {
			if arrayRandomNumbers[i] < arrayRandomNumbers[j] {
				arrayRandomNumbers[j], arrayRandomNumbers[i] = arrayRandomNumbers[i], arrayRandomNumbers[j]
			}
		}
	}
}

func stoogesort(list []int) {
	last := len(list) - 1
	if list[last] < list[0] {
		list[0], list[last] = list[last], list[0]
	}
	if last > 1 {
		t := len(list) / 3
		stoogesort(list[:len(list)-t])
		stoogesort(list[t:])
		stoogesort(list[:len(list)-t])
	}
}

func main() {
	fmt.Println(string("App is running!!!"))

	// var randomArray [100000]int
	// rand.Seed(time.Now().Unix())
	// for i := range randomArray {
	// 	randomArray[i] = rand.Intn(1000000)
	// }
	// fmt.Println(randomArray)
	// bubbleSort(&randomArray)
	// fmt.Println(randomArray)

	// anotherRandom := make([]int, 0, 10000)
	// for i := 0; i < 10000; i++ {
	// 	// anotherRandom[i] = rand.Intn(100)
	// 	anotherRandom = append(anotherRandom, rand.Intn(100000))
	// }
	// fmt.Println(anotherRandom)
	// stoogesort(anotherRandom)
	// fmt.Println(anotherRandom)

	// mergeRandom := make([]int, 0, 1000000000)
	// for i := 0; i < 10000; i++ {
	// 	mergeRandom = append(mergeRandom, rand.Intn(10000000000))
	// }
	// // fmt.Println(mergeRandom)
	// mergeRandom = mergeSort(mergeRandom)
	// fmt.Println(mergeRandom)

	quickSortRandom := make([]int, 0, 2500000000)
	for i := 0; i < 10000; i++ {
		quickSortRandom = append(quickSortRandom, rand.Intn(10000000000))
	}
	// fmt.Println(quickSortRandom)
	quickSort(quickSortRandom)
	fmt.Println(quickSortRandom)

	fmt.Println(string("App is done!!!"))
}

//Merge Sort
func mergeSort(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	left := make([]int, 0)
	right := make([]int, 0)
	m := len(a) / 2

	for i, x := range a {
		switch {
		case i < m:
			left = append(left, x)
		case i >= m:
			right = append(right, x)
		}
	}

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {

	results := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				results = append(results, left[0])
				left = left[1:len(left)]
			} else {
				results = append(results, right[0])
				right = right[1:len(right)]
			}
		} else if len(left) > 0 {
			results = append(results, left[0])
			left = left[1:len(left)]
		} else if len(right) > 0 {
			results = append(results, right[0])
			right = right[1:len(right)]
		}
	}

	return results
}

//quick sort
func partition(lo int, piv int, arr []int) int {
	is := lo

	for i := lo; i < piv; i++ {
		if arr[i] < arr[piv] {
			if i != is {
				arr[i], arr[is] = arr[is], arr[i]
			}

			is++
		}
	}

	arr[is], arr[piv] = arr[piv], arr[is]

	if is-1 > lo {
		partition(lo, is-1, arr)
	}
	if is+1 < piv {
		partition(is+1, piv, arr)
	}

	return is
}

func quickSort(arr []int) []int {
	l := len(arr)
	piv := l - 1

	partition(0, piv, arr)

	return arr
}
