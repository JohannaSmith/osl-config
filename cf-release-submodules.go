package main

type CfReleaseSubmodules []struct {
	Name    string
	Path    string
	LtsPath string
}

func (c CfReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Path
	}
	return p
}

func (c CfReleaseSubmodules) LtsPaths() []string {
	lts := c.Lts()
	p := make([]string, len(lts))
	for k, v := range lts {
		p[k] = v.LtsPath
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
			Name: "cf-routing-release",
			Path: "src/cf-routing-release",
		},
	}
}
