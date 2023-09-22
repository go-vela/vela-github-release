// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// execCmd is a helper function to
// run the provided command.
func execCmd(e *exec.Cmd, f *os.File) error {
	logrus.Tracef("executing cmd %s", strings.Join(e.Args, " "))

	// set command stout to OS stdout
	e.Stdout = os.Stdout
	// set command stderr to OS stderr
	e.Stderr = os.Stderr

	// check if file provided is empty
	if f != nil {
		// set command stdin to file
		e.Stdin = f
	}

	// output "trace" string for command
	fmt.Println("$", strings.Join(e.Args, " "))

	return e.Run()
}

// versionCmd is a helper function to output
// the gh version information.
func versionCmd() *exec.Cmd {
	logrus.Trace("creating gh version command")

	// variable to store flags for command
	var flags []string

	// add flag for version gh command
	flags = append(flags, "version")

	return exec.Command(_gh, flags...)
}
