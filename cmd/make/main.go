package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

type TemplateContent struct {
	SuiteStructName  string
	FunctionName     string
	TestFunctionName string
	PackageName      string
}

func main() {
	files := [][]string{
		{"../../src/easy/%s/%s.go", "templates/template.txt"},
		{"../../src/easy/%s/%s_test.go", "templates/template_test.txt"},
		{"../../src/easy/%s/suite_%s_test.go", "templates/template_suite_test.txt"},
	}

	packageName := os.Args[1]

	content := TemplateContent{
		SuiteStructName:  fmt.Sprintf("%sSuite", strcase.ToCamel(packageName)),
		FunctionName:     strcase.ToLowerCamel(packageName),
		TestFunctionName: strcase.ToCamel(packageName),
		PackageName:      packageName,
	}

	for _, file := range files {
		err := writeFile(file[0], file[1], content)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeFile(filePath string, templateFile string, content TemplateContent) error {
	tmp, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	fileComplete := fmt.Sprintf(filePath, content.PackageName, content.PackageName)

	split := strings.Split(fileComplete, "/")
	path := strings.Join(split[:len(split)-1], "/")
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(fileComplete)
	if err != nil {
		return err
	}

	err = tmp.Execute(file, content)

	return err
}
