package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func run() {
	if len(os.Args) < 3 {
		fmt.Println("no param")
		return
	}

	p, err := exec.LookPath(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Readlink(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	var args []string
	if len(os.Args) > 3 {
		args = os.Args[3:]
	} else {
		args = make([]string, 0)
	}

	cmd := exec.Command(f, args...)
	cmd.Dir = filepath.Dir(f)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
