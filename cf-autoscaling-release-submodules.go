package main

type CfAutoscalingReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c CfAutoscalingReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewCfAutoscalingReleaseSubmodules() CfAutoscalingReleaseSubmodules {
	return CfAutoscalingReleaseSubmodules{
		{
			Name: "cf-autoscaling",
			Path: "src/cf-autoscaling",
		},
	}
}
