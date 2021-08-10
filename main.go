package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func runWith(shell string) {
	if path, err := exec.LookPath(shell); err == nil {
		var out bytes.Buffer

		fmt.Printf("%s:\n", path)
		cmd := exec.Command(path, `-c`, `echo '\\'`)
		cmd.Stdout = &out
		cmd.Stderr = &out

		_ = cmd.Run()
		fmt.Printf("%s\n", out.String())
	} else {
		fmt.Printf("%s does not exist; skipping\n\n", shell)
	}
}

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("%v\n", os.Args)
	} else {
		// Calling Bash in multiple patterns in case we don't know which Bash is in the path
		runWith("bash")
		runWith(`C:\cygwin64\bin\bash.exe`)
		runWith(`C:\Program Files\Git\usr\bin\bash.exe`)

		// Calling myself for reference
		runWith("./main")

		// Executing AOT-compiled argv.c for reference
		// Run both executables by a native compiler and Cygwin GCC
		runWith("./argv")
		runWith("./argv-cygwin")
	}
}
