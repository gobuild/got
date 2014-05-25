package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	downTempl = "http://gobuild.io/%s/%s/%s/%s"
)

func url(ver, pkgPath string) string {
	return fmt.Sprintf(downTempl, pkgPath, ver, runtime.GOOS, runtime.GOARCH)
}

func download() {
	idx := findPara(1)
	if idx == -1 {
		fmt.Println("need a param")
		return
	}

	if os.Args[idx] == "download" {
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
	}
}

func binName(pkg string) string {
	ss := strings.Split(pkg, "/")
	return ss[len(ss)-1]
}

func binPath(ver, pkg string) string {
	ss := strings.Split(pkg, "/")
	ps := append([]string{repoDir}, ss[:len(ss)-1]...)
	dir, bin := filepath.Join(ps...), ss[len(ss)-1]
	return filepath.Join(dir, bin, ver, ss[len(ss)-1])
}

func get(ver, pkg string) error {
	u := url(ver, pkg)
	vPrintln("Request", u)
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprint("error status code:", resp.StatusCode))
	}

	ss := strings.Split(pkg, "/")
	ps := append([]string{repoDir}, ss[:len(ss)-1]...)
	dir, bin := filepath.Join(ps...), ss[len(ss)-1]
	os.MkdirAll(dir, os.ModePerm)
	binPath := filepath.Join(dir, bin+"-"+ver+".zip")
	vPrintln("Save to", binPath)
	f, err := os.OpenFile(binPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	extractDir := filepath.Join(dir, bin, ver)
	os.MkdirAll(extractDir, os.ModePerm)
	vPrintln("Extract to", extractDir)
	err = unzip(extractDir, binPath)
	if err != nil {
		return err
	}

	vPrintln("Clear", binPath)
	os.Remove(binPath)
	return nil
}
