# Notes

* `Println` doesn't do formatting. Use `Printf` instead
* You have to seed the random generator `rand.Seed(xxx)`. If you want a different result each time `rand.Seed(time.Now().UnixNano())`
* For arrays, square brackets are placed before the data type. `[]int` vs `int[]`
* Arrays are fixed-length. Length is part of the type definition, too. You can't pass `[5]int{ 1,2,3,4,5 }` into a function that receives `[4]int`
* `range` is like a `foreach` with an index and the value. `for i, n := range array {}` vs `foreach (var item in array.Select((n, i) => new { n, i })`
* variadic functions are like functions with `params`
* Go has slices like Python: `array[start:stop]`