package main

import (
	"html/template"
	//"os"
	"path/filepath"
)

var templates = make(map[string]*template.Template)

var tmplFuncs = template.FuncMap{
	"add": add,
}

func InitTemplates() error {
	files, err := filepath.Glob(filepath.Join("./html", "*.page.tmpl"))
	if err != nil {
		return err
	}
	for _, file := range files {
		name := filepath.Base(file)
		tmpl, err := template.New(name).Funcs(tmplFuncs).Parse(string(file))
		if err != nil {
			return err
		}
		templates[tmpl.Name()] = tmpl
	}
	return nil
}

func add(a, b, c int) int {
	return a + b + c
}
