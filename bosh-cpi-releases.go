package main

const ReleaseName = "bosh-cpi-releases"

func NewReports() Reports {
	return Reports{
		{
			Name:    "bosh-aws-cpi",
			RepoURL: "https://github.com/cloudfoundry-incubator/bosh-aws-cpi-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "bosh-aws-cpi",
					Path: "src/bosh_aws_cpi",
				},
			},
		},
		{
			Name:    "bosh-vcloud-cpi",
			RepoURL: "https://github.com/cloudfoundry-incubator/bosh-vcloud-cpi-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "bosh-vcloud-cpi",
					Path: "src/bosh_vcloud_cpi",
				},
			},
		},
		{
			Name:    "bosh-vsphere-cpi",
			RepoURL: "https://github.com/cloudfoundry-incubator/bosh-vsphere-cpi-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "bosh-vsphere-cpi",
					Path: "src/vsphere_cpi",
				},
			},
		},
	}
}
