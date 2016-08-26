package main

const ReleaseName = "cf-routing-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "cf-routing",
			RepoURL: "https://github.com/cloudfoundry-incubator/cf-routing-release",
			Branch:  "develop",
			Paths: ReportPaths{
				{
					Name: "cf-tcp-router",
					Path: "src/code.cloudfoundry.org/cf-tcp-router",
				},
				{
					Name: "routing-api",
					Path: "src/code.cloudfoundry.org/routing-api",
				},
				{
					Name: "tcp-emitter",
					Path: "src/code.cloudfoundry.org/tcp-emitter",
				},
				{
					Name: "gorouter",
					Path: "src/code.cloudfoundry.org/gorouter",
				},
			},
		},
	}
}
