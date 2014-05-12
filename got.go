package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const APP_VER = "0.1.0512"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var (
	verbore = flag.Bool("v", false, "show the detail info")
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

func help() {
	/*fmt.Printf(`NAME:
	   got - Go Binary Package Manager

	USAGE:
	   got [global options] [command] [package path] [arguments...]

	VERSION:
	   %s

	COMMANDS:
	   list		list all packages
	   install	download and install binary of package
	   run		install and run the command
	   down		download only
	   rm		remove package
	   update	update one package
	   upgrade	update self
	   help		Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   -v	print the version
	`, APP_VER)*/
	fmt.Printf(`NAME:
   got - Go Binary Package Manager

USAGE:
   got [global options] [command] [package path] [arguments...]

VERSION:
   %s

COMMANDS:
   list		list all packages
   install	download and install binary of package
   help		Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -v	print the version
`, APP_VER)
}

// exePath returns the executable path.
func exePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
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

	p, err := exePath()
	if err == nil {
		return filepath.Dir(p)
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
		vPrintf("Extracting %s:\n", f.Name)
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

func remove() {

}

func update() {

}

func upgrade() {

}

func search() {

}

func main() {
	flag.Parse()

	l := len(os.Args)
	if l == 1 {
		help()
		return
	}

	idx := findPara(1)
	if idx == -1 {
		help()
		return
	}

	switch strings.ToLower(os.Args[idx]) {
	case "list":
		list()
	case "help":
		help()
	case "rm":
		remove()
	case "download":
		download()
	case "search":
		search()
	case "upgrade":
		upgrade()
	case "update":
		update()
	case "install":
		install()
	default:
		install()
	}
}
