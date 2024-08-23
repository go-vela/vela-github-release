// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const createAction = "create"

var (
	// ErrorNoCreateTag is returned when the plugin is missing the create tag.
	ErrorNoCreateTag = errors.New("no create tag provided")

	// ErrorNoCreateTarget is returned when the plugin is missing the create target.
	ErrorNoCreateTarget = errors.New("no create target provided")
)

// Create represents the plugin configuration for Create config information.
type Create struct {
	// save the release as a draft instead of publishing it
	Draft bool
	// list of asset files to be given to create the release
	Files []string
	// create release notes
	Notes string
	// read release notes from file
	NotesFile string
	// mark the release as a prerelease
	Prerelease bool
	// tag name to create a release from
	Tag string
	// target branch or commit SHA (default: main branch)
	Target string
	// release title
	Title string
}

// Command formats and outputs the Create command from
// the provided configuration to create resources.
func (c *Create) Command() *exec.Cmd {
	logrus.Trace("creating gh create command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// add flag for release command
	flags = append(flags, releaseCmd)

	// add flag for create command
	flags = append(flags, createAction)

	// check if create tag is provided
	if len(c.Tag) > 0 {
		// add flag for tag from provided create tag
		flags = append(flags, c.Tag)
	}

	// iterate through the files and add them as parameters
	for _, file := range c.Files {
		f, err := filepath.Glob(file)
		if err != nil {
			logrus.Warnf("bad file pattern: %v", err)
		}

		if f == nil {
			logrus.Warnf("no file matches found for %s", file)

			continue
		}

		flags = append(flags, f...)
	}

	// add flag for draft from provided create draft
	flags = append(flags, fmt.Sprintf("--draft=%t", c.Draft))

	// check if create notes is provided
	if len(c.Notes) > 0 {
		// add flag for notes from provided create notes
		flags = append(flags, fmt.Sprintf("--notes=%s", c.Notes))
	}

	// check if create notesfile is provided
	if len(c.NotesFile) > 0 {
		// add flag for notesfile from provided create notesfile
		flags = append(flags, fmt.Sprintf("--notes-file=%s", c.NotesFile))
	}

	// add flag for prerelease from provided create prerelease
	flags = append(flags, fmt.Sprintf("--prerelease=%t", c.Prerelease))

	// check if create target branch is provided
	if len(c.Target) > 0 {
		// add flag for target from provided create target
		flags = append(flags, fmt.Sprintf("--target=%s", c.Target))
	}

	// check if create title is provided
	if len(c.Title) > 0 {
		// add flag for title from provided create title
		flags = append(flags, fmt.Sprintf("--title=%s", c.Title))
	}

	return exec.Command(_gh, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resources.
func (c *Create) Exec() error {
	logrus.Debug("running create with provided configuration")

	// create command for the target branch
	cmd := c.Command()

	// run the create command for the target branch
	err := execCmd(cmd, nil)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the Create is properly configured.
func (c *Create) Validate() error {
	logrus.Trace("validating create configuration")

	// verify target branch is provided if no target provided error
	if len(c.Target) == 0 {
		return ErrorNoCreateTarget
	}

	// verify create tag is provided if no tag provided error
	if len(c.Tag) == 0 {
		return ErrorNoCreateTag
	}

	return nil
}
