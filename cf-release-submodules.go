package main

type CfReleaseSubmodules[]struct {
	Name   string
	Path string
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
