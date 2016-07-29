package main

type DiegoReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c DiegoReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewDiegoReleaseSubmodules() DiegoReleaseSubmodules {
	return DiegoReleaseSubmodules{
		{
			Name: "auction",
			Path: "src/code.cloudfoundry.org/auction",
		},
		{
			Name: "auctioneer",
			Path: "src/code.cloudfoundry.org/auctioneer",
		},
		{
			Name: "bbs",
			Path: "src/code.cloudfoundry.org/bbs",
		},
		{
			Name: "benchmark-bbs",
			Path: "src/code.cloudfoundry.org/benchmarkbbs",
		},
		{
			Name: "file-server",
			Path: "src/code.cloudfoundry.org/fileserver",
		},
		{
			Name: "executor",
			Path: "src/code.cloudfoundry.org/executor",
		},
		{
			Name: "rep",
			Path: "src/code.cloudfoundry.org/rep",
		},
		{
			Name: "cf-debug-server",
			Path: "src/code.cloudfoundry.org/debugserver",
		},
		{
			Name: "route-emitter",
			Path: "src/code.cloudfoundry.org/route-emitter",
		},
		{
			Name: "routing-info",
			Path: "src/github.com/cloudfoundry-incubator/routing-info",
		},
		{
			Name: "runtime-schema",
			Path: "src/code.cloudfoundry.org/runtimeschema",
		},
		{
			Name: "locket",
			Path: "src/code.cloudfoundry.org/locket",
		},
		{
			Name: "cf_http",
			Path: "src/code.cloudfoundry.org/cfhttp",
		},
		{
			Name: "buildpack_app_lifecycle",
			Path: "src/code.cloudfoundry.org/buildpackapplifecycle",
		},
		{
			Name: "docker_app_lifecycle",
			Path: "src/code.cloudfoundry.org/dockerapplifecycle",
		},
		{
			Name: "diegocanaryapp",
			Path: "src/code.cloudfoundry.org/diegocanaryapp",
		},
		{
			Name: "consuladapter",
			Path: "src/code.cloudfoundry.org/consuladapter",
		},
		{
			Name: "diego-ssh",
			Path: "src/code.cloudfoundry.org/diego-ssh",
		},
		{
			Name: "cacheddownloader",
			Path: "src/code.cloudfoundry.org/cacheddownloader",
		},
		{
			Name: "healthcheck",
			Path: "src/code.cloudfoundry.org/healthcheck",
		},
		{
			Name: "vizzini",
			Path: "src/code.cloudfoundry.org/vizzini",
		},
	}
}
