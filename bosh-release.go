package main

const ReleaseName = "bosh-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "bosh",
			RepoURL: "https://github.com/cloudfoundry/bosh",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "bosh",
					Path: ".",
				},
				{
					Name: "bosh-davcli",
					Path: "go/src/github.com/cloudfoundry/bosh-davcli",
				},
				{
					Name: "bosh-agent",
					Path: "go/src/github.com/cloudfoundry/bosh-agent",
				},
			},
		},
	}
}
