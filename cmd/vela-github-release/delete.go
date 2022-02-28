// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// nolint: dupl // ignore code similarity to keep consistent structure
package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const deleteAction = "delete"

var (
	// ErrorNoDeleteTag is returned when the plugin is missing the delete tag.
	ErrorNoDeleteTag = errors.New("no delete tag provided")
)

// Delete represents the plugin configuration for Delete config information.
type Delete struct {
	// tag name to delete a release from
	Tag string
	// Skip the confirmation prompt
	Yes bool
}

// Command formats and outputs the Delete command from
// the provided configuration to delete resources.
func (d *Delete) Command() *exec.Cmd {
	logrus.Trace("creating gh delete command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// add flag for release command
	flags = append(flags, releaseCmd)

	// add flag for delete command
	flags = append(flags, deleteAction)

	// check if delete tag is provided
	if len(d.Tag) > 0 {
		// add flag for tag from provided delete tag
		flags = append(flags, d.Tag)
	}

	// add flag for delete from provided delete
	flags = append(flags, fmt.Sprintf("--yes=%t", d.Yes))

	return exec.Command(_gh, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resources.
func (d *Delete) Exec() error {
	logrus.Debug("running delete with provided configuration")

	// delete command for the target branch
	cmd := d.Command()

	// run the delete command for the target branch
	err := execCmd(cmd, nil)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the Delete is properly configured.
func (d *Delete) Validate() error {
	logrus.Trace("validating delete configuration")

	// verify delete tag is provided if no tag provided error
	if len(d.Tag) == 0 {
		return ErrorNoDeleteTag
	}

	return nil
}
