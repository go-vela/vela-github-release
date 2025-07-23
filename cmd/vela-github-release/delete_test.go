// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func TestGithubRelease_Delete_Command(t *testing.T) {
	// setup types
	d := &Delete{
		Tag: "tag",
		Yes: false,
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		deleteAction,
		d.Tag,
		fmt.Sprintf("--yes=%t", false),
	)

	got := d.Command()

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

func TestGithubRelease_Delete_Exec_Error(t *testing.T) {
	// setup types
	d := &Delete{
		Tag: "tag",
		Yes: false,
	}

	err := d.Exec()
	if err == nil {
		t.Errorf("Exec should have returned err: %v", err)
	}
}

func TestGithubRelease_Delete_Validate_Success(t *testing.T) {
	// setup types
	d := &Delete{
		Tag: "tag",
		Yes: false,
	}

	err := d.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGithubRelease_Delete_Validate_Error(t *testing.T) {
	// setup types
	d := &Delete{
		Tag: "",
		Yes: false,
	}

	err := d.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err: %v", ErrorNoDeleteTag)
	}

	if !errors.Is(err, ErrorNoDeleteTag) {
		t.Errorf("Validate should have returned err: %v, instead returned %v", ErrorNoDeleteTag, err)
	}
}
