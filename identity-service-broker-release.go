package main

const ReleaseName = "identity-service-broker-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "identity-service-broker",
			RepoURL: "git@github.com:pivotal-cf/identity-service-broker-release.git",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "identity-service-broker",
					Path: "src/identity-service-broker",
				},
			},
		},
	}
}
