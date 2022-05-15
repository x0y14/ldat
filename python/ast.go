package python

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// ef. ast_dump.py
const astDumpPy = `
import ast
import sys

if __name__ == '__main__':

    args = sys.argv

    if len(args) < 2:
        # not found : analyze target (*.py)
        print("not found : analyze target (*.py)")
        exit(1)

    target = args[1]

    # is python file? (ended with ".py"?)
    if target[-3:] != ".py":
        print("target suffix is not '.py'")
        exit(1)

    with open(args[1], "r") as f:
        tree = ast.parse(f.read())
        print(ast.dump(tree))

    exit(0)

`

func AstDump(filepath string) (string, error) {
	tmpF, _ := ioutil.TempFile("", "astDumpPy")
	defer os.Remove(tmpF.Name())

	_, err := tmpF.WriteString(astDumpPy)
	if err != nil {
		return "", err
	}

	out, err := exec.Command("python", tmpF.Name(), filepath).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
