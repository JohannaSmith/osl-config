package main

const ReleaseName = "notifications-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "notifications",
			RepoURL: "git@github.com:cloudfoundry-incubator/notifications-release.git",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "notifications",
					Path: "src/notifications",
				},
			},
		},
	}
}
