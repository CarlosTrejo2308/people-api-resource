package abort

func AbortOnError(err error) {
	if err != nil {
		panic(err)
	}
}
