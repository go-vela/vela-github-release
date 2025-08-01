// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/urfave/cli/v3"

// flags is a helper function to return all
// supported command line interface (CLI) flags
// for the plugin.
func flags() []cli.Flag {
	var _flags []cli.Flag

	_flags = append(_flags, coreFlags()...)
	_flags = append(_flags, githubConfigFlags()...)
	_flags = append(_flags, releaseOperationFlags()...)
	_flags = append(_flags, utilityFlags()...)

	return _flags
}

// coreFlags returns the core application flags.
func coreFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "files",
			Usage: "files name used for action",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_FILES"),
				cli.EnvVar("GITHUB_RELEASE_FILES"),
				cli.File("/vela/parameters/github-release/files"),
				cli.File("/vela/secrets/github-release/files"),
			),
		},
		&cli.StringFlag{
			Name:  "log.level",
			Value: "info",
			Usage: "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_LOG_LEVEL"),
				cli.EnvVar("VELA_LOG_LEVEL"),
				cli.EnvVar("GITHUB_RELEASE_LOG_LEVEL"),
				cli.File("/vela/parameters/github-release/log_level"),
				cli.File("/vela/secrets/github-release/log_level"),
			),
		},
		&cli.StringFlag{
			Name:  "tag",
			Usage: "tag name used for action",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_TAG"),
				cli.EnvVar("GITHUB_RELEASE_TAG"),
				cli.File("/vela/parameters/github-release/tag"),
				cli.File("/vela/secrets/github-release/tag"),
			),
		},
		&cli.StringFlag{
			Name:  "gh.version",
			Usage: "set gh version for plugin",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_VERSION"),
				cli.EnvVar("VELA_GH_VERSION"),
				cli.EnvVar("GH_VERSION"),
				cli.File("/vela/parameters/github-release/gh/version"),
				cli.File("/vela/secrets/github-release/version"),
			),
		},
	}
}

// githubConfigFlags returns configuration-related flags.
func githubConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "config.action",
			Usage: "action to perform against github instance",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_ACTION"),
				cli.EnvVar("CONFIG_ACTION"),
				cli.File("/vela/parameters/github-release/config/action"),
				cli.File("/vela/secrets/github-release/config/action"),
			),
		},
		&cli.StringFlag{
			Name:  "config.hostname",
			Value: "github.com",
			Usage: "hostname to set for github instance",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_HOSTNAME"),
				cli.EnvVar("CONFIG_HOSTNAME"),
				cli.EnvVar("GH_HOST"),
				cli.EnvVar("GITHUB_HOST"),
				cli.File("/vela/parameters/github-release/config/hostname"),
				cli.File("/vela/secrets/github-release/config/hostname"),
			),
		},
		&cli.StringFlag{
			Name:  "config.token",
			Usage: "token to set to authenticate to github instance",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_TOKEN"),
				cli.EnvVar("CONFIG_TOKEN"),
				cli.EnvVar("GH_TOKEN"),
				cli.EnvVar("GITHUB_TOKEN"),
				cli.File("/vela/parameters/github-release/config/token"),
				cli.File("/vela/secrets/github-release/config/token"),
			),
		},
	}
}

// releaseOperationFlags returns flags for create, delete, and view operations.
func releaseOperationFlags() []cli.Flag {
	return []cli.Flag{
		// Create Flags
		&cli.BoolFlag{
			Name:  "create.draft",
			Usage: "save the release as a draft instead of publishing it",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_DRAFT"),
				cli.EnvVar("CREATE_DRAFT"),
				cli.File("/vela/parameters/github-release/create/draft"),
				cli.File("/vela/secrets/github-release/create/draft"),
			),
		},
		&cli.StringFlag{
			Name:  "create.notes",
			Usage: "create release notes",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_NOTES"),
				cli.EnvVar("CREATE_NOTES"),
				cli.File("/vela/parameters/github-release/create/notes"),
				cli.File("/vela/secrets/github-release/create/notes"),
			),
		},
		&cli.StringFlag{
			Name:  "create.notes_file",
			Usage: "read release notes from file",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_NOTES_FILE"),
				cli.EnvVar("CREATE_NOTES_FILE"),
				cli.File("/vela/parameters/github-release/create/notes_file"),
				cli.File("/vela/secrets/github-release/create/notes_file"),
			),
		},
		&cli.BoolFlag{
			Name:  "create.prerelease",
			Usage: "mark the release as a prerelease",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_PRERELEASE"),
				cli.EnvVar("CREATE_PRERELEASE"),
				cli.File("/vela/parameters/github-release/create/prerelease"),
				cli.File("/vela/secrets/github-release/create/prerelease"),
			),
		},
		&cli.StringFlag{
			Name:  "create.target",
			Value: "main",
			Usage: "target branch or commit SHA",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_TARGET"),
				cli.EnvVar("CREATE_TARGET"),
				cli.File("/vela/parameters/github-release/create/target"),
				cli.File("/vela/secrets/github-release/create/target"),
			),
		},
		&cli.StringFlag{
			Name:  "create.title",
			Usage: "Release title",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_TITLE"),
				cli.EnvVar("CREATE_TITLE"),
				cli.File("/vela/parameters/github-release/create/title"),
				cli.File("/vela/secrets/github-release/create/title"),
			),
		},
		// Delete Flags
		&cli.BoolFlag{
			Name:  "delete.yes",
			Usage: "skip the confirmation prompt",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_YES"),
				cli.EnvVar("DELETE_YES"),
				cli.File("/vela/parameters/github-release/delete/yes"),
				cli.File("/vela/secrets/github-release/delete/yes"),
			),
			//  TODO: should this be set with default to bypass the prompt? : Value: true,
		},
		// View Flags
		&cli.BoolFlag{
			Name:  "view.web",
			Usage: "open the release in the browser",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_WEB"),
				cli.EnvVar("VIEW_WEB"),
				cli.File("/vela/parameters/github-release/view/web"),
				cli.File("/vela/secrets/github-release/view/web"),
			),
		},
	}
}

// utilityFlags returns flags for download, list, and upload operations.
func utilityFlags() []cli.Flag {
	return []cli.Flag{
		// Download Flags
		&cli.StringFlag{
			Name:  "download.dir",
			Value: ".",
			Usage: "the directory to download files",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_DIR"),
				cli.EnvVar("DOWNLOAD_DIR"),
				cli.File("/vela/parameters/github-release/download/dir"),
				cli.File("/vela/secrets/github-release/download/dir"),
			),
		},
		&cli.StringSliceFlag{
			Name:  "download.patterns",
			Usage: "download only assets that match glob patterns",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_PATTERNS"),
				cli.EnvVar("DOWNLOAD_PATTERNS"),
				cli.File("/vela/parameters/github-release/download/patterns"),
				cli.File("/vela/secrets/github-release/download/patterns"),
			),
		},
		// List Flags
		&cli.IntFlag{
			Name:  "list.limit",
			Value: 30,
			Usage: "maximum number of items to fetch for list action",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_LIMIT"),
				cli.EnvVar("LIST_LIMIT"),
				cli.File("/vela/parameters/github-release/list/limit"),
				cli.File("/vela/secrets/github-release/list/limit"),
			),
		},
		// Upload Flags
		&cli.BoolFlag{
			Name:  "upload.clobber",
			Usage: "overwrite existing assets of the same name",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_CLOBBER"),
				cli.EnvVar("UPLOAD_CLOBBER"),
				cli.File("/vela/parameters/github-release/upload/clobber"),
				cli.File("/vela/secrets/github-release/upload/clobber"),
			),
		},
	}
}
