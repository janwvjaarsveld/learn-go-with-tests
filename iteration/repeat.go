package iteration

// Repeat takes a character and repeats it 5 times
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}
	return
}
