package main

type EtcdReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c EtcdReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewEtcdReleaseSubmodules() EtcdReleaseSubmodules {
	return EtcdReleaseSubmodules{
		{
			Name: "etcd",
			Path: "src/github.com/coreos/etcd",
		},
		{
			Name: "etcd-metrics-server",
			Path: "src/etcd-metrics-server",
		},
	}
}
