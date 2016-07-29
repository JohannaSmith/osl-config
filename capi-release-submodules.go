package main

type CapiReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c CapiReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewCapiReleaseSubmodules() CapiReleaseSubmodules {
	return CapiReleaseSubmodules{
		{
			Name: "cloud_controller_ng",
			Path: "src/cloud_controller_ng",
		},
		{
			Name: "pat",
			Path: "src/github.com/bmizerany/pat",
		},
		{
			Name: "bbs",
			Path: "src/code.cloudfoundry.org/bbs",
		},
		{
			Name: "buildpackapplifecycle",
			Path: "src/code.cloudfoundry.org/buildpackapplifecycle",
		},
		{
			Name: "cc-uploader",
			Path: "src/code.cloudfoundry.org/cc-uploader",
		},
		{
			Name: "debugserver",
			Path: "src/code.cloudfoundry.org/debugserver",
		},
		{
			Name: "cflager",
			Path: "src/code.cloudfoundry.org/cflager",
		},
		{
			Name: "cfhttp",
			Path: "src/code.cloudfoundry.org/cfhttp",
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
			Name: "dockerapplifecycle",
			Path: "src/code.cloudfoundry.org/dockerapplifecycle",
		},
		{
			Name: "locket",
			Path: "src/github.com/cloudfoundry-incubator/locket",
		},
		{
			Name: "nsync",
			Path: "src/code.cloudfoundry.org/nsync",
		},
		{
			Name: "routing-info",
			Path: "src/github.com/cloudfoundry-incubator/routing-info",
		},
		{
			Name: "runtimeschema",
			Path: "src/code.cloudfoundry.org/runtimeschema",
		},
		{
			Name: "stager",
			Path: "src/code.cloudfoundry.org/stager",
		},
		{
			Name: "tps",
			Path: "src/code.cloudfoundry.org/tps",
		},
		{
			Name: "blobstore_url_signer",
			Path: "src/github.com/cloudfoundry/blobstore_url_signer",
		},
		{
			Name: "dropsonde",
			Path: "src/github.com/cloudfoundry/dropsonde",
		},
		{
			Name: "gosteno",
			Path: "src/github.com/cloudfoundry/gosteno",
		},
		{
			Name: "gunk",
			Path: "src/github.com/cloudfoundry/gunk",
		},
		{
			Name: "noaa",
			Path: "src/github.com/cloudfoundry/noaa",
		},
		{
			Name: "sonde-go",
			Path: "src/github.com/cloudfoundry/sonde-go",
		},
		{
			Name: "gogo-protobuf",
			Path: "src/github.com/gogo/protobuf",
		},
		{
			Name: "protobuf",
			Path: "src/github.com/golang/protobuf",
		},
		{
			Name: "websocket",
			Path: "src/github.com/gorilla/websocket",
		},
		{
			Name: "consul",
			Path: "src/github.com/hashicorp/consul",
		},
		{
			Name: "gouuid",
			Path: "src/github.com/nu7hatch/gouuid",
		},
		{
			Name: "ginkgo",
			Path: "src/github.com/onsi/ginkgo",
		},
		{
			Name: "gomega",
			Path: "src/github.com/onsi/gomega",
		},
		{
			Name: "clock",
			Path: "src/code.cloudfoundry.org/clock",
		},
		{
			Name: "lager",
			Path: "src/code.cloudfoundry.org/lager",
		},
		{
			Name: "ifrit",
			Path: "src/github.com/tedsuo/ifrit",
		},
		{
			Name: "rata",
			Path: "src/github.com/tedsuo/rata",
		},
		{
			Name: "go-sse",
			Path: "src/github.com/vito/go-sse",
		},
		{
			Name: "crypto",
			Path: "src/golang.org/x/crypto",
		},
	}
}
