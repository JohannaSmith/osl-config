package main

const ReleaseName = "ops-manager-installation-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "ops-manager-installation",
			RepoURL: "git@github.com:pivotal-cf/installation",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "ops-manager-web",
					Path: "web",
				},
			},
		},
	}
}
