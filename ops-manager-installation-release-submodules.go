package main

type OpsManagerInstallationReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c OpsManagerInstallationReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Path
	}
	return p
}

func NewOpsManagerInstallationReleaseSubmodules() OpsManagerInstallationReleaseSubmodules {
	return OpsManagerInstallationReleaseSubmodules{
		{
			Name: "web",
			Path: "web",
		},
	}
}
