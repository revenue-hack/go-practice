package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Package struct {
	Name       string
	ImportPath string
	Deps       []string
}

func main() {
	in := os.Args[1:]
	pList, err := goList(in)
	if err != nil {
		log.Printf("%v\n", err)
	}
	pListDeps, err := goList([]string{"..."})
	printPackageList(pList, pListDeps)
}

func printPackageList(pList, pListDeps []*Package) {
	fmt.Println("input packages")
	var inputPackageList []string
	for _, p := range pList {
		fmt.Printf("%s ", p.ImportPath)
		inputPackageList = append(inputPackageList, p.ImportPath)
	}
	fmt.Println()
	fmt.Println("dependent packages")
	for _, p := range pListDeps {
		if !inDepend(inputPackageList, p) {
			continue
		}
		fmt.Printf("%s ", p.ImportPath)
	}
	fmt.Println()
}

func inDepend(array []string, p *Package) bool {
loop:
	for _, v := range array {
		for _, pv := range p.Deps {
			if pv == v {
				continue loop
			}
		}
		return false
	}
	return true
}

func goList(packages []string) ([]*Package, error) {
	var args = []string{"list", "-e", "-json"}
	args = append(args, packages...)
	cmd := exec.Command("go", args...)
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdOut.Close()
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
	}()
	jsonDecoder := json.NewDecoder(stdOut)
	var packageList []*Package
	for {
		var p Package
		err := jsonDecoder.Decode(&p)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			return packageList, nil
		}
		packageList = append(packageList, &p)
	}
}
