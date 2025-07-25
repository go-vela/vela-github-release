// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"

	"github.com/spf13/afero"
)

func TestGithubRelease_Config_Command(t *testing.T) {
	// setup types
	c := &Config{
		Action:   "action",
		Hostname: "hostname",
		Path:     tokenFile,
		Token:    "token",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.CommandContext(
		t.Context(),
		_gh,
		"auth",
		"login",
		fmt.Sprintf("--hostname=%s", c.Hostname),
		"--with-token",
	)

	got := c.Command(t.Context())

	if got.Path != want.Path {
		t.Errorf("Command path is %v, want %v", got.Path, want.Path)
	}

	if len(got.Args) != len(want.Args) {
		t.Errorf("Command args length is %v, want %v", len(got.Args), len(want.Args))
	}

	for i, arg := range got.Args {
		if i < len(want.Args) && arg != want.Args[i] {
			t.Errorf("Command args[%d] is %v, want %v", i, arg, want.Args[i])
		}
	}
}

func TestGithubRelease_Config_Exec_Error(t *testing.T) {
	// setup types
	c := &Config{
		Action:   "action",
		Hostname: "hostname",
		Path:     tokenFile,
		Token:    "token",
	}

	err := c.Exec(t.Context())
	if err == nil {
		t.Errorf("Exec should have returned err: %v", err)
	}
}

func TestGithubRelease_Config_Validate_Success(t *testing.T) {
	// setup types
	c := &Config{
		Action:   "action",
		Hostname: "hostname",
		Path:     tokenFile,
		Token:    "token",
	}

	err := c.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGithubRelease_Config_Validate_Error(t *testing.T) {
	tests := []struct {
		name    string
		c       *Config
		wantErr error
	}{
		{
			name: "No action provided",
			c: &Config{
				Action:   "",
				Hostname: "hostname",
				Path:     tokenFile,
				Token:    "token",
			},
			wantErr: ErrorNoConfigAction,
		},
		{
			name: "No token provided",
			c: &Config{
				Action:   "action",
				Hostname: "hostname",
				Path:     tokenFile,
				Token:    "",
			},
			wantErr: ErrorNoConfigGitToken,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.c.Validate(); err == nil {
				t.Errorf("Validate() should have raised an error %v", test.wantErr)
			} else if !errors.Is(err, test.wantErr) {
				t.Errorf("Validate() error = %v, wantErr = %v", err, test.wantErr)
			}
		})
	}
}

func TestGithubRelease_Config_Write(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	c := &Config{
		Action:   "action",
		Hostname: "hostname",
		Path:     tokenFile,
		Token:    "token",
	}

	err := c.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}

func TestGithubRelease_Config_Write_Error(t *testing.T) {
	// setup filesystem
	appFS = afero.NewReadOnlyFs(afero.NewMemMapFs())

	// setup types
	c := &Config{
		Action:   "action",
		Hostname: "hostname",
		Path:     tokenFile,
		Token:    "token",
	}

	err := c.Write()
	if err == nil {
		t.Errorf("Write should have returned err")
	}
}

func TestGithubRelease_Config_Write_NoFile(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	c := &Config{
		Action:   "action",
		Hostname: "hostname",
		Path:     tokenFile,
		Token:    "token",
	}

	err := c.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}
