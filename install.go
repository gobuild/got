package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

var (
	tmplProxy = `package main

import (
	"os/exec"
	"os"
)

func main() {
	cmd := exec.Command("{{.Path}}", os.Args[1:]...)
	cmd.Dir = "{{.Dir}}"
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
`
)

func proxy(binName, src string) error {
	bPath := getPath()
	vPrintln("Build proxy from", src, "to", filepath.Join(bPath, binName))

	f := filepath.Join(os.TempDir(), "got", binName)
	fmt.Println("1", f)
	os.MkdirAll(f, os.ModePerm)
	m, err := os.Create(filepath.Join(f, "main.go"))
	if err != nil {
		return err
	}

	t, err := template.New("t").Parse(tmplProxy)
	if err != nil {
		m.Close()
		return err
	}
	data := struct {
		Path string
		Dir  string
	}{
		src,
		filepath.Dir(src),
	}
	err = t.Execute(m, data)
	if err != nil {
		m.Close()
		return err
	}
	m.Close()
	cmdBuild := exec.Command("go", "build")
	cmdBuild.Dir = f
	cmdBuild.Stdin = os.Stdin
	cmdBuild.Stdout = os.Stdout
	err = cmdBuild.Run()
	if err != nil {
		return err
	}
	return os.Rename(filepath.Join(f, binName), filepath.Join(bPath, binName))
}

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

	//if true {
	if runtime.GOOS == "windows" {
		err = proxy(binName(pkgPath), binPath(ver, pkgPath))
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		binName := binName(pkgPath)
		/*dstPath := filepath.Join(bPath, binName)
		stat, err := os.Stat(dstPath)
		if err != nil {
			fmt.Println(err)
			return
		}*/

		err = link(binName, binPath(ver, pkgPath))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
