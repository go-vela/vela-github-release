// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func TestGithubRelease_Create_Command(t *testing.T) {
	// setup types
	c := &Create{
		Draft:      false,
		Files:      []string{"testdata/file"},
		Notes:      "notes",
		NotesFile:  "notes_file",
		Prerelease: false,
		Tag:        "tag",
		Target:     "target",
		Title:      "title",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.CommandContext(
		t.Context(),
		_gh,
		releaseCmd,
		createAction,
		"tag",
		"testdata/file",
		fmt.Sprintf("--draft=%t", false),
		fmt.Sprintf("--notes=%s", c.Notes),
		fmt.Sprintf("--notes-file=%s", c.NotesFile),
		fmt.Sprintf("--prerelease=%t", false),
		fmt.Sprintf("--target=%s", c.Target),
		fmt.Sprintf("--title=%s", c.Title),
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

func TestGithubRelease_Create_Command_FileMissing(t *testing.T) {
	// setup types
	c := &Create{
		Draft:      false,
		Files:      []string{"testdata/file_missing"},
		Notes:      "notes",
		NotesFile:  "notes_file",
		Prerelease: false,
		Tag:        "tag",
		Target:     "target",
		Title:      "title",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.CommandContext(
		t.Context(),
		_gh,
		releaseCmd,
		createAction,
		"tag",
		fmt.Sprintf("--draft=%t", false),
		fmt.Sprintf("--notes=%s", c.Notes),
		fmt.Sprintf("--notes-file=%s", c.NotesFile),
		fmt.Sprintf("--prerelease=%t", false),
		fmt.Sprintf("--target=%s", c.Target),
		fmt.Sprintf("--title=%s", c.Title),
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

func TestGithubRelease_Create_Command_MultipleFiles(t *testing.T) {
	// setup types
	c := &Create{
		Draft:      false,
		Files:      []string{"testdata/test1.txt", "testdata/test2.txt"},
		Notes:      "notes",
		NotesFile:  "notes_file",
		Prerelease: false,
		Tag:        "tag",
		Target:     "target",
		Title:      "title",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.CommandContext(
		t.Context(),
		_gh,
		releaseCmd,
		createAction,
		"tag",
		"testdata/test1.txt",
		"testdata/test2.txt",
		fmt.Sprintf("--draft=%t", false),
		fmt.Sprintf("--notes=%s", c.Notes),
		fmt.Sprintf("--notes-file=%s", c.NotesFile),
		fmt.Sprintf("--prerelease=%t", false),
		fmt.Sprintf("--target=%s", c.Target),
		fmt.Sprintf("--title=%s", c.Title),
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

func TestGithubRelease_Create_Command_MultipleFilesGlob(t *testing.T) {
	// setup types
	c := &Create{
		Draft:      false,
		Files:      []string{"testdata/*.txt"},
		Notes:      "notes",
		NotesFile:  "notes_file",
		Prerelease: false,
		Tag:        "tag",
		Target:     "target",
		Title:      "title",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.CommandContext(
		t.Context(),
		_gh,
		releaseCmd,
		createAction,
		"tag",
		"testdata/test1.txt",
		"testdata/test2.txt",
		fmt.Sprintf("--draft=%t", false),
		fmt.Sprintf("--notes=%s", c.Notes),
		fmt.Sprintf("--notes-file=%s", c.NotesFile),
		fmt.Sprintf("--prerelease=%t", false),
		fmt.Sprintf("--target=%s", c.Target),
		fmt.Sprintf("--title=%s", c.Title),
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

	err := c.Exec(t.Context())
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
			} else if !errors.Is(err, test.wantErr) {
				t.Errorf("Validate() error = %v, wantErr = %v", err, test.wantErr)
			}
		})
	}
}
