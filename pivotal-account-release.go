package main

const ReleaseName = "pivotal-account-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "pivotal-account",
			RepoURL: "git@github.com:pivotal-cf/pivotal-account-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "pivotal-account",
					Path: "src/pivotal-account",
				},
			},
		},
	}
}
