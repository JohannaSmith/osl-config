package main

type Environment struct {
	Name   string
	Bucket string
}

func NewEnvironment() Environment {
	return Environment{
		Name:   "production",
		Bucket: "osl-reports",
	}
}
