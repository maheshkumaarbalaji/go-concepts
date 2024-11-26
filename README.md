# Go Concepts

This repository contains implementation in Go, for several core concepts in Computer Science like *Searching*, *Sorting*, *Data Structures* etc.

## Packages

This repository contains the following packages:
- **search** - Contains the implementation for linear and binary search using slices. This can also be applied to arrays in Go.
- **sorting** - Contains the implementation for various sorting techniques for slice elements in Go.
- **ds** - Implements the prefix tree data structure in Go.

### Testing

Each package also contains the unit test scripts which can be identified by the "_test.go" suffix present in the files. To run all the test scripts, execute the following command.

```bash
go test ./lib/... -v -cover
```

- The **-v** command line option enables the verbose logs to be printed as the test scripts are being executed.
- The **-cover** command line option displays the code coverage information for the test cases executed.