package main

import (
	"context"
	"testing"

	"github.com/perfectgentlemande/go-simple-worker-pool/dataset"
)

func TestProcessEverything(t *testing.T) {
	ctx := context.Background()

	type testCase struct {
		items        []dataset.SomeStruct
		errorItemIDs []int
	}

	cases := []testCase{
		{
			items: dataset.FillDefaultData(0),
		},
		{
			items: dataset.FillDefaultData(1),
		},
		{
			items:        dataset.FillDefaultData(2),
			errorItemIDs: []int{1},
		},
		{
			items:        dataset.FillDefaultData(2),
			errorItemIDs: []int{0},
		},
		{
			items: dataset.FillDefaultData(10),
		},
		{
			items:        dataset.FillDefaultData(10),
			errorItemIDs: []int{4, 6, 9},
		},
		{
			items:        dataset.FillDefaultData(10),
			errorItemIDs: dataset.GenereateAllErrorIDs(10),
		},
		{
			items: dataset.FillDefaultData(10000),
		},
		{
			items:        dataset.FillDefaultData(10000),
			errorItemIDs: []int{4, 6, 8311, 9, 8, 898, 254, 3751, 3718, 22, 9072, 3234, 1223, 444, 6908, 807, 5516},
		},
		{
			items:        dataset.FillDefaultData(10000),
			errorItemIDs: dataset.GenereateAllErrorIDs(10000),
		},
	}

	for i := range cases {
		goodItems, err := processEverything(ctx, cases[i].items, 10, cases[i].errorItemIDs)
		if err != nil && len(cases[i].errorItemIDs) == 0 {
			t.Errorf("unexpected error occurred: %v", err)
			continue
		}

		if err != nil && len(goodItems) != 0 {
			t.Errorf("got error: %v, but there are %d elements in output", err, len(cases[i].items))
			continue
		}

		if err == nil && len(cases[i].items) != len(goodItems) {
			t.Errorf("expected: %d items, but got %d items", len(cases[i].items), len(goodItems))
			continue
		}
	}
}
