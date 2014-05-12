package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

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

	url := fmt.Sprintf(downTempl, pkgPath, ver, runtime.GOOS, runtime.GOARCH)
	vPrintln("getting from", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("error status code:", resp.StatusCode)
		return
	}

	dir, bin := getPath(), ss[len(ss)-1]
	binPath := filepath.Join(dir, bin+".zip")
	vPrintln("writting to", binPath)
	f, err := os.OpenFile(binPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	extractDir := filepath.Join(dir, "temp")
	os.MkdirAll(extractDir, os.ModePerm)
	vPrintln("unzip to", extractDir)
	err = unzip(extractDir, binPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	dst := filepath.Join(dir, bin)
	src := filepath.Join(extractDir, bin)
	vPrintln("moving", src, "to", dst)
	if runtime.GOOS == "windows" {
		err = os.Rename(src+".exe", dst+".exe")
	} else {
		err = os.Rename(src, dst)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	vPrintln("clear", extractDir)
	os.RemoveAll(extractDir)

	vPrintln("clear", binPath)
	os.Remove(binPath)
}
