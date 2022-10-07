package main

import (
	"embed"
	"html/template"
	"os"
)

//go:embed *.html
var files embed.FS

func main() {
	tmpl, _ := template.ParseFS(files, "home.html")
	tmpl.Execute(os.Stdout, map[string]string{"name": "saurabh"})
}
