# Go Leetcodes 🎉

Welcome to the Go Leetcodes repository! This repository contains a collection of Leetcode problems solved in Go with unit tests as well as speed testing.

## Testing

### Testing ALL

Run the following command in the root directory of the project to run all tests

```bash
$ go test ./...
```

can also add the `-v` flag to get more information or the `-cover` flag to get coverage information. (NOTE: Timing Functions do not have tests and same with helper functions as they are just function generics used to reuse functions rather than having repetitive code)

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
$ ./leetcodes
```
### Timing specific groups or functions``

The syntax for timing in the command line is `./leetcodes [method] [input_size]`

Method can be a group, for example `ArraysAndHashing` or `TwoPointers`, or a specific function, for example `ContainsDuplicate` or `ValidPalindrome`.

Input size can be any integer (However, I would recommend staying below 100,000,000 as it can take a very long time for some functions).

Examples:

```bash
$ ./leetcodes ArraysAndHashing 10
```

```bash
$ ./leetcodes TwoPointers 10
```

```bash
$ ./leetcodes ContainsDuplicate 10
```

```bash
$ ./leetcodes ValidPalindrome 10
```

If you ever get confused on which methods you can use, you can use the `help` flag to see a list of available methods as well as the instructions on how to use the command

```bash
$ ./leetcodes help
```

## Contributing

Contributions are welcome! If you have any suggestions or improvements, please open an issue or submit a pull request.

When adding a solution to a leetcode problem, make sure to take care of the following areas:

- Add a folder for the problem in the appropriate category (e.g., ArraysAndHashing, TwoPointers, etc.)
- In this folder, create a file with the name of the problem (e.g., "two_sum.go") and add the solution to the problem.
- Add a test file with the name of the problem (e.g., "two_sum_test.go") and add the test cases for the problem.
- (Optional) Add a README.md file with a brief description of the problem and the solution.

- Add it to the timings in the `cmd/leetcodes/timers` as `TimeMethodName` folder.
- Add it to the main function in the `cmd/leetcodes/main.go` file.
- Add it to the `TimeAllFunctions` function in the `cmd/leetcodes/timers/TimeAllFunctions.go` file.
- (Optional BUT Recommended) If helper functions are needed, add them to the `helper_functions.go` file. An example would be if there does not already exist a function that handles the type of input and output 

Basically, just make sure it's integrated with the CLI for timing the functions

If you want to contribute to refactoring the code, or adding new functionality feel free to do that and submit a pr and I will look over it.

## Forking

You can also fork this repository and make your own changes.

For example, if you want to use a different solution to a leetcode problem, you can fork this repository and make your own changes and then run the timer on your solution and see how fast it is with different values of `n`.
