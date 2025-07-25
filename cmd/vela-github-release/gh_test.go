// SPDX-License-Identifier: Apache-2.0

package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestGithub_CLI_install(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// run test
	err := install(t.Context(), "2.14.4", "2.14.4")
	if err != nil {
		t.Errorf("install returned err: %v", err)
	}
}

func TestGithub_CLI_install_NoBinary(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// run test
	err := install(t.Context(), "2.14.3", "2.14.4")
	if err == nil {
		t.Errorf("install should have returned err ")
	}
}

func TestGithub_CLI_install_NotWritable(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	a := &afero.Afero{
		Fs: appFS,
	}

	// create binary file
	err := a.WriteFile(_gh, []byte("!@#$%^&*()"), 0777)
	if err != nil {
		t.Errorf("Unable to write file %s: %v", _gh, err)
	}

	// run test
	err = install(t.Context(), "2.14.3", "2.14.4")
	if err == nil {
		t.Errorf("install should have returned err")
	}
}
