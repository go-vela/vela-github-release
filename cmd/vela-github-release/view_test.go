// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func TestGithubRelease_View_Command(t *testing.T) {
	// setup types
	v := &View{
		Tag: "tag",
		Web: false,
	}

	//nolint:gosec // ignore for testing purposes
	want := exec.CommandContext(
		context.Background(),
		_gh,
		releaseCmd,
		viewAction,
		"tag",
		fmt.Sprintf("--web=%t", v.Web),
	)

	got := v.Command(context.Background())

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

func TestGithubRelease_View_Exec_Error(t *testing.T) {
	// setup types
	v := &View{
		Tag: "tag",
		Web: false,
	}

	err := v.Exec(context.Background())
	if err == nil {
		t.Errorf("Exec should have returned err: %v", err)
	}
}

func TestGithubRelease_View_Validate_Success(t *testing.T) {
	// setup types
	v := &View{
		Tag: "tag",
		Web: false,
	}

	err := v.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGithubRelease_View_Validate_Error(t *testing.T) {
	// setup types
	v := &View{
		Tag: "",
		Web: false,
	}

	err := v.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err: %v", ErrorNoViewTag)
	}

	if !errors.Is(err, ErrorNoViewTag) {
		t.Errorf("Validate should have returned err: %v, instead returned %v", ErrorNoViewTag, err)
	}
}
