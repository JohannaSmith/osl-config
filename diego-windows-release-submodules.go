package main

type DiegoWindowsReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c DiegoWindowsReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Path
	}
	return p
}

func NewDiegoWindowsReleaseSubmodules() DiegoWindowsReleaseSubmodules {
	return DiegoWindowsReleaseSubmodules{
		{
			Name: "greenhouse-install-script-generator",
			Path: "greenhouse-install-script-generator",
		},
		{
			Name: "hakim",
			Path: "hakim",
		},
		{
			Name: "metron",
			Path: "loggregator/src/metron",
		},
		{
			Name: "diego-windows-release",
			Path: "DiegoWindowsRelease",
		},
	}
}
