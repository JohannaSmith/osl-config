package main

type NotificationsReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c NotificationsReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewNotificationsReleaseSubmodules() NotificationsReleaseSubmodules {
	return NotificationsReleaseSubmodules{
		{
			Name: "notifications",
			Path: "src/notifications",
		},
	}
}
