package main

type Repos []struct {
	Name   string
	Repo   string
	Branch string
}

func NewRepos() Repos {
	return Repos{
		{
			Name:   "cloud_controller_ng",
			Repo:   "https://github.com/cloudfoundry/cloud_controller_ng.git",
			Branch: "master",
		},
		{
			Name:   "loggregator",
			Repo:   "https://github.com/cloudfoundry/loggregator.git",
			Branch: "master",
		},
	}
}
