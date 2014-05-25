package main

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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
		vPrintln("Extracting", f.Name)

		if !strings.HasSuffix(f.Name, "/") {
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
		} else {
			os.MkdirAll(filepath.Join(dest, f.Name), os.ModePerm)
		}
	}
	return nil
}

// exePath returns the executable path.
func exePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}
