package main

type CfRoutingReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c CfRoutingReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewCfRoutingReleaseSubmodules() CfRoutingReleaseSubmodules {
	return CfRoutingReleaseSubmodules{
		{
			Name: "cf-tcp-router",
			Path: "src/github.com/cloudfoundry-incubator/cf-tcp-router",
		},
		{
			Name: "routing-api",
			Path: "src/github.com/cloudfoundry-incubator/routing-api",
		},
		{
			Name: "tcp-emitter",
			Path: "src/github.com/cloudfoundry-incubator/tcp-emitter",
		},
		{
			Name: "gorouter",
			Path: "src/github.com/cloudfoundry/gorouter",
		},
	}
}
