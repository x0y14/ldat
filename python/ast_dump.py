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
