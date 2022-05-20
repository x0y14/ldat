package ldat

import (
	"ldat/interfaces"
	"ldat/python"
	"strings"
)

type LDat struct {
	mainFile string
	libFiles []string
}

func NewLDat(mainFilePath string, libFilePaths []string) *LDat {
	return &LDat{
		mainFile: mainFilePath,
		libFiles: libFilePaths,
	}
}

// DumpAsts []string[0] main AST, []string[1:] lib ASTs
func (ld *LDat) DumpAsts() ([]string, error) {
	queue := []string{ld.mainFile}
	queue = append(queue, ld.libFiles...)

	var asts []string

	var psr interfaces.IParser

	for _, file := range queue {
		if strings.HasSuffix(file, ".py") {
			psr = &python.PyParser{}
			_ = psr.Parse()
		} else if strings.HasSuffix(file, ".rb") {

		}
	}

	//if strings.HasSuffix(ld.mainFile, ".py") {
	//	result, err := python.AstDumpWithFilePath(ld.mainFile)
	//	if err != nil {
	//		return nil, err
	//	}
	//	fmt.Printf("%v", result)
	//} else if strings.HasSuffix(ld.mainFile, ".rb") {
	//	result, err := ruby.AstDumpWithFilePath(ld.mainFile)
	//	if err != nil {
	//		return nil, err
	//	}
	//	fmt.Printf("%v", result)
	//}

	return asts, nil
}
