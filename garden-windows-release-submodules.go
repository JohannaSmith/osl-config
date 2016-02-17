package main

type GardenWindowsReleaseSubmodules []struct {
	Name    string // the name of the project
	Path    string // the path of the project in the latest release
	LtsPath string // the path of the project in the latest release
}

func (c GardenWindowsReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func (c GardenWindowsReleaseSubmodules) LtsPaths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.LtsPath
	}
	return p
}

func NewGardenWindowsReleaseSubmodules() GardenWindowsReleaseSubmodules {
	return GardenWindowsReleaseSubmodules{
		{
			Name:    "GardenWindowsRelease",
			Path:    "GardenWindowsRelease",
			LtsPath: "GardenWindowsRelease",
		},
		{
			Name:    "Containerizer",
			Path:    "src/github.com/cloudfoundry/garden-windows/Containerizer/",
			LtsPath: "src/github.com/cloudfoundry-incubator/garden-windows/Containerizer/",
		},
		{
			Name:    "garden-windows",
			Path:    "src/github.com/cloudfoundry/garden-windows/",
			LtsPath: "src/github.com/cloudfoundry-incubator/garden-windows/",
		},
	}
}
