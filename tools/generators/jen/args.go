package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/june-style/go-sample/domain/derrors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewArgs() (Args, error) {
	flag.Parse()
	args := flag.Args()

	if len(args) != argsIndexCount {
		return nil, derrors.NewInvalidArgument("invalid file path")
	}

	flame := derrors.Caller(0)

	if flame.Dir() == "" {
		return nil, derrors.NewInvalidArgument("invalid base directory")
	}

	return &argsImpl{
		basePath:    flame.Dir(),
		filePath:    strings.TrimRight(strings.ToLower(args[argsIndexPath]), ".go"),
		packageName: "main",
	}, nil
}

type Args interface {
	FullPath() string
	PackageName() string
	MethodName() string
}

const (
	argsIndexPath = iota
	argsIndexCount
)

type argsImpl struct {
	basePath    string
	filePath    string
	packageName string
}

func (a *argsImpl) FullPath() string {
	return fmt.Sprintf("%s/%s_gen.go", a.basePath, a.filePath)
}

func (a *argsImpl) PackageName() string {
	return a.packageName
}

func (a *argsImpl) MethodName() string {
	return cases.Title(language.Und).String(a.filePath)
}
