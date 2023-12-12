// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"reflect"
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
		fmt.Sprintf(d.Tag),
		fmt.Sprintf("--yes=%t", false),
	)

	got := d.Command()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("execCmd is %v, want %v", got, want)
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
