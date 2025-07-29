// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const uploadAction = "upload"

// ErrorNoUploadTag is returned when the plugin is missing the upload tag.
var ErrorNoUploadTag = errors.New("no upload tag provided")

// Upload represents the plugin configuration for Upload config information.
type Upload struct {
	// list of asset files to be given to upload
	Files []string
	// overwrite existing assets of the same name
	Clobber bool
	// tag name to upload a release from
	Tag string
}

// Command formats and outputs the Upload command from
// the provided configuration to upload resources.
func (u *Upload) Command(ctx context.Context) *exec.Cmd {
	logrus.Trace("creating gh upload command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// add flag for release command
	flags = append(flags, releaseCmd)

	// add flag for upload command
	flags = append(flags, uploadAction)

	// check if upload tag is provided
	if len(u.Tag) > 0 {
		// add flag for tag from provided upload tag
		flags = append(flags, u.Tag)
	}

	// iterate through the files and add them as parameters
	for _, file := range u.Files {
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

	// add flag for upload from provided upload
	flags = append(flags, fmt.Sprintf("--clobber=%t", u.Clobber))

	return exec.CommandContext(ctx, _gh, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resources.
func (u *Upload) Exec(ctx context.Context) error {
	logrus.Debug("running upload with the provided configuration")

	// upload command for the existing asset
	cmd := u.Command(ctx)

	// run the upload command for the existng asset
	err := execCmd(cmd, nil)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the Upload is properly configured.
func (u *Upload) Validate() error {
	logrus.Trace("validating upload configuration")

	// verify upload tag is provided if no tag provided error
	if len(u.Tag) == 0 {
		return ErrorNoUploadTag
	}

	return nil
}
