package main

type CfReleaseSubmodules[]struct {
	Name   string
	Path string
}

func (c CfReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Path
	}
	return p
}

func NewCfReleaseSubmodules() CfReleaseSubmodules{
	return CfReleaseSubmodules{
		{
			Name:   "cloud_controller_ng",
			Path:   "src/cloud_controller_ng",
		},
		{
			Name:   "loggregator",
			Path:   "src/loggregator",
		},
	}
}
