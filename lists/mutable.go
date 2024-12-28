package lists

import (
	"math/rand"
	"sort"
)

// one
func InsertFirst[A any](slice *[]A, element A) {
	*slice = append([]A{element}, *slice...)
}

func InsertLast[A any](slice *[]A, element A) {
	*slice = append(*slice, element)
}

func InsertAt[A any](slice *[]A, index int, element A) {
	*slice = append((*slice)[:index+1], (*slice)[index:]...)
	(*slice)[index] = element
}

// all
func InsertAllFirst[A any](slice *[]A, elements []A) {
	*slice = append(elements, *slice...)
}

func InsertAll[A any](slice *[]A, elements []A) {
	*slice = append(*slice, elements...)
}

func InsertAllLast[A any](slice *[]A, elements []A) {
	InsertAll(slice, elements)
}

func InsertAllAt[A any](slice *[]A, index int, elements []A) {
	*slice = append((*slice)[:index], append(elements, (*slice)[index:]...)...)
}

// remove
func RemoveFirst[A any](slice *[]A) {
	*slice = (*slice)[1:]
}

func RemoveAt[A any](slice *[]A, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func RemoveLast[A any](slice *[]A) {
	*slice = (*slice)[:len(*slice)-1]
}

func Toggle[A comparable](slice *[]A, element A) {
	indexToRemove := IndexOf(*slice, element)
	if indexToRemove == -1 {
		InsertLast(slice, element)
	} else {
		RemoveAt(slice, indexToRemove)
	}
}

// range
func FilterOutRange[A any](slice *[]A, fromIndex int, toIndex int) {
	*slice = append((*slice)[:fromIndex], (*slice)[toIndex+1:]...)
}

func FilterInRange[A any](slice *[]A, fromIndex int, toIndex int) {
	*slice = (*slice)[fromIndex:toIndex]
}

// filter
func Filter[A any](slice *[]A, cb func(A) bool) {
	newSlice := []A{}
	for _, item := range *slice {
		if cb(item) {
			newSlice = append(newSlice, item)
		}
	}
	*slice = newSlice
}

func FilterIndexed[A any](slice *[]A, cb func(A, int) bool) {
	newSlice := []A{}
	for i, item := range *slice {
		if cb(item, i) {
			newSlice = append(newSlice, item)
		}
	}
	*slice = newSlice
}

// reverse
func Reverse[A any](slice []A) {
	left, right := 0, len(slice)-1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}

// times
func Times[A any](count int, cb func(int) A) []A {
	newSlice := make([]A, count)
	for i := 0; i < count; i++ {
		newSlice[i] = cb(i)
	}
	return newSlice
}

// uniquify
func Uniquify[A comparable](slice *[]A) {
	for i := 0; i < len(*slice); i++ {
		for j := 0; j < i; j++ {
			if (*slice)[i] == (*slice)[j] {
				RemoveAt(slice, i)
				i--
				break
			}
		}
	}
}

// sort
func Sort[A any](slice []A, cb func(A, A) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return cb(slice[i], slice[j])
	})
}

// shuffle
func Shuffle[A any](slice []A) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ShuffleRNG[A any](slice []A, random rand.Rand) {
	for i := range slice {
		j := random.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Pick[A any](slice []A) A {
	return slice[rand.Intn(len(slice))]
}

func PickRNG[A any](slice []A, random rand.Rand) A {
	return slice[random.Intn(len(slice))]
}

// clear
func Clear[A any](slice *[]A) {
	*slice = []A{}
}

// synonyms
func Push[A any](slice *[]A, element A) {
	InsertLast(slice, element)
}

func Pop[A any](slice *[]A, element A) {
	RemoveLast(slice)
}
