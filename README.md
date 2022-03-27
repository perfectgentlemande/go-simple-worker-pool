# go-simple-worker-pool

## Simple worker pool with error checking

Main idea is:
- you need to process some number of items with function that can return error
- you should do that concurrently
- error-returning items are generated randomly

There is package called `dataset` that includes functions that generate test data.

### Case 1: collecting all the errors

When you need to collect all the errors while doing concurrent procession

### Case 2: stopping after the first error

When you need abort the process after the first error