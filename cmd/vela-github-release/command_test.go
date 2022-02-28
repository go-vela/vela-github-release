// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestGithubRelease_execCmd(t *testing.T) {
	// setup types
	e := exec.Command("echo", "hello")

	err := execCmd(e, os.Stdin)
	if err != nil {
		t.Errorf("execCmd returned err: %v", err)
	}
}

func TestGithubRelease_versionCmd(t *testing.T) {
	want := exec.Command(
		_gh,
		"version",
	)

	got := versionCmd()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("versionCmd is %v, want %v", got, want)
	}
}
