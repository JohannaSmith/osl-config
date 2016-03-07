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

func (c DiegoReleaseSubmodules) LtsPaths() []string {
	lts := c.Lts()
	p := make([]string, len(lts))
	for k, v := range lts {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

//Diego-release LTS is locked at a SHA. Refer to README.md for the actual SHA
func (c DiegoReleaseSubmodules) Lts() DiegoReleaseSubmodules {
	out := DiegoReleaseSubmodules{
		{
			Name: "diego-smoke-tests",
			Path: "src/github.com/cloudfoundry-incubator/diego-smoke-tests",
		},
		{
			Name: "auction",
			Path: "src/github.com/cloudfoundry-incubator/auction",
		},
		{
			Name: "auctioneer",
			Path: "src/github.com/cloudfoundry-incubator/auctioneer",
		},
		{
			Name: "bbs",
			Path: "src/github.com/cloudfoundry-incubator/bbs",
		},
		{
			Name: "benchmark-bbs",
			Path: "src/github.com/cloudfoundry-incubator/benchmark-bbs",
		},
		{
			Name: "file-server",
			Path: "src/github.com/cloudfoundry-incubator/file-server",
		},
		{
			Name: "executor",
			Path: "src/github.com/cloudfoundry-incubator/executor",
		},
		{
			Name: "stager",
			Path: "src/github.com/cloudfoundry-incubator/stager",
		},
		{
			Name: "rep",
			Path: "src/github.com/cloudfoundry-incubator/rep",
		},
		{
			Name: "nsync",
			Path: "src/github.com/cloudfoundry-incubator/nsync",
		},
		{
			Name: "cf-debug-server",
			Path: "src/github.com/cloudfoundry-incubator/cf-debug-server",
		},
		{
			Name: "converger",
			Path: "src/github.com/cloudfoundry-incubator/converger",
		},
		{
			Name: "route-emitter",
			Path: "src/github.com/cloudfoundry-incubator/route-emitter",
		},
		{
			Name: "runtime-schema",
			Path: "src/github.com/cloudfoundry-incubator/runtime-schema",
		},
		{
			Name: "locket",
			Path: "src/github.com/cloudfoundry-incubator/locket",
		},
		{
			Name: "tps",
			Path: "src/github.com/cloudfoundry-incubator/tps",
		},
		{
			Name: "cf_http",
			Path: "src/github.com/cloudfoundry-incubator/cf_http",
		},
		{
			Name: "buildpack_app_lifecycle",
			Path: "src/github.com/cloudfoundry-incubator/buildpack_app_lifecycle",
		},
		{
			Name: "docker_app_lifecycle",
			Path: "src/github.com/cloudfoundry-incubator/docker_app_lifecycle",
		},
		{
			Name: "diegocanaryapp",
			Path: "src/github.com/cloudfoundry-incubator/diegocanaryapp",
		},
		{
			Name: "consuladapter",
			Path: "src/github.com/cloudfoundry-incubator/consuladapter",
		},
		{
			Name: "diego-ssh",
			Path: "src/github.com/cloudfoundry-incubator/diego-ssh",
		},
		{
			Name: "cacheddownloader",
			Path: "src/github.com/cloudfoundry-incubator/cacheddownloader",
		},
		{
			Name: "healthcheck",
			Path: "src/github.com/cloudfoundry-incubator/healthcheck",
		},
		{
			Name: "cc-uploader",
			Path: "src/github.com/cloudfoundry-incubator/cc-uploader",
		},
		{
			Name: "vizzini",
			Path: "src/github.com/cloudfoundry-incubator/vizzini",
		},
	}
	return out
}

func NewDiegoReleaseSubmodules() DiegoReleaseSubmodules {
	return DiegoReleaseSubmodules{
		{
			Name: "auction",
			Path: "src/github.com/cloudfoundry-incubator/auction",
		},
		{
			Name: "auctioneer",
			Path: "src/github.com/cloudfoundry-incubator/auctioneer",
		},
		{
			Name: "bbs",
			Path: "src/github.com/cloudfoundry-incubator/bbs",
		},
		{
			Name: "benchmark-bbs",
			Path: "src/github.com/cloudfoundry-incubator/benchmark-bbs",
		},
		{
			Name: "file-server",
			Path: "src/github.com/cloudfoundry-incubator/file-server",
		},
		{
			Name: "executor",
			Path: "src/github.com/cloudfoundry-incubator/executor",
		},
		{
			Name: "stager",
			Path: "src/github.com/cloudfoundry-incubator/stager",
		},
		{
			Name: "rep",
			Path: "src/github.com/cloudfoundry-incubator/rep",
		},
		{
			Name: "nsync",
			Path: "src/github.com/cloudfoundry-incubator/nsync",
		},
		{
			Name: "cf-debug-server",
			Path: "src/github.com/cloudfoundry-incubator/cf-debug-server",
		},
		{
			Name: "converger",
			Path: "src/github.com/cloudfoundry-incubator/converger",
		},
		{
			Name: "route-emitter",
			Path: "src/github.com/cloudfoundry-incubator/route-emitter",
		},
		{
			Name: "routing-info",
			Path: "src/github.com/cloudfoundry-incubator/routing-info",
		},
		{
			Name: "runtime-schema",
			Path: "src/github.com/cloudfoundry-incubator/runtime-schema",
		},
		{
			Name: "locket",
			Path: "src/github.com/cloudfoundry-incubator/locket",
		},
		{
			Name: "tps",
			Path: "src/github.com/cloudfoundry-incubator/tps",
		},
		{
			Name: "cf_http",
			Path: "src/github.com/cloudfoundry-incubator/cf_http",
		},
		{
			Name: "buildpack_app_lifecycle",
			Path: "src/github.com/cloudfoundry-incubator/buildpack_app_lifecycle",
		},
		{
			Name: "docker_app_lifecycle",
			Path: "src/github.com/cloudfoundry-incubator/docker_app_lifecycle",
		},
		{
			Name: "diegocanaryapp",
			Path: "src/github.com/cloudfoundry-incubator/diegocanaryapp",
		},
		{
			Name: "consuladapter",
			Path: "src/github.com/cloudfoundry-incubator/consuladapter",
		},
		{
			Name: "diego-ssh",
			Path: "src/github.com/cloudfoundry-incubator/diego-ssh",
		},
		{
			Name: "cacheddownloader",
			Path: "src/github.com/cloudfoundry-incubator/cacheddownloader",
		},
		{
			Name: "healthcheck",
			Path: "src/github.com/cloudfoundry-incubator/healthcheck",
		},
		{
			Name: "cc-uploader",
			Path: "src/github.com/cloudfoundry-incubator/cc-uploader",
		},
		{
			Name: "vizzini",
			Path: "src/github.com/cloudfoundry-incubator/vizzini",
		},
	}
}
