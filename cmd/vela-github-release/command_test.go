// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"
	"os/exec"
	"testing"
)

func TestGithubRelease_execCmd(t *testing.T) {
	// setup types
	e := exec.CommandContext(context.Background(), "echo", "hello")

	err := execCmd(e, os.Stdin)
	if err != nil {
		t.Errorf("execCmd returned err: %v", err)
	}
}

func TestGithubRelease_versionCmd(t *testing.T) {
	want := exec.CommandContext(
		context.Background(),
		_gh,
		"version",
	)

	got := versionCmd(context.Background())

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
