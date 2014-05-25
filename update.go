package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

var (
	pkgListUrl = "http://beta.gobuild.io/api/pkglist?os=%s&arch=%s"
)

type Branch struct {
	Name        string
	Sha         string
	Updated     string
	Os          string
	Arch        string
	Zipball_url string
}

type Pkg struct {
	Name        string
	Description string
	Branches    []Branch
}

func pkgUrl() string {
	return fmt.Sprintf(pkgListUrl, runtime.GOOS, runtime.GOARCH)
}

func update() {
	resp, err := http.Get(pkgUrl())
	if err != nil {
		fmt.Println(err)
		return
	}

	var pkgs = make([]Pkg, 0)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&pkgs)
	if err != nil {
		fmt.Println(err)
		return
	}
}
