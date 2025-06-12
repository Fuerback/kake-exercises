## Exercise 2

Now we want to build a interface that supports the following functions:

```go
getMaxColumnSize() int
isSquare() bool
```

Then implement this interface to take a 2D array as input.

### Example 1
```go
md := NewMapDetailImpl([][]int{
	{1, 2, 3},
	{4, 5, 4},
	{1, 1, 1},
})
fmt.Println(md.getMaxColumnSize())
fmt.Println(md.isSquare())
```

```go
3
true
```

### Example 2
```go
md := NewMapDetailImpl([][]int{
	{1},
	{4, 5, 4, 5, 5, 1},
	{},
	{7, 8, 9},
})
fmt.Println(md.getMaxColumnSize())
fmt.Println(md.isSquare())
```

```go
6
false
```

