package main

type Environment struct {
	Name   string
	Bucket string
}

func NewEnvironment() Environment {
	return Environment{
		Name:   "staging",
		Bucket: "osl-reports-staging",
	}
}
