package ldat

import (
	"fmt"
	"ldat/python"
	"ldat/ruby"
	"strings"
)

type LDat struct {
	mainFile string
	libFiles []string
}

func NewLDat(m string, l []string) *LDat {
	return &LDat{
		mainFile: m,
		libFiles: l,
	}
}

func (ld *LDat) Load() error {
	if strings.HasSuffix(ld.mainFile, ".py") {
		result, err := python.AstDump(ld.mainFile)
		if err != nil {
			return err
		}
		fmt.Printf("%v", result)
	} else if strings.HasSuffix(ld.mainFile, ".rb") {
		result, err := ruby.AstDump(ld.mainFile)
		if err != nil {
			return err
		}
		fmt.Printf("%v", result)
	}

	return nil
}
