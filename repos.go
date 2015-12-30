package main

type Repo struct {
	Name   string
	Repo   string
	Branch string
}

type Repos []Repo

type ReposStructure struct {
	Repos []Repo
}

func (rs *ReposStructure) Flat() (Repos, error) {
	return rs.Repos, nil
}

func NewRepos() *ReposStructure {
	return &ReposStructure{
		Repos: []Repo {
			{
				Name: "cloud_controller_ng",
				Repo: "https://github.com/cloudfoundry/cloud_controller_ng.git",
				Branch: "master",
			},
			{
				Name: "loggregator",
				Repo: "https://github.com/cloudfoundry/loggregator.git",
				Branch: "master",
			},
		},
	}
}
