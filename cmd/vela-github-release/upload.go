// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const uploadAction = "upload"

var (
	// ErrorNoUploadTag is returned when the plugin is missing the upload tag.
	ErrorNoUploadTag = errors.New("no upload tag provided")
)

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
func (u *Upload) Command() *exec.Cmd {
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

	// add flag for files provided by upload files
	flags = append(flags, u.Files...)

	// add flag for upload from provided upload
	flags = append(flags, fmt.Sprintf("--clobber=%t", u.Clobber))

	return exec.Command(_gh, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resources.
func (u *Upload) Exec() error {
	logrus.Debug("running upload with the provided configuration")

	// upload command for the existing asset
	cmd := u.Command()

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
