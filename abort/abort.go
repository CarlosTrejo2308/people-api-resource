package abort

// AbortOnError recives an error
// and panic when there's an error
func AbortOnError(err error) {
	if err != nil {
		panic(err)
	}
}
