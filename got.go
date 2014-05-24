package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const APP_VER = "0.2.0524"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var (
	verbore = flag.Bool("v", false, "show the detail info")
	repoDir = "/usr/local/got"
)

func vPrintln(args ...interface{}) {
	if *verbore {
		fmt.Print("[Info] ")
		fmt.Println(args...)
	}
}

func vPrintf(format string, args ...interface{}) {
	if *verbore {
		fmt.Printf("[Info] "+format, args...)
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
	   download	download only
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

	os.MkdirAll(repoDir, os.ModePerm)

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
