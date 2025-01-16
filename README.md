# Go Leetcodes 🎉

Welcome to the Go Leetcodes repository! This repository contains a collection of Leetcode problems solved in Go with unit tests as well as speed testing.

## Testing

### General testing:

In order to run a test, simply navigate to the function you want to test and run `go test -v`

```bash
go test -v
```

## Coverage:
```bash
$ go test -v -coverprofile=coverage.out
$ go tool cover -html=coverage.out
```

## Benchmarking

```bash
$ go test -bench .
```

## Time Functions

In the root directory of the project, run:

This will default to timing ALL functions with a default input size of 1,000,000.
```bash
$ ./time_functions
```
### Timing specific groups or functions``

The syntax for timing in the command line is `./time_functions [method] [input_size]`

Method can be a group, for example `ArraysAndHashing` or `TwoPointers`, or a specific function, for example `ContainsDuplicate` or `ValidPalindrome`.

Input size can be any integer (However, I would recommend staying below 100,000,000 as it can take a very long time for some functions).

Examples:

```bash
$ ./time_functions ArraysAndHashing 10
```

```bash
$ ./time_functions TwoPointers 10
```

```bash
$ ./time_functions ContainsDuplicate 10
```

```bash
$ ./time_functions ValidPalindrome 10
```

If you ever get confused on which methods you can use, you can use the `help` flag to see a list of available methods as well as the instructions on how to use the command

```bash
$ ./time_functions help
```
