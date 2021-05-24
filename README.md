# go-by-tests

## run tests
```bash
$ go test ./... -cover -v
```

## create an example test
```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```

## check out documentation
`godoc -http 8000` and find your package

## run benchmark tests
`go test -bench=.`
