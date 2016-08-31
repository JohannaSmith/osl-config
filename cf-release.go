package main

const ReleaseName = "cf-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "cf",
			RepoURL: "https://github.com/cloudfoundry/cf-release",
			Branch:  "develop",
			Paths: ReportPaths{
				{
					Name: "loggregator",
					Path: "src/loggregator",
				},
				{
					Name: "collector",
					Path: "src/collector",
				},
				{
					Name: "uaa",
					Path: "src/uaa-release/src/uaa",
				},
				{
					Name: "dea_next",
					Path: "src/dea_next",
				},
				{
					Name: "smoke-tests",
					Path: "src/smoke-tests",
				},
				{
					Name: "cf-acceptance-tests",
					Path: "src/github.com/cloudfoundry/cf-acceptance-tests/",
				},
				{
					Name: "statsd-injector",
					Path: "src/loggregator/src/statsd-injector",
				},
				{
					Name: "hm9000",
					Path: "src/hm9000",
				},
				{
					Name: "nats",
					Path: "src/nats",
				},
				{
					Name: "consul-release",
					Path: "src/consul-release",
				},
				{
					Name: "etcd",
					Path: "src/github.com/coreos/etcd",
				},
				{
					Name: "etcd-metrics-server",
					Path: "src/etcd-metrics-server",
				},
				{
					Name: "warden",
					Path: "src/warden/warden",
				},
				{
					Name: "warden-client",
					Path: "src/warden/warden-client",
				},
				{
					Name: "warden-protocol",
					Path: "src/warden/warden-protocol",
				},
				{
					Name: "em-posix-spawn",
					Path: "src/warden/em-posix-spawn",
				},
			},
		},
	}
}
