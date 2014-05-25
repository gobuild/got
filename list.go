package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func list() {
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

	for _, pkg := range pkgs {
		fmt.Println(pkg.Name, "\t", pkg.Description)
	}
}
