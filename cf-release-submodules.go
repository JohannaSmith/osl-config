package main

type CfReleaseSubmodules[]struct {
	Name		string
	Path		string
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
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.LtsPath
	}
	return p
}

func NewCfReleaseSubmodules() CfReleaseSubmodules{
	return CfReleaseSubmodules{
		{
			Name:   "cloud_controller_ng",
			Path:   "src/cloud_controller_ng",
			LtsPath:   "src/cloud_controller_ng",
		},
		{
			Name:   "loggregator",
			Path:   "src/loggregator",
			LtsPath:   "src/loggregator",
		},
		{
			Name:   "uaa",
			Path:   "src/uaa-release/src/uaa",
			LtsPath:   "src/uaa",
		},
	}
}
