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
		{Name: "cloud_controller_ng", Path: "src/cloud_controller_ng"},
		{Name: "pat", Path: "src/github.com/bmizerany/pat"},
		{Name: "bbs", Path: "src/github.com/cloudfoundry-incubator/bbs"},
		{Name: "buildpack_app_lifecycle", Path: "src/github.com/cloudfoundry-incubator/buildpack_app_lifecycle"},
		{Name: "cc-uploader", Path: "src/github.com/cloudfoundry-incubator/cc-uploader"},
		{Name: "cf-debug-server", Path: "src/github.com/cloudfoundry-incubator/cf-debug-server"},
		{Name: "cf-lager", Path: "src/github.com/cloudfoundry-incubator/cf-lager"},
		{Name: "cf_http", Path: "src/github.com/cloudfoundry-incubator/cf_http"},
		{Name: "consuladapter", Path: "src/github.com/cloudfoundry-incubator/consuladapter"},
		{Name: "diego-ssh", Path: "src/github.com/cloudfoundry-incubator/diego-ssh"},
		{Name: "docker_app_lifecycle", Path: "src/github.com/cloudfoundry-incubator/docker_app_lifecycle"},
		{Name: "locket", Path: "src/github.com/cloudfoundry-incubator/locket"},
		{Name: "nsync", Path: "src/github.com/cloudfoundry-incubator/nsync"},
		{Name: "routing-info", Path: "src/github.com/cloudfoundry-incubator/routing-info"},
		{Name: "runtime-schema", Path: "src/github.com/cloudfoundry-incubator/runtime-schema"},
		{Name: "stager", Path: "src/github.com/cloudfoundry-incubator/stager"},
		{Name: "tps", Path: "src/github.com/cloudfoundry-incubator/tps"},
		{Name: "blobstore_url_signer", Path: "src/github.com/cloudfoundry/blobstore_url_signer"},
		{Name: "dropsonde", Path: "src/github.com/cloudfoundry/dropsonde"},
		{Name: "gosteno", Path: "src/github.com/cloudfoundry/gosteno"},
		{Name: "gunk", Path: "src/github.com/cloudfoundry/gunk"},
		{Name: "noaa", Path: "src/github.com/cloudfoundry/noaa"},
		{Name: "sonde-go", Path: "src/github.com/cloudfoundry/sonde-go"},
		{Name: "gogo-protobuf", Path: "src/github.com/gogo/protobuf"},
		{Name: "protobuf", Path: "src/github.com/golang/protobuf"},
		{Name: "websocket", Path: "src/github.com/gorilla/websocket"},
		{Name: "consul", Path: "src/github.com/hashicorp/consul"},
		{Name: "gouuid", Path: "src/github.com/nu7hatch/gouuid"},
		{Name: "ginkgo", Path: "src/github.com/onsi/ginkgo"},
		{Name: "gomega", Path: "src/github.com/onsi/gomega"},
		{Name: "clock", Path: "src/github.com/pivotal-golang/clock"},
		{Name: "lager", Path: "src/github.com/pivotal-golang/lager"},
		{Name: "ifrit", Path: "src/github.com/tedsuo/ifrit"},
		{Name: "rata", Path: "src/github.com/tedsuo/rata"},
		{Name: "go-sse", Path: "src/github.com/vito/go-sse"},
		{Name: "crypto", Path: "src/golang.org/x/crypto"},
	}
}
