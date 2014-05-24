package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func link(binName, src string) error {
	bPath := getPath()
	vPrintln("link", src, "to", filepath.Join(bPath, binName))
	return makeLink(src, filepath.Join(bPath, binName))
}

func install() {
	idx := findPara(1)
	if idx == -1 {
		fmt.Println("need a go package path")
		return
	}

	if os.Args[idx] == "install" {
		idx = findPara(idx + 1)
	}
	if idx == -1 {
		fmt.Println("need a go package path")
		return
	}

	var ver, pkgPath = "master", os.Args[idx]

	idx = findPara(idx + 1)
	if idx != -1 {
		ver = os.Args[idx]
	}

	ss := strings.Split(pkgPath, "/")
	if len(ss) < 2 {
		fmt.Println("not a go package path")
		return
	}
	if !strings.Contains(ss[0], ".") {
		pkgPath = "github.com/" + pkgPath
	}

	err := get(ver, pkgPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = link(binName(pkgPath), binPath(ver, pkgPath))
	if err != nil {
		fmt.Println(err)
		return
	}
}
