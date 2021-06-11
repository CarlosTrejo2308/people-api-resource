package peopleApiResource

func abortOnError(err error) {
	if err != nil {
		panic(err)
	}
}
