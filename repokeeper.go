package main

import (
	"fmt"
	"os"
)

// Refactor : Return Package Object instead string and bool
func findPackage(name string) (Package, bool) {
	repo := getLocalRepo()
	found := false
	fpath := ""
	for _, path := range repo {
		found = isExist(path+"/"+name) && isExist(path+"/"+name+"/dist.toml")
		if found {
			fpath = (path + "/" + name)
			return tomlToPackage(fpath + "/dist.toml"), found
		}
	}
	fmt.Println("Package not found: ", name)
	os.Exit(1)
	p := Package{}
	return p, false
}

func getPackage(name string) {
	p, _ := findPackage(name)
	fmt.Println("[*] Fetching resource ...")
	p.Download()
}

func installPackage(name string) {
	p, found := findPackage(name)
	if !found {
		os.Exit(1)
	}
	p.Download()
}

func extractPackage(name string) {
	p, found := findPackage(name)
	if !found {
		os.Exit(1)
	}
	fmt.Println("[", name, "] Extarcting resources...")
	p.extract("")
}
func buildPackage(name string) {

}
func removePackage(name string) {
}
