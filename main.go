package main

import (
	"go-reloaded/core"
	"os"
)

func main() {
	path1 := os.Args[1]
	path2 := os.Args[2]
	resultfile := core.CreateFile(path2)
	text := string(core.GetFileContent(path1))
	resultfile.WriteString(core.ApplyChanges(text))
	resultfile.Close()
}
