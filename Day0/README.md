# Notes

## Tests

* Tests should be in a file XXX_test.go
* Test names should start with `Test`
* Tests should have a single param `t *testing.T`
* There isn't anything like `Assert.AreEqual` by default. See: [testify](https://github.com/stretchr/testify)
* In `Errorf`, `%q` encloses a string inside quotes
* Examples should contain an "Output" comment
```
func ExampleHello() {
    fmt.Println(Hello())
    // Output:
    // Hello, world!
}
```

## Syntax and other

* Param types follow param names in functions `func MyFunc(str string)`
* Return type are between closing parenthesis and opening bracket `func MyFunc(str string) string {}`

* Semicolons are optional at the end of line

* Unused imports give compilation error

* There is no method overloading

* `:=` vs `=`: declaration + assignment vs assignment only. See [Difference between := and = operators in Go](https:stackoverflow.com/a/17891297)

* `case` statements inside `switch` don't have `break`
* Multiple `case`s are separated by comma instead of body-less `case`