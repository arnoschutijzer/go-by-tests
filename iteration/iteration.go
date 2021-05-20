package iteration

func Repeat(sequence string, count int) string {
	// zero-valued init
	var repeated string
	for i := 0; i < count; i++ {
		repeated += sequence
	}
	return repeated
}
