// SPDX-License-Identifier: Apache-2.0

//nolint: dupl // ignore code similarity to keep consistent structure
package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const viewAction = "view"

var (
	// ErrorNoViewTag is returned when the plugin is missing the view tag.
	ErrorNoViewTag = errors.New("no view tag provided")
)

// View represents the plugin configuration for View config information.
type View struct {
	// tag name to view a release from
	Tag string
	// open the release in the browser
	Web bool
}

// Command formats and outputs the View command from
// the provided configuration to view resources.
func (v *View) Command() *exec.Cmd {
	logrus.Trace("creatig gh view command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// add flag for release command
	flags = append(flags, releaseCmd)

	// add flag for view command
	flags = append(flags, viewAction)

	// check if view tag is provided
	if len(v.Tag) > 0 {
		// add flag for tag from provided view tag
		flags = append(flags, v.Tag)
	}

	// add flag for the view from provided view
	flags = append(flags, fmt.Sprintf("--web=%t", v.Web))

	return exec.Command(_gh, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resources.
func (v *View) Exec() error {
	logrus.Debug("running view with provided configuration")

	// view command for the release in a browser
	cmd := v.Command()

	// run the view command for the release in a browser
	err := execCmd(cmd, nil)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the View is properly configured.
func (v *View) Validate() error {
	logrus.Trace("validating view configuration")

	// verify view tag is provided if no tag provided error
	if len(v.Tag) == 0 {
		return ErrorNoViewTag
	}

	return nil
}
