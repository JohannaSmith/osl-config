package main

const ReleaseName = "garden-linux-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "garden-linux",
			RepoURL: "https://github.com/cloudfoundry-incubator/garden-linux-release",
			Branch:  "develop",
			Paths: ReportPaths{
				{
					Name: "garden",
					Path: "src/code.cloudfoundry.org/garden",
				},
				{
					Name: "garden-linux",
					Path: "src/code.cloudfoundry.org/garden-linux",
				},
				{
					Name: "garden-integration-tests",
					Path: "src/code.cloudfoundry.org/garden-integration-tests",
				},
				{
					Name: "garden-shed",
					Path: "src/code.cloudfoundry.org/garden-shed",
				},
			},
		},
	}
}
