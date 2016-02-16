package main

type BoshCpiReleasesSubmodules []struct {
	Name     string // the name of the project
	Path     string
	RepoName string
	RepoURL  string
}

func NewBoshCpiReleasesSubmodules() BoshCpiReleasesSubmodules {
	return BoshCpiReleasesSubmodules{
		{
			Name:     "bosh-aws-cpi",
			Path:     "src/bosh_aws_cpi",
			RepoName: "bosh-aws-cpi",
			RepoURL:  "https://github.com/cloudfoundry-incubator/bosh-aws-cpi-release",
		},
		{
			Name:     "bosh-vcloud-cpi",
			Path:     "src/bosh_vcloud_cpi",
			RepoName: "bosh-vcloud-cpi",
			RepoURL:  "https://github.com/cloudfoundry-incubator/bosh-vcloud-cpi-release",
		},
		{
			Name:     "bosh-vsphere-cpi",
			Path:     "src/vsphere_cpi",
			RepoName: "bosh-vsphere-cpi",
			RepoURL:  "https://github.com/cloudfoundry-incubator/bosh-vsphere-cpi-release",
		},
	}
}
