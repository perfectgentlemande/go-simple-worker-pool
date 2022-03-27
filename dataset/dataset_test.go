package dataset

import (
	"testing"
)

func TestFillDefaultData(t *testing.T) {
	type testCase struct {
		numOfItems int
	}

	cases := []testCase{
		{numOfItems: 0}, {numOfItems: 1000}, {numOfItems: 1000000},
	}

	for i := range cases {
		dataSet := FillDefaultData(cases[i].numOfItems)

		if len(dataSet) != cases[i].numOfItems {
			t.Errorf("expected %d items, got: %d", cases[i].numOfItems, len(dataSet))
		}

		if len(dataSet) == 0 {
			continue
		}

		if dataSet[len(dataSet)-1].ID != len(dataSet)-1 {
			t.Errorf("expected last item's ID: %d, got: %d", len(dataSet)-1, dataSet[len(dataSet)-1].ID)
		}
	}
}

func TestGenerateAlLErrorIDs(t *testing.T) {
	type testCase struct {
		numOfItems int
	}

	cases := []testCase{
		{numOfItems: 0}, {numOfItems: 1000}, {numOfItems: 1000000},
	}

	for i := range cases {
		errorIDs := GenereateAllErrorIDs(cases[i].numOfItems)

		if cases[i].numOfItems != len(errorIDs) {
			t.Errorf("expected num of items: %d, got: %d", cases[i].numOfItems, len(errorIDs))
			continue
		}

		for j := range errorIDs {
			if errorIDs[j] != j {
				t.Errorf("expected errorID: %d, got: %d for num of items: %d", j, errorIDs[j], cases[i].numOfItems)
				break
			}
		}
	}
}
func TestGenerateRandomErrorIDs(t *testing.T) {
	type testCase struct {
		numOfItems int
	}

	cases := []testCase{
		{numOfItems: 0}, {numOfItems: 1000}, {numOfItems: 1000000},
	}

	for i := range cases {
		randomErrorIDs := GenerateRandomErrorIDs(cases[i].numOfItems)

		for j := range randomErrorIDs {
			if randomErrorIDs[j] < 0 || randomErrorIDs[j] >= cases[i].numOfItems {
				t.Errorf("wrong index of error element: %d, len of sample data: %d", randomErrorIDs[j], cases[i].numOfItems)
			}
		}
	}
}
func TestShouldReturnError(t *testing.T) {
	type testCase struct {
		item         SomeStruct
		errorItemIDs []int
		returnsError bool
	}

	cases := []testCase{
		{
			item: SomeStruct{
				ID: 0,
			},
			errorItemIDs: []int{},
			returnsError: false,
		},
		{
			item: SomeStruct{
				ID: 0,
			},
			errorItemIDs: []int{1, 2, 3},
			returnsError: false,
		},
		{
			item: SomeStruct{
				ID: 0,
			},
			errorItemIDs: []int{0, 5, 6},
			returnsError: true,
		},
	}

	for i := range cases {
		err := ShouldReturnError(cases[i].item, cases[i].errorItemIDs)
		if err != nil != cases[i].returnsError {
			t.Errorf("returned error: %v, expected: %v", err != nil, cases[i].returnsError)
		}
	}

}
