// Package dataset for generating example datasets
package dataset

import (
	"fmt"
	"math/rand"
	"strconv"
)

type SomeStruct struct {
	ID  int
	Val string
}
type FullOutput struct {
	Result SomeStruct
	Error  error
}

// FillDefaultData for generating dataset
func FillDefaultData(numberOfItems int) []SomeStruct {
	someItems := make([]SomeStruct, numberOfItems)

	for i := range someItems {
		someItems[i] = SomeStruct{
			ID:  i,
			Val: "val " + strconv.Itoa(i),
		}
	}

	return someItems
}

func GenereateAllErrorIDs(n int) []int {
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = i
	}

	return res
}

// GenerateRandomErrorIDs for generating ids that should return error
func GenerateRandomErrorIDs(n int) []int {
	if n == 0 {
		return make([]int, 0)
	}

	numberOfErrors := rand.Intn(n)

	res := make([]int, numberOfErrors)
	for i := 0; i < numberOfErrors; i++ {
		res[i] = rand.Intn(n)
	}

	return res
}

// ShouldReturnError for imitation of error returning process
func ShouldReturnError(item SomeStruct, errorItemIDs []int) error {
	for i := range errorItemIDs {
		if errorItemIDs[i] == item.ID {
			return fmt.Errorf("id %d returned error", item.ID)
		}
	}

	return nil
}

// For educational purposes we'll use the id's to generate errors
func DoSomething(item SomeStruct, errorItemIDs []int) (SomeStruct, error) {
	if err := ShouldReturnError(item, errorItemIDs); err != nil {
		return SomeStruct{}, err
	}

	item.Val += " was changed"
	return item, nil
}
