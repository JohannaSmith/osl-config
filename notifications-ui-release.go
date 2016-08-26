package main

const ReleaseName = "notifications-ui-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "notifications-ui",
			RepoURL: "git@github.com:pivotal-cf/notifications-ui-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "notifications-ui",
					Path: "src/notifications-ui",
				},
			},
		},
	}
}
