package main

type PivotalAccountReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c PivotalAccountReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewPivotalAccountReleaseSubmodules() PivotalAccountReleaseSubmodules {
	return PivotalAccountReleaseSubmodules{
		{
			Name: "pivotal-account",
			Path: "src/pivotal-account",
		},
	}
}
