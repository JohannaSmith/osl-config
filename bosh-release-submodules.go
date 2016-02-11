package main

type BoshReleaseSubmodules []struct {
	Name     string // the name of the project
	Path     string
	RepoName string
	RepoURL  string
}

func NewBoshReleaseSubmodules() BoshReleaseSubmodules {
	return BoshReleaseSubmodules{
		{
			Name:     "bosh-aws-cpi",
			RepoName: "bosh-aws-cpi",
			Path:     "src/bosh_aws_cpi",
			RepoURL:  "https://github.com/cloudfoundry-incubator/bosh-aws-cpi-release",
		},
		{
			Name:     "bosh-vcloud-cpi",
			RepoName: "bosh-vcloud-cpi",
			Path:     "src/bosh_vcloud_cpi",
			RepoURL:  "https://github.com/cloudfoundry-incubator/bosh-vcloud-cpi-release",
		},
		{
			Name:     "bosh-vsphere-cpi",
			RepoName: "bosh-vsphere-cpi",
			Path:     "src/vsphere_cpi",
			RepoURL:  "https://github.com/cloudfoundry-incubator/bosh-vsphere-cpi-release",
		},
		{
			Name:     "bosh",
			RepoName: "bosh",
			Path:     ".",
			RepoURL:  "https://github.com/cloudfoundry/bosh",
		},
		{
			Name:     "bosh-davcli",
			RepoName: "bosh-davcli",
			Path:     ".",
			RepoURL:  "https://github.com/cloudfoundry/bosh-davcli",
		},
	}
}
