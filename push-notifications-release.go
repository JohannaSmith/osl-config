package main

const ReleaseName = "push-notifications-release"

func NewReports() Reports {
	return Reports{
		{
			Name:    "push-api",
			RepoURL: "git@github.com:cfmobile/push.git",
			Branch:  "dev",
			Paths: ReportPaths{
				{
					Name:           "push-api",
					Path:           "push",
					PrepareCommand: "pushd .. && mvn clean install -Pprod -DskipTests && popd",
				},
			},
		},
		{
			Name:    "push-dashboard",
			RepoURL: "git@github.com:cfmobile/push-dashboard.git",
			Branch:  "dev",
			Paths: ReportPaths{
				{
					Name: "push-dashboard",
					Path: ".",
				},
			},
		},
		{
			Name:    "push-scheduler",
			RepoURL: "git@github.com:cfmobile/push-scheduler.git",
			Branch:  "dev",
			Paths: ReportPaths{
				{
					Name: "push-scheduler",
					Path: ".",
				},
			},
		},
		{
			Name:    "push-analytics",
			RepoURL: "git@github.com:cfmobile/push-analytics.git",
			Branch:  "dev",
			Paths: ReportPaths{
				{
					Name:           "push-analytics",
					Path:           "push-analytics",
					PrepareCommand: "pushd .. && mvn clean install -DskipTests && popd",
				},
			},
		},
		{
			Name:    "push-smoke-tests",
			RepoURL: "git@github.com:cfmobile/push-smoke-tests.git",
			Branch:  "develop",
			Paths: ReportPaths{
				{
					Name: "push-smoke-tests",
					Path: ".",
				},
			},
		},
	}
}
