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
		goodItems, errs := processEverything(ctx, cases[i].items, 10, cases[i].errorItemIDs)
		if len(errs) != len(cases[i].errorItemIDs) {
			t.Errorf("expected %d errors, got: %d", len(cases[i].errorItemIDs), len(errs))
		}

		if len(errs) == 0 {
			if len(goodItems) != len(cases[i].items) {
				t.Errorf("case_id: %d: expected good items: %d, got: %d", i, len(cases[i].items), len(goodItems))
			}
		}
	}
}
