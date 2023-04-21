package iteration

func Repeat(x string) string {
	var repeated string

	for i :=0; i < 5; i++ {
		repeated = repeated + x
	}

	return repeated
}