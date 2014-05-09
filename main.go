package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const APP_VER = "0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var (
	verbore   = flag.Bool("v", false, "show the detail info")
	isRemove  = flag.Bool("r", false, "remove a binary")
	isUpdate  = flag.Bool("u", false, "update a binary")
	isUpgrade = flag.Bool("p", false, "upgrade self")
)

var downTempl = "http://gobuild.io/%s/%s/%s/%s"

func vPrintln(args ...interface{}) {
	if *verbore {
		fmt.Println(args...)
	}
}

func vPrintf(format string, args ...interface{}) {
	if *verbore {
		fmt.Printf(format, args...)
	}
}

func getPath() string {
	path := os.Getenv("PATH")
	gPath := os.Getenv("GOPATH")
	if gPath != "" {
		bin := filepath.Join(gPath, "bin")
		if strings.Contains(path, bin) {
			return bin
		}
	}

	rootPath := os.Getenv("GOROOT")
	if rootPath != "" {
		bin := filepath.Join(rootPath, "bin")
		if strings.Contains(path, bin) {
			return bin
		}
	}

	//TODO: where is the default binary path?
	return "/usr/local/bin"
}

func findPara(start int) int {
	if start >= len(os.Args) {
		return -1
	}

	for i := start; i < len(os.Args); i++ {
		if !strings.HasPrefix(os.Args[i], "-") {
			return i
		}
	}
	return -1
}

func unzip(dest, fPath string) error {
	r, err := zip.OpenReader(fPath)
	if err != nil {
		return err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		newPath := filepath.Join(dest, f.Name)
		os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
		nf, err := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer nf.Close()

		_, err = io.Copy(nf, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func download() {

}

func install() {
	idx := findPara(1)
	if idx == -1 {
		fmt.Println("not indicate a go package path")
		return
	}

	var ver, pkgPath = "master", os.Args[idx]

	idx = findPara(idx + 1)
	if idx != -1 {
		ver = os.Args[idx]
	}

	ss := strings.Split(pkgPath, "/")
	if len(ss) == 2 {
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
	err = os.Rename(src, dst)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.RemoveAll(extractDir)
	os.Remove(binPath)
}

func remove() {

}

func update() {

}

func upgrade() {

}

func main() {
	flag.Parse()
	if *isRemove {
		remove()
	} else if *isUpdate {
		update()
	} else if *isUpgrade {
		upgrade()
	} else {
		install()
	}
}
