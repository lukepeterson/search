# Binary search in Go

This repository contains a simple implementation of the binary search algorithm in Go. Binary search is an efficient algorithm for finding an item from a sorted list of items. It works by repeatedly dividing in half the portion of the list that could contain the item, until we've narrowed down the possible locations to just one.

## How it works

The binary search function takes a sorted slice of integers (`data`) and a target integer (`target`) to find within the slice. The function returns the index of the target if it is present in the slice. If the target is not found, the function returns `-1`.

The binary search algorithm runs in O(log n) time, making it much faster than an O(n) linear search for large data sets.

## Usage

1. **Clone**
   ```bash
   git clone git@github.com:lukepeterson/search.git 
   ```

2. **Navigate**
   ```bash
   cd search
   ```

3. **Run**
   ```bash
   go run main.go
   ```

## Example

Given a sorted slice `data`:

```go
data := []int{1, 2, 4, 5, 7, 8, 12}
```

And a target value `target`:

```go
target := 5
```

The output will be:

```
Target 5 found at index 3
```

If the target is not found, the output will be:

```
Target not found
```