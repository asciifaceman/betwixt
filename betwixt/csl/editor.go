package csl

import (
	"os"
	"os/exec"
)

// Open given filename with given editor
// Note: VSCode doesn't work as editor in WSL+Windows
func Open(editor string, filename string) error {
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
