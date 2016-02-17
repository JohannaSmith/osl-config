package main

type NotificationsUiReleaseSubmodules []struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c NotificationsUiReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewNotificationsUiReleaseSubmodules() NotificationsUiReleaseSubmodules {
	return NotificationsUiReleaseSubmodules{
		{
			Name: "notifications-ui",
			Path: "src/notifications-ui",
		},
	}
}
