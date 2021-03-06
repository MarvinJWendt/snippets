package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"
)

type ReadmeData struct {
	TotalSnippetCount      int
	GoSnippetCount         int
	MarkdownSnippetCount   int
	JavaScriptSnippetCount int

	SnippetTree string

	LastUpdateTime string
}

var RD ReadmeData

var ciPath string
var projectPath string

func main() {
	RD.LastUpdateTime = time.Now().UTC().Format(time.RFC822)
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
	var result int

	result = countSnippets("go")
	RD.TotalSnippetCount += result
	RD.GoSnippetCount = result

	result = countSnippets("markdown")
	RD.TotalSnippetCount += result
	RD.MarkdownSnippetCount = result

	result = countSnippets("javascript")
	RD.TotalSnippetCount += result
	RD.JavaScriptSnippetCount = result
}

func countSnippets(lang string) int {
	var result int
	var toc string
	toc += "\n## " + lang + "\n"
	check(filepath.Walk(filepath.Join(projectPath, "/"+lang), func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".") || strings.Contains(path, "_") || filepath.Base(path) == lang {
			return nil
		}
		result++
		_, f := filepath.Split(path)
		toc += fmt.Sprintf("- [%s](%s/%s)\n", f, lang, f)
		return nil
	}))
	RD.SnippetTree += toc
	return result
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
