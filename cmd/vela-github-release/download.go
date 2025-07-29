// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const downloadAction = "download"

var (
	// ErrorNoDownloadDirectory is returned when no download directory is provided.
	ErrorNoDownloadDirectory = errors.New("no download directory provided")

	// ErrorNoDownloadTag is returned when the plugin is missing the download tag.
	ErrorNoDownloadTag = errors.New("no download tag provided")
)

// Download represents the plugin configuration for Download config information.
type Download struct {
	// the directory to download files into (default ".")
	Directory string
	// download only assets that match a glob pattern
	Patterns []string
	// tag name to download a release from
	Tag string
}

// Command formats and outputs the Download command from
// the provided configuration to dowmload resources.
func (d *Download) Command(ctx context.Context) *exec.Cmd {
	logrus.Trace("creating gh download command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// add flag for release command
	flags = append(flags, releaseCmd)

	// add flag for download command
	flags = append(flags, downloadAction)

	// check if download tag is provided
	if len(d.Tag) > 0 {
		// add flag for tag from provided download tag
		flags = append(flags, d.Tag)
	}

	// add flag for directory from provided download directory
	flags = append(flags, fmt.Sprintf("--dir=%s", d.Directory))

	// iterate through all downloads patterns
	for _, pattern := range d.Patterns {
		// add flag for patterns provided by download patterns
		flags = append(flags, fmt.Sprintf("--pattern=%s", pattern))
	}

	return exec.CommandContext(ctx, _gh, flags...)
}

func (d *Download) Exec(ctx context.Context) error {
	logrus.Debug("running download with provided configuration")

	// download command for the directory
	cmd := d.Command(ctx)

	// run the download command for the directory
	err := execCmd(cmd, nil)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the Download is properly configured.
func (d *Download) Validate() error {
	logrus.Trace("validating download configuration")

	// verify directory field is provided
	if len(d.Directory) == 0 {
		return ErrorNoDownloadDirectory
	}

	// verify download tag is provided if no tag provided error
	if len(d.Tag) == 0 {
		return ErrorNoDownloadTag
	}

	return nil
}
