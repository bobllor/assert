## About

A simple assertion library for Go.

## Install

```sh
go get https://github.com/bobllor/assert
```

## Usage

Type `assert` and use the available functions to get started.

```go
_, err := os.ReadDir("/non/existent/path")
assert.Nil(t, err) // example_test.go:30 (assert.TestExampleReadMe assertion failed): value is not nil (open ...)

assert.True(t, 1 == 2-1) // passes assertion
```

Multiple args are supported with functions that uses `All` or `Any`:

```go
assert.TrueAll(t, 1 == 1, 2 == 2, 3 == 30/10) // passes assertion, all is true

assert.ContainsAny(t, "welcome to the jungle", "x3d", "jungle") // passes assertion, "jungle" is found despite "x3d" not found

assert.NotContainsAny(t, "sentence string", "z", "sentence") // fails assertion due to "sentence"
```