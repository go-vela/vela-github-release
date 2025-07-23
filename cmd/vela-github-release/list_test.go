// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func TestGithubRelease_List_Command(t *testing.T) {
	// setup types
	l := &List{
		Limit: 30,
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		listAction,
		fmt.Sprintf("--limit=%d", l.Limit),
	)

	got := l.Command()

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

func TestGithubRelease_List_Exec_Error(t *testing.T) {
	// setup types
	l := &List{
		Limit: 30,
	}

	err := l.Exec()
	if err == nil {
		t.Errorf("Exec should have returned err: %v", err)
	}
}

func TestGithubRelease_List_Validate(t *testing.T) {
	// setup types
	l := &List{
		Limit: 30,
	}

	err := l.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGithubRelease_List_Validate_Error(t *testing.T) {
	// setup types
	l := &List{
		Limit: 0,
	}

	err := l.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err: %v", ErrorInvalidListLimit)
	}

	if !errors.Is(err, ErrorInvalidListLimit) {
		t.Errorf("Validate should have returned err: %v, instead returned %v", ErrorInvalidListLimit, err)
	}
}
