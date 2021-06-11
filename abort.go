package people-api-resource

func abortOnError(err error) {
	if err != nil {
		panic(err)
	}
}
