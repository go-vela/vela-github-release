// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func TestGithubRelease_Upload_Command(t *testing.T) {
	// setup types
	u := &Upload{
		Clobber: false,
		Files:   []string{"testdata/file"},
		Tag:     "tag",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		uploadAction,
		"tag",
		"testdata/file",
		fmt.Sprintf("--clobber=%t", u.Clobber),
	)

	got := u.Command()

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

func TestGithubRelease_Upload_Command_FileMissing(t *testing.T) {
	// setup types
	u := &Upload{
		Clobber: false,
		Files:   []string{"testdata/file_missing"},
		Tag:     "tag",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		uploadAction,
		"tag",
		fmt.Sprintf("--clobber=%t", u.Clobber),
	)

	got := u.Command()

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

func TestGithubRelease_Upload_Command_MultipleFiles(t *testing.T) {
	// setup types
	u := &Upload{
		Clobber: false,
		Files:   []string{"testdata/test1.txt", "testdata/test2.txt"},
		Tag:     "tag",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		uploadAction,
		"tag",
		"testdata/test1.txt",
		"testdata/test2.txt",
		fmt.Sprintf("--clobber=%t", u.Clobber),
	)

	got := u.Command()

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

func TestGithubRelease_Upload_Command_MultipleFilesGlob(t *testing.T) {
	// setup types
	u := &Upload{
		Clobber: false,
		Files:   []string{"testdata/*.txt"},
		Tag:     "tag",
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		uploadAction,
		"tag",
		"testdata/test1.txt",
		"testdata/test2.txt",
		fmt.Sprintf("--clobber=%t", u.Clobber),
	)

	got := u.Command()

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

func TestGithubRelease_Upload_Exec_Error(t *testing.T) {
	// setup types
	u := &Upload{
		Clobber: false,
		Files:   []string{"files"},
		Tag:     "tag",
	}

	err := u.Exec()
	if err == nil {
		t.Errorf("Exec should have returned err: %v", err)
	}
}

func TestGithubRelease_Upload_Validate(t *testing.T) {
	// setup types
	u := &Upload{
		Clobber: true,
		Files:   []string{"files"},
		Tag:     "tag",
	}

	err := u.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGithubRelease_Upload_Validate_Error(t *testing.T) {
	// setup types
	u := &Upload{
		Clobber: false,
		Files:   []string{""},
		Tag:     "",
	}

	err := u.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err: %v", ErrorNoUploadTag)
	}

	if !errors.Is(err, ErrorNoUploadTag) {
		t.Errorf("Validate should have returned err: %v, instead returned %v", ErrorNoUploadTag, err)
	}
}
