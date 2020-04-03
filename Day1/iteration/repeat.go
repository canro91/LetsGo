package iteration

const repeatedCount = 5

func Repeat(str string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += str
	}
	return repeated
}