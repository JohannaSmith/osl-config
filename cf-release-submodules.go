package main

type CfReleaseSubmodules []struct {
	Name    string // the name of the project
	Path    string // the path of the project in the latest release
	LtsPath string // the path in the lts release
}

func (c CfReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func (c CfReleaseSubmodules) LtsPaths() []string {
	lts := c.Lts()
	p := make([]string, len(lts))
	for k, v := range lts {
		p[k] = v.Name + ":" + v.LtsPath
	}
	return p
}

func (c CfReleaseSubmodules) Lts() CfReleaseSubmodules {
	var out CfReleaseSubmodules
	for _, v := range c {
		if v.LtsPath != "" {
			out = append(out, v)
		}
	}
	return out
}

func NewCfReleaseSubmodules() CfReleaseSubmodules {
	return CfReleaseSubmodules{
		{
			Name:    "cloud_controller_ng",
			Path:    "src/cloud_controller_ng",
			LtsPath: "src/cloud_controller_ng",
		},
		{
			Name:    "loggregator",
			Path:    "src/loggregator",
			LtsPath: "src/loggregator",
		},
		{
			Name:    "collector",
			Path:    "src/collector",
			LtsPath: "src/collector",
		},
		{
			Name:    "uaa",
			Path:    "src/uaa-release/src/uaa",
			LtsPath: "src/uaa",
		},
		{
			Name:    "gorouter",
			Path:    "src/github.com/cloudfoundry/gorouter",
			LtsPath: "src/github.com/cloudfoundry/gorouter",
		},
		{
			Name:    "dea_next",
			Path:    "src/dea_next",
			LtsPath: "src/dea_next",
		},
		{
			Name:    "smoke-tests",
			Path:    "src/smoke-tests",
			LtsPath: "src/smoke-tests",
		},
		{
			Name:    "statsd-injector",
			Path:    "src/statsd-injector",
			LtsPath: "src/statsd-injector",
		},
		{
			Name:    "hm9000",
			Path:    "src/hm9000",
			LtsPath: "src/hm9000",
		},
		{
			Name:    "cf-routing-release",
			Path:    "src/cf-routing-release",
			LtsPath: "src/github.com/cloudfoundry-incubator/routing-api",
		},
		{
			Name:    "nats",
			Path:    "src/nats",
			LtsPath: "src/nats",
		},
		{
			Name:    "route-registrar",
			Path:    "src/github.com/cloudfoundry-incubator/route-registrar",
			LtsPath: "src/route_registrar",
		},
		{
			Name: "consul-release",
			Path: "src/consul-release",
		},
		{
			Name:    "etcd",
			Path:    "src/github.com/coreos/etcd",
			LtsPath: "src/github.com/coreos/etcd",
		},
		{
			Name:    "etcd-metrics-server",
			Path:    "src/etcd-metrics-server",
			LtsPath: "src/etcd-metrics-server",
		},
		{
			Name:    "warden",
			Path:    "src/warden/warden",
			LtsPath: "src/warden/warden",
		},
		{
			Name:    "warden-client",
			Path:    "src/warden/warden-client",
			LtsPath: "src/warden/warden-client",
		},
		{
			Name:    "warden-protocol",
			Path:    "src/warden/warden-protocol",
			LtsPath: "src/warden/warden-protocol",
		},
		{
			Name: "em-posix-spawn",
			Path: "src/warden/em-posix-spawn",
		},
	}
}
