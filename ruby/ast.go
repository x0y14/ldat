package ruby

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// ef. ast_dump.rb
const astDumpRb = `
if ARGV.size == 0 then
    puts "not found : analyze target (*.rb)"
    exit(1)
end

#if ARGV[0].end_with?(".rb") == false then
#    puts "target suffix is not '.rb'"
#    exit(1)
#end

pp RubyVM::AbstractSyntaxTree.parse_file(ARGV[0])
exit(0)

`

func AstDumpWithFilePath(filepath string) (string, error) {
	// make create tmp file
	tmpF, _ := ioutil.TempFile("", "astDumpRb")
	defer os.Remove(tmpF.Name())

	// write file content
	_, err := tmpF.WriteString(astDumpRb)
	if err != nil {
		return "", err
	}

	// dump ast
	out, err := exec.Command("ruby", tmpF.Name(), filepath).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
