package main

const ReleaseName = "diego-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "diego",
			RepoURL: "https://github.com/cloudfoundry-incubator/diego-release",
			Branch:  "develop",
			Paths: ReportPaths{
				{
					Name: "auction",
					Path: "src/code.cloudfoundry.org/auction",
				},
				{
					Name: "auctioneer",
					Path: "src/code.cloudfoundry.org/auctioneer",
				},
				{
					Name: "bbs",
					Path: "src/code.cloudfoundry.org/bbs",
				},
				{
					Name: "benchmarkbbs",
					Path: "src/code.cloudfoundry.org/benchmarkbbs",
				},
				{
					Name: "fileserver",
					Path: "src/code.cloudfoundry.org/fileserver",
				},
				{
					Name: "executor",
					Path: "src/code.cloudfoundry.org/executor",
				},
				{
					Name: "rep",
					Path: "src/code.cloudfoundry.org/rep",
				},
				{
					Name: "debugserver",
					Path: "src/code.cloudfoundry.org/debugserver",
				},
				{
					Name: "route-emitter",
					Path: "src/code.cloudfoundry.org/route-emitter",
				},
				{
					Name: "routing-info",
					Path: "src/github.com/cloudfoundry-incubator/routing-info",
				},
				{
					Name: "runtimeschema",
					Path: "src/code.cloudfoundry.org/runtimeschema",
				},
				{
					Name: "locket",
					Path: "src/code.cloudfoundry.org/locket",
				},
				{
					Name: "cfhttp",
					Path: "src/code.cloudfoundry.org/cfhttp",
				},
				{
					Name: "buildpackapplifecycle",
					Path: "src/code.cloudfoundry.org/buildpackapplifecycle",
				},
				{
					Name: "dockerapplifecycle",
					Path: "src/code.cloudfoundry.org/dockerapplifecycle",
				},
				{
					Name: "diegocanaryapp",
					Path: "src/code.cloudfoundry.org/diegocanaryapp",
				},
				{
					Name: "consuladapter",
					Path: "src/code.cloudfoundry.org/consuladapter",
				},
				{
					Name: "diego-ssh",
					Path: "src/code.cloudfoundry.org/diego-ssh",
				},
				{
					Name: "cacheddownloader",
					Path: "src/code.cloudfoundry.org/cacheddownloader",
				},
				{
					Name: "healthcheck",
					Path: "src/code.cloudfoundry.org/healthcheck",
				},
				{
					Name: "vizzini",
					Path: "src/code.cloudfoundry.org/vizzini",
				},
			},
		},
	}
}
