// athena-dbg is a simple wrapper around godbg with sensible defaults to run athena.py
// under the gdb debugger.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	_ "github.com/sirnewton01/godbg"
)

func main() {
	fmt.Printf(":: running athena-dbg with args: %v\n", os.Args[1:])

	pyexe, err := exec.LookPath("python")
	if err != nil {
		fmt.Fprintf(os.Stderr, "**error** looking for python executable: %v\n", err)
		os.Exit(1)
	}

	athena, err := exec.LookPath("athena.py")
	if err != nil {
		fmt.Fprintf(os.Stderr, "**error** looking for athena.py script: %v\n", err)
		os.Exit(1)
	}

	godbg, err := exec.LookPath("godbg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "**error** looking for godbg executable: %v\n", err)
		os.Exit(1)
	}

	args := []string{pyexe, athena}
	args = append(args, os.Args[1:]...)
	cmd := exec.Command(godbg, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "**error** running godbg: %v\n", err)
		os.Exit(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
	}
}

// EOF
