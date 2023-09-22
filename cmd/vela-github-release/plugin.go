// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

// releaseCmd defines the core command for gh release plugin.
const releaseCmd = "release"

var (
	// ErrInvalidAction defines the error type when the
	// Action provided to the Plugin is unsupported.
	ErrInvalidAction = errors.New("invalid action provided")
)

// Plugin represents the configuration loaded for the plugin.
type Plugin struct {
	// config arguments loaded for the plugin
	Config *Config
	// create arguments loaded for the plugin
	Create *Create
	// delete arguments loaded for the plugin
	Delete *Delete
	// download arguments loaded for the plugin
	Download *Download
	// list arguments loaded for the plugin
	List *List
	// upload arguments loaded for the plugin
	Upload *Upload
	// view arguments loaded fo rthe plugin
	View *View
}

// Exec formats and runs the commands for gh plugin.
func (p *Plugin) Exec() error {
	logrus.Debug("running plugin with provided configuration")

	// output gh version for troubleshooting
	err := execCmd(versionCmd(), nil)
	if err != nil {
		return err
	}

	// execute config configuration
	err = p.Config.Exec()
	if err != nil {
		return err
	}

	// execute action specific configuration
	switch p.Config.Action {
	case createAction:
		// execute create action
		return p.Create.Exec()
	case deleteAction:
		// execute delete action
		return p.Delete.Exec()
	case downloadAction:
		// execute download action
		return p.Download.Exec()
	case listAction:
		// execute list action
		return p.List.Exec()
	case uploadAction:
		// execute upload action
		return p.Upload.Exec()
	case viewAction:
		// execute view action
		return p.View.Exec()
	default:
		return fmt.Errorf(
			"%w: %s (Valid actions: %s, %s, %s, %s, %s, %s)",
			ErrInvalidAction,
			p.Config.Action,
			createAction,
			deleteAction,
			downloadAction,
			listAction,
			uploadAction,
			viewAction,
		)
	}
}

// Validate verfies the plugin is properly configured.
func (p *Plugin) Validate() error {
	logrus.Debug("validating plugin configuration")

	// validate config configuration
	err := p.Config.Validate()
	if err != nil {
		return err
	}

	// validate action specific configuration
	switch p.Config.Action {
	case createAction:
		// validate create configuration
		return p.Create.Validate()
	case deleteAction:
		// validate delete configuration
		return p.Delete.Validate()
	case downloadAction:
		// validate download configuration
		return p.Download.Validate()
	case listAction:
		// validate list configuration
		return p.List.Validate()
	case uploadAction:
		// validate upload configuration
		return p.Upload.Validate()
	case viewAction:
		// validate view configuration
		return p.View.Validate()
	default:
		return fmt.Errorf(
			"%w: %s (Valid actions: %s, %s, %s, %s, %s, %s)",
			ErrInvalidAction,
			p.Config.Action,
			createAction,
			deleteAction,
			downloadAction,
			listAction,
			uploadAction,
			viewAction,
		)
	}
}
