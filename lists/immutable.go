package lists

import "math/rand"

func ToFiltered[A any](slice []A, cb func(A) bool) []A {
	newSlice := []A{}
	for _, item := range slice {
		if cb(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

func ToFilteredIndexed[A any](slice []A, cb func(A, int) bool) []A {
	newSlice := []A{}
	for i, item := range slice {
		if cb(item, i) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

// map
func Map[A any, B any](slice []A, cb func(A) B) []B {
	newSlice := make([]B, len(slice))
	for i, item := range slice {
		newSlice[i] = cb(item)
	}
	return newSlice
}

func MapIndexed[A any, B any](slice []A, cb func(A, int) B) []B {
	newSlice := make([]B, len(slice))
	for i, item := range slice {
		newSlice[i] = cb(item, i)
	}
	return newSlice
}

// flat map
func FlatMap[A any, B any](slice []A, cb func(A) []B) []B {
	newSlice := []B{}
	for _, item := range slice {
		newSlice = append(newSlice, cb(item)...)
	}
	return newSlice
}

func FlatMapIndexed[A any, B any](slice []A, cb func(A, int) []B) []B {
	newSlice := []B{}
	for i, item := range slice {
		newSlice = append(newSlice, cb(item, i)...)
	}
	return newSlice
}

// map not null
func MapNotNull[A any, B any](slice []A, cb func(A) *B) []B {
	newSlice := []B{}
	for _, item := range slice {
		val := cb(item)
		if val != nil {
			newSlice = append(newSlice, *val)
		}
	}
	return newSlice
}

func MapNotNullIndexed[A any, B any](slice []A, cb func(A, int) *B) []B {
	newSlice := []B{}
	for i, item := range slice {
		val := cb(item, i)
		if val != nil {
			newSlice = append(newSlice, *val)
		}
	}
	return newSlice
}

// flat map not null
func FlatMapNotNull[A any, B any](slice []A, cb func(A) []*B) []B {
	newSlice := []B{}
	for _, item := range slice {
		res := cb(item)
		list := MapNotNull(res, func(e *B) *B { return e })
		newSlice = append(newSlice, list...)
	}
	return newSlice
}

func FlatMapNotNullIndexed[A any, B any](slice []A, cb func(A, int) []*B) []B {
	newSlice := []B{}
	for i, item := range slice {
		res := cb(item, i)
		list := MapNotNull(res, func(e *B) *B { return e })
		newSlice = append(newSlice, list...)
	}
	return newSlice
}

// reduce / fold
func Reduce[A any](slice []A, cb func(A, A) A) A {
	initialValue := slice[0]
	for _, item := range slice[1:] {
		initialValue = cb(initialValue, item)
	}

	return initialValue
}

func ReduceIndexed[A any](slice []A, cb func(A, A, int) A) A {
	initialValue := slice[0]
	for i, item := range slice[1:] {
		initialValue = cb(initialValue, item, i)
	}

	return initialValue
}

func Fold[A any, B any](slice []A, cb func(B, A) B, initialValue B) B {
	for _, item := range slice {
		initialValue = cb(initialValue, item)
	}

	return initialValue
}

func FoldIndexed[A any, B any](slice []A, cb func(B, A, int) B, initialValue B) B {
	for i, item := range slice {
		initialValue = cb(initialValue, item, i)
	}

	return initialValue
}

// clone
func CloneShallow[A any](slice []A) []A {
	newSlice := make([]A, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// reverse
func ToReversed[A any](slice []A) []A {
	newSlice := CloneShallow(slice)
	Reverse(newSlice)
	return newSlice
}

// reverse
func ToUniquify[A comparable](slice []A) []A {
	newSlice := CloneShallow(slice)
	Uniquify(&newSlice)
	return newSlice
}

// sort
func ToSorted[A any](slice []A, cb func(A, A) bool) []A {
	newSlice := CloneShallow(slice)
	Sort(newSlice, cb)
	return newSlice
}

// shuffle
func ToShuffled[A any](slice []A) []A {
	newSlice := CloneShallow(slice)
	Shuffle(newSlice)
	return newSlice
}

func ToShuffledRNG[A any](slice []A, random rand.Rand) []A {
	newSlice := CloneShallow(slice)
	ShuffleRNG(newSlice, random)
	return newSlice
}

// find
func FindFirst[A any](slice []A, cb func(A) bool) (A, int) {
	for i, item := range slice {
		if cb(item) {
			return item, i
		}
	}
	var zeroValue A
	return zeroValue, -1
}

func FindFirstIndexed[A any](slice []A, cb func(A, int) bool) (A, int) {
	for i, item := range slice {
		if cb(item, i) {
			return item, i
		}
	}
	var zeroValue A
	return zeroValue, -1
}

func FindLast[A any](slice []A, cb func(A) bool) (A, int) {
	for i := len(slice) - 1; i >= 0; i-- {
		if cb(slice[i]) {
			return slice[i], i
		}
	}
	var zeroValue A
	return zeroValue, -1
}

func FindLastIndexed[A any](slice []A, cb func(A, int) bool) (A, int) {
	for i := len(slice) - 1; i >= 0; i-- {
		if cb(slice[i], i) {
			return slice[i], i
		}
	}
	var zeroValue A
	return zeroValue, -1
}

// find index
func FindFirstIndex[A any](slice []A, cb func(A) bool) int {
	for i, item := range slice {
		if cb(item) {
			return i
		}
	}
	return -1
}

func FindFirstIndexIndexed[A any](slice []A, cb func(A, int) bool) int {
	for i, item := range slice {
		if cb(item, i) {
			return i
		}
	}
	return -1
}

func FindLastIndex[A any](slice []A, cb func(A) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if cb(slice[i]) {
			return i
		}
	}
	return -1
}

func FindLastIndexIndexed[A any](slice []A, cb func(A, int) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if cb(slice[i], i) {
			return i
		}
	}
	return -1
}

// indexOf
func FirstIndexOf[A comparable](slice []A, element A) int {
	for i, item := range slice {
		if item == element {
			return i
		}
	}
	return -1
}

func LastIndexOf[A comparable](slice []A, element A) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == element {
			return i
		}
	}
	return -1
}

// contains
func Contains[A comparable](slice []A, element A) bool {
	return FirstIndexOf(slice, element) != -1
}

// checkers
func Some[A comparable](slice []A, cb func(A) bool) bool {
	for _, item := range slice {
		if cb(item) {
			return true
		}
	}
	return false
}

func SomeIndexed[A comparable](slice []A, cb func(A, int) bool) bool {
	for i, item := range slice {
		if cb(item, i) {
			return true
		}
	}
	return false
}

func Every[A comparable](slice []A, cb func(A) bool) bool {
	for _, item := range slice {
		if !cb(item) {
			return false
		}
	}
	return true
}

func EveryIndexed[A comparable](slice []A, cb func(A, int) bool) bool {
	for i, item := range slice {
		if !cb(item, i) {
			return false
		}
	}
	return true
}

// zip
type ZipItem[A any, B any] struct {
	a A
	b B
}

func Zip[A any, B any](sliceA []A, sliceB []B) []ZipItem[A, B] {
	end := min(len(sliceA), len(sliceB))
	newSlice := make([]ZipItem[A, B], end)
	for i := 0; i < end; i++ {
		newSlice[i] = ZipItem[A, B]{sliceA[i], sliceB[i]}
	}
	return newSlice
}

type ZipIndexedItem[A any, B any] struct {
	a     A
	b     B
	index int
}

func ZipIndexed[A any, B any](sliceA []A, sliceB []B) []ZipIndexedItem[A, B] {
	end := min(len(sliceA), len(sliceB))
	newSlice := make([]ZipIndexedItem[A, B], end)
	for i := 0; i < end; i++ {
		newSlice[i] = ZipIndexedItem[A, B]{sliceA[i], sliceB[i], i}
	}
	return newSlice
}

// synonyms
func Find[A any](slice []A, cb func(A) bool) (A, int) {
	return FindFirst(slice, cb)
}

func FindIndex[A any](slice []A, cb func(A) bool) int {
	return FindFirstIndex(slice, cb)
}

func IndexOf[A comparable](slice []A, element A) int {
	return FirstIndexOf(slice, element)
}
