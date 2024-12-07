package template

import (
	"fmt"
	"pg-backend/util"
)

type CodeSample = util.CodeSample

func GenerateBase(code *CodeSample) string {
	return fmt.Sprintf(`
%s

def main():
	print(add(10, 45))

if __name__ == "__main__":
	main()
	`, code.Content)
}