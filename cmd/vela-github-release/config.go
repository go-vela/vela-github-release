// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/sirupsen/logrus"
)

var (
	appFS = afero.NewOsFs()

	// ErrorNoConfigAction is return when an action isn't provided.
	ErrorNoConfigAction = errors.New("no config action provided")

	// ErrorNoConfigGitToken is returned when the github token isn't provided.
	ErrorNoConfigGitToken = errors.New("no config github token provided")
)

const tokenFile = "/root/token"

// Config represents the plugin configuration for github information.
type Config struct {
	// action to perform against gh
	Action string
	// hostname to set for gh
	Hostname string
	// path to tokenFile for gh auth login
	Path string
	// token to provide to authenticate to github hostname
	Token string
}

// Command formats and outputs the Config command from
// the provided configuration to config resources.
func (c *Config) Command() *exec.Cmd {
	logrus.Trace("creating gh config command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// add flag for auth command
	flags = append(flags, "auth", "login")

	// check if hostname is provided
	if len(c.Hostname) > 0 {
		// add flag for hostname from provided config hostname
		flags = append(flags, fmt.Sprintf("--hostname=%s", c.Hostname))
	}
	// add flag for with token command
	flags = append(flags, "--with-token")

	return exec.Command(_gh, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resouces.
func (c *Config) Exec() error {
	logrus.Debug("running config with provided configuration")

	// create gh token file for authentication
	err := c.Write()
	if err != nil {
		return err
	}

	// open file in stdin
	file, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	// create command for the gh auth login
	cmd := c.Command()

	// run the config command for the gh auth login
	err = execCmd(cmd, file)
	if err != nil {
		return err
	}

	return nil
}

// Write creates a file in the home directory of the current user.
func (c *Config) Write() error {
	logrus.Trace("writing gh token file")

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if token file content is provided
	if len(c.Token) == 0 {
		return nil
	}

	// capture current user running commands
	u, err := user.Current()
	if err == nil {
		// create full path for token file
		c.Path = filepath.Join(u.HomeDir, "token")
	}

	logrus.Debug("Creating gh token file ", c.Path)

	// send Filesystem call to create directory path for token file
	err = a.Fs.MkdirAll(filepath.Dir(c.Path), 0777)
	if err != nil {
		return err
	}

	return a.WriteFile(c.Path, []byte(c.Token), 0600)
}

// Validate verifies the Config is properly configured.
func (c *Config) Validate() error {
	logrus.Trace("validating config configuration")

	// verify action is provided
	if len(c.Action) == 0 {
		return ErrorNoConfigAction
	}

	// verify token is provided
	if len(c.Token) == 0 {
		return ErrorNoConfigGitToken
	}

	return nil
}
