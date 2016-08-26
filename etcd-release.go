package main

const ReleaseName = "etcd-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "etcd",
			RepoURL: "git@github.com:cloudfoundry-incubator/etcd-release.git",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "etcd",
					Path: "src/github.com/coreos/etcd",
				},
				{
					Name: "etcd-metrics-server",
					Path: "src/etcd-metrics-server",
				},
			},
		},
	}
}
