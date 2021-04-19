package main

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

type ReadmeData struct {
	TotalSnippetCount int
	GoSnippetCount    int
}

var RD ReadmeData

var ciPath string
var projectPath string

func main() {
	_, scriptPath, _, _ := runtime.Caller(0)
	ciPath = filepath.Join(scriptPath, "../")
	projectPath = filepath.Join(ciPath, "../")

	fillReadmeData()

	tmpl, err := template.ParseFiles(filepath.Join(ciPath, "/README.template.md"))
	check(err)

	var tpl bytes.Buffer
	check(tmpl.Execute(&tpl, RD))

	check(ioutil.WriteFile(filepath.Join(ciPath, "../README.md"), tpl.Bytes(), 0600))
}

func fillReadmeData() {
	countGoSnippets()
}

func countGoSnippets() {
	var result int
	check(filepath.Walk(filepath.Join(projectPath, "/go"), func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".") {
			return nil
		}

		result++

		return nil
	}))
	RD.TotalSnippetCount += result
	RD.GoSnippetCount = result
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
