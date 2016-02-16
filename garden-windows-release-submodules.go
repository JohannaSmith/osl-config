package main

type GardenWindowsReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c GardenWindowsReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewGardenWindowsReleaseSubmodules() GardenWindowsReleaseSubmodules {
	return GardenWindowsReleaseSubmodules{
		{
			Name: "Containerizer",
			Path: "src/github.com/cloudfoundry/garden-windows/Containerizer/",
		},
		{
			Name: "garden-windows",
			Path: "src/github.com/cloudfoundry/garden-windows/",
		},
	}
}
