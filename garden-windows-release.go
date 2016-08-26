package main

const ReleaseName = "garden-windows-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "garden-windows",
			RepoURL: "https://github.com/cloudfoundry/garden-windows-release",
			Branch:  "master",
			Paths: ReportPaths{
				{
					Name: "GardenWindowsRelease",
					Path: "GardenWindowsRelease",
				},
				{
					Name: "Containerizer",
					Path: "src/github.com/cloudfoundry/garden-windows/Containerizer/",
				},
				{
					Name: "garden-windows",
					Path: "src/github.com/cloudfoundry/garden-windows/",
				},
			},
		},
	}
}
