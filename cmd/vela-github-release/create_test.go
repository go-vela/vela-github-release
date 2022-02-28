// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestGithubRelease_Create_Command(t *testing.T) {
	// setup types
	c := &Create{
		Draft:      false,
		Files:      []string{"file"},
		Notes:      "notes",
		NotesFile:  "notes_file",
		Prerelease: false,
		Tag:        "tag",
		Target:     "target",
		Title:      "title",
	}

	// nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		createAction,
		"tag",
		"file",
		fmt.Sprintf("--draft=%t", false),
		fmt.Sprintf("--notes=%s", c.Notes),
		fmt.Sprintf("--notes-file=%s", c.NotesFile),
		fmt.Sprintf("--prerelease=%t", false),
		fmt.Sprintf("--target=%s", c.Target),
		fmt.Sprintf("--title=%s", c.Title),
	)

	got := c.Command()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("execCmd is %v, want %v", got, want)
	}
}

func TestGithubRelease_Create_Exec_Error(t *testing.T) {
	// setup types
	c := &Create{
		Draft:      false,
		Files:      []string{"file"},
		Notes:      "notes",
		NotesFile:  "notes_file",
		Prerelease: false,
		Tag:        "tag",
		Target:     "target",
		Title:      "title",
	}

	err := c.Exec()
	if err == nil {
		t.Errorf("Exec should have returned err: %v", err)
	}
}

func TestGithubRelease_Create_Validate_Success(t *testing.T) {
	// setup types
	c := &Create{
		Draft:      false,
		Files:      []string{"file"},
		Notes:      "notes",
		NotesFile:  "notes_file",
		Prerelease: false,
		Tag:        "tag",
		Target:     "target",
		Title:      "title",
	}

	err := c.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGithubRelease_Create_Validate_Error(t *testing.T) {
	tests := []struct {
		name    string
		c       *Create
		wantErr error
	}{
		{
			name: "No target provided",
			c: &Create{
				Draft:      false,
				Files:      []string{"file"},
				Notes:      "notes",
				NotesFile:  "notes_file",
				Prerelease: false,
				Tag:        "tag",
				Target:     "",
				Title:      "title",
			},
			wantErr: ErrorNoCreateTarget,
		},
		{
			name: "No tag provided",
			c: &Create{
				Draft:      false,
				Files:      []string{"file"},
				Notes:      "notes",
				NotesFile:  "notes_file",
				Prerelease: false,
				Tag:        "",
				Target:     "target",
				Title:      "title",
			},
			wantErr: ErrorNoCreateTag,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.c.Validate(); err == nil {
				t.Errorf("Validate() should have raised an error %v", test.wantErr)
			} else if errors.Is(err, test.wantErr) {
				t.Errorf("Validate() error = %v, wantErr = %v", err, test.wantErr)
			}
		})
	}
}
