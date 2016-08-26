package main

const ReleaseName = "cf-autoscaling-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "cf-autoscaling",
			RepoURL: "git@github.com:pivotal-cf/cf-autoscaling-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "cf-autoscaling",
					Path: "src/cf-autoscaling",
				},
			},
		},
	}
}
