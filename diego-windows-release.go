package main

const ReleaseName = "diego-windows-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "diego-windows",
			RepoURL: "https://github.com/cloudfoundry/diego-windows-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "greenhouse-install-script-generator",
					Path: "greenhouse-install-script-generator",
				},
				{
					Name: "hakim",
					Path: "hakim",
				},
				{
					Name: "metron",
					Path: "loggregator/src/metron",
				},
				{
					Name: "DiegoWindowsRelease",
					Path: "DiegoWindowsRelease",
				},
			},
		},
	}
}
