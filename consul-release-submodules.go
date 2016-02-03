package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type ConsulReleaseSubmodules []ConsulReleaseSubmodule
type ConsulReleaseSubmodule struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c ConsulReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Path
	}
	return p
}

func NewConsulReleaseSubmodules() ConsulReleaseSubmodules {
	tmpDir, _ := ioutil.TempDir(os.TempDir(), "")
	defer os.RemoveAll(tmpDir)
	_, err := exec.Command("git", "clone", "https://github.com/cloudfoundry-incubator/consul-release", tmpDir).Output()
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("git", "-C", tmpDir, "submodule").Output()
	if err != nil {
		log.Fatal(err)
	}
	list := make(ConsulReleaseSubmodules, 0)
	for _, v := range strings.Split(string(out), "\n") {
		if v == "" || strings.Contains(v, "acceptance-test") {
			continue
		}
		module := strings.Split(v, " ")
		list = append(list, ConsulReleaseSubmodule{
			Name: path.Base(module[1]),
			Path: module[1],
		})
	}
	return list
}
