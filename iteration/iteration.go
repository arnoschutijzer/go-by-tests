package iteration

func Repeat(sequence string) string {
	// zero-valued init
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += sequence
	}
	return repeated
}
