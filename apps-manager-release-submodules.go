package main

type AppsManagerReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c AppsManagerReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewAppsManagerReleaseSubmodules() AppsManagerReleaseSubmodules {
	return AppsManagerReleaseSubmodules{
		{
			Name: "invitations",
			Path: "src/invitations",
		},
		{
			Name: "app-usage-service",
			Path: "src/app-usage-service",
		},
		{
			Name: "apps-manager-js",
			Path: "src/apps-manager-js",
		},
	}
}
