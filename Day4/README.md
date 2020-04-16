# Notes

* String methods from standard library are "static" methods. So `string.TrimSuffix(str, suffix)` instead of `str.TrimSuffix(suffix)`
* You can combine an assignment and a condition in an `if`. For example, `if err := myFunc(); err != nil {}`
* You can unpack an slice with `...` to pass it to a function with variadic arguments. See: [three dots](https://programming.guide/go/three-dots-ellipsis.html)