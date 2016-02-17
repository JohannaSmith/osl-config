package main

type BoshReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c BoshReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewBoshReleaseSubmodules() BoshReleaseSubmodules {
	return BoshReleaseSubmodules{
		{
			Name: "bosh",
			Path: ".",
		},
		{
			Name: "bosh-davcli",
			Path: "go/src/github.com/cloudfoundry/bosh-davcli",
		},
		{
			Name: "bosh-agent",
			Path: "go/src/github.com/cloudfoundry/bosh-agent",
		},
	}
}
