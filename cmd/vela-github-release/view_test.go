// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestGithubRelease_View_Command(t *testing.T) {
	// setup types
	v := &View{
		Tag: "tag",
		Web: false,
	}

	// nolint:gosec // ignore for testing purposes
	want := exec.Command(
		_gh,
		releaseCmd,
		viewAction,
		"tag",
		fmt.Sprintf("--web=%t", v.Web),
	)

	got := v.Command()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("execCmd is %v, want %v", got, want)
	}
}

func TestGithubRelease_View_Exec_Error(t *testing.T) {
	// setup types
	v := &View{
		Tag: "tag",
		Web: false,
	}

	err := v.Exec()
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
