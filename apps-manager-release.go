package main

const ReleaseName = "apps-manager-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "apps-manager",
			RepoURL: "git@github.com:pivotal-cf/apps-manager-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "invitations",
					Path: "src/invitations",
				},
				{
					Name: "app-usage-service",
					Path: "src/app-usage-service",
				},
				{
					Name: "apps-manager-js",
					Path: "src/apps-manager-js",
				},
			},
		},
	}
}
