package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/perfectgentlemande/go-simple-worker-pool/dataset"
)

func doWork(ctx context.Context, wg *sync.WaitGroup, itemsInputCh <-chan dataset.SomeStruct, itemsOutputCh chan<- dataset.FullOutput, errorItemIDs []int) {
	defer wg.Done()

	for {
		// we do not read anything when the context is closed
		select {
		case <-ctx.Done():
			return
		case item, ok := <-itemsInputCh:
			if !ok {
				return
			}

			res := dataset.FullOutput{}
			item, err := dataset.DoSomething(item, errorItemIDs)
			if err != nil {
				res.Error = fmt.Errorf("failed to do something: %w", err)
			}
			res.Result = item

			// we do not send anything when the context is closed
			select {
			case <-ctx.Done():
				return
			case itemsOutputCh <- res:
				continue
			}
		}
	}
}

func processEverything(ctx context.Context, items []dataset.SomeStruct, wgCount int, errorItemIDs []int) ([]dataset.SomeStruct, []error) {
	itemsInputCh := make(chan dataset.SomeStruct)
	itemsOutputCh := make(chan dataset.FullOutput)
	wg := &sync.WaitGroup{}

	output := make([]dataset.SomeStruct, 0)
	outputErrors := make([]error, 0)

	// worker pool input
	// we do not send anything when the context is closed
	go func(ctx context.Context) {
		defer close(itemsInputCh)

		for i := range items {
			select {
			case <-ctx.Done():
				return
			case itemsInputCh <- items[i]:
				continue
			}
		}
	}(ctx)

	// worker pool itself
	go func() {
		for i := 0; i < wgCount; i++ {
			wg.Add(1)

			// worker itself
			go doWork(ctx, wg, itemsInputCh, itemsOutputCh, errorItemIDs)
		}
		wg.Wait()
		close(itemsOutputCh)
	}()

	// worker pool output
	for {
		select {
		case <-ctx.Done():
			return nil, outputErrors
		case res, ok := <-itemsOutputCh:
			if !ok {
				return output, outputErrors
			}
			output = append(output, res.Result)
			if res.Error != nil {
				outputErrors = append(outputErrors, res.Error)
			}
		}
	}
}

func main() {
	items := dataset.FillDefaultData(20)
	errorItemIDs := dataset.GenerateRandomErrorIDs(len(items))
	ctx := context.Background()

	fmt.Printf("generated %d items\n", len(items))
	fmt.Printf("items: %v\n", items)
	fmt.Printf("generated %d errors\n", len(errorItemIDs))
	fmt.Printf("error item IDs: %v\n", errorItemIDs)

	goodItems, errs := processEverything(ctx, items, len(items), errorItemIDs)
	if len(errs) != 0 {
		fmt.Printf("there were %d actual errors in \n", len(errs))
		fmt.Printf("actual errors: %v\n", errs)
	}

	fmt.Printf("got %d good items\n", len(goodItems))
	fmt.Printf("good items: %v\n", goodItems)
}
