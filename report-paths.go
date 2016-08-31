package main

type ReportPaths []struct {
	Name           string
	Path           string
	PrepareCommand string
}

func (r ReportPaths) ToString() []string {
	p := make([]string, len(r))
	for k, v := range r {
		p[k] = v.Name + ":" + v.Path
	}
	return p
}

func NewReportPaths() ReportPaths {
	return ReportPaths{}
}

func (r Reports) ReleaseName() string {
	return ReleaseName
}

type Reports []struct {
	Name    string
	RepoURL string
	Branch  string
	Paths   ReportPaths
}
