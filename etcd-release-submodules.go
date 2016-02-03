package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type EtcdReleaseSubmodules []EtcdReleaseSubmodule
type EtcdReleaseSubmodule struct {
	Name string // the name of the project
	Path string // the path of the project in the latest release
}

func (c EtcdReleaseSubmodules) Paths() []string {
	p := make([]string, len(c))
	for k, v := range c {
		p[k] = v.Path
	}
	return p
}

func NewEtcdReleaseSubmodules() EtcdReleaseSubmodules {
	tmpDir, _ := ioutil.TempDir(os.TempDir(), "")
	defer os.RemoveAll(tmpDir)
	_, err := exec.Command("git", "clone", "https://github.com/cloudfoundry-incubator/etcd-release", tmpDir).Output()
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("git", "-C", tmpDir, "submodule").Output()
	if err != nil {
		log.Fatal(err)
	}
	list := make(EtcdReleaseSubmodules, 0)
	for _, v := range strings.Split(string(out), "\n") {
		if v == "" {
			continue
		}
		module := strings.Split(v, " ")
		list = append(list, EtcdReleaseSubmodule{
			Name: path.Base(module[1]),
			Path: module[1],
		})
	}
	return list
}
