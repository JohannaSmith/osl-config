package main

type GardenLinuxReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c GardenLinuxReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewGardenLinuxReleaseSubmodules() GardenLinuxReleaseSubmodules {
	return GardenLinuxReleaseSubmodules{
		{
			Name: "garden",
			Path: "src/code.cloudfoundry.org/garden",
		},
		{
			Name: "garden-linux",
			Path: "src/code.cloudfoundry.org/garden-linux",
		},
		{
			Name: "garden-integration-tests",
			Path: "src/code.cloudfoundry.org/garden-integration-tests",
		},
		{
			Name: "garden-shed",
			Path: "src/code.cloudfoundry.org/garden-shed",
		},
	}
}
