// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func TestGithubRelease_Download_Command(t *testing.T) {
	// setup types
	d := &Download{
		Directory: "dir",
		Patterns:  []string{"pattern"},
		Tag:       "tag",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.CommandContext(
		t.Context(),
		_gh,
		releaseCmd,
		downloadAction,
		d.Tag,
		fmt.Sprintf("--dir=%s", d.Directory),
		fmt.Sprintf("--pattern=%s", "pattern"),
	)

	got := d.Command(t.Context())

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

func TestGithubRelease_Download_Exec_Error(t *testing.T) {
	// setup types
	d := &Download{
		Directory: "dir",
		Patterns:  []string{"patterns"},
		Tag:       "tag",
	}

	err := d.Exec(t.Context())
	if err == nil {
		t.Errorf("Exec should have returned err: %v", err)
	}
}

func TestGithubRelease_Download_Validate_Success(t *testing.T) {
	// setup types
	d := &Download{
		Directory: "dir",
		Patterns:  []string{"patterns"},
		Tag:       "tag",
	}

	err := d.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGithubRelease_Download_Validate_Error(t *testing.T) {
	tests := []struct {
		name    string
		d       *Download
		wantErr error
	}{
		{
			name: "No directory provided",
			d: &Download{
				Directory: "",
				Patterns:  []string{"patterns"},
				Tag:       "tag",
			},
			wantErr: ErrorNoDownloadDirectory,
		},
		{
			name: "No tag provided",
			d: &Download{
				Directory: "dir",
				Patterns:  []string{"patterns"},
				Tag:       "",
			},
			wantErr: ErrorNoDownloadTag,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.d.Validate(); err == nil {
				t.Errorf("Validate() should have raised an error %v", test.wantErr)
			} else if !errors.Is(err, test.wantErr) {
				t.Errorf("Validate() error = %v, wantErr = %v", err, test.wantErr)
			}
		})
	}
}
