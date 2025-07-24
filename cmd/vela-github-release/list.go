// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const listAction = "list"

var (
	// ErrorInvalidListLimit is returned when the limit is not provided.
	ErrorInvalidListLimit = errors.New("List limit cannot be zero")
)

// List represents the plugin configuration for List config information.
type List struct {
	// maximum number of items to fetch (default 30)
	Limit int
}

// Command formats and outputs the List command from
// the provided configuration to list resources.
func (l *List) Command(ctx context.Context) *exec.Cmd {
	logrus.Trace("creating gh list command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// add flag for release command
	flags = append(flags, releaseCmd)

	// add flag for list command
	flags = append(flags, listAction)

	// add flag for list from provided list
	flags = append(flags, fmt.Sprintf("--limit=%d", l.Limit))

	return exec.CommandContext(ctx, _gh, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resources.
func (l *List) Exec(ctx context.Context) error {
	logrus.Debug("running list with provided configuration")

	// list command for the number of items to fetch
	cmd := l.Command(ctx)

	// run the list command for the limit of items
	err := execCmd(cmd, nil)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the List is properly configured.
func (l *List) Validate() error {
	logrus.Trace("validating list configuration")

	// verify limit is provided if no limit is provided error
	if l.Limit == 0 {
		return ErrorInvalidListLimit
	}

	return nil
}
