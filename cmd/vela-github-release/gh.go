// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"runtime"
	"strings"

	getter "github.com/hashicorp/go-getter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

const (
	_gh       = "/bin/gh"
	_ghTmp    = "/bin/download"
	_download = "https://github.com/cli/cli/releases/download/v%s/gh_%s_%s_%s.tar.gz//gh_%s_%s_%s/bin"
)

// install downloads a custom version of the gh cli.
func install(customVer, defaultVer string) error {
	logrus.Infof("custom gh version requested: %s", customVer)

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if the custom version matches the default version
	if strings.EqualFold(customVer, defaultVer) {
		// the gh versions match so no action required
		return nil
	}

	logrus.Debugf("custom version does not match default: %s", defaultVer)
	// rename the old gh binary since we can't overwrite it for now
	// https://github.com/hashicorp/go-getter/issues/219
	err := a.Rename(_gh, fmt.Sprintf("%s.default", _gh))
	if err != nil {
		return err
	}

	// create the download URL to install gh
	url := fmt.Sprintf(_download, customVer, customVer, runtime.GOOS, runtime.GOARCH, customVer, runtime.GOOS, runtime.GOARCH)

	logrus.Infof("downloading gh version from: %s", url)
	// send the HTTP request to install gh
	err = getter.Get(_ghTmp, url, []getter.ClientOption{}...)
	if err != nil {
		return err
	}

	// getter installed a directory of files, move the binary from that to the _gh location
	err = a.Rename(_ghTmp+"/gh", _gh)
	if err != nil {
		return err
	}

	logrus.Debugf("changing ownership of the file: %s", _gh)
	// ensure the gh binary is executable
	err = a.Chmod(_gh, 0700)
	if err != nil {
		return err
	}

	return nil
}
