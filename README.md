# Go Leetcodes 🎉

Welcome to the Go Leetcodes repository! This repository contains a collection of Leetcode problems solved in Go with unit tests as well as speed testing.

## Testing

### General testing:
```bash
go test -v
```

### Specific testing:
```bash
go test -v -run <test_name>

# Example
# go test -v -run TestContainsDuplicate
```

## Coverage:
```bash
go test -v -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Benchmarking

```bash
go test -bench .
```

## Speed Testing
In the root directory of the project, run:

```bash
./time_functions
```
