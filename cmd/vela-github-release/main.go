// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	_ "github.com/joho/godotenv/autoload"

	"github.com/go-vela/vela-github-release/version"
)

func main() {
	// capture application version information.
	pluginVersion := version.New()
	// serialize the version information as pretty JSON
	bytes, err := json.MarshalIndent(pluginVersion, "", "  ")
	if err != nil {
		logrus.Fatal(err)
	}

	// output the version information to stdout
	fmt.Fprintf(os.Stdout, "%s\n", string(bytes))

	// create new CLI application
	app := &cli.Command{
		Name:      "vela-github-release",
		Usage:     "Vela Github Release plugin for managing Gihub Releases in a Vela Pipeline.",
		Copyright: "Copyright 2022 Target Brands, Inc. All rights reserved.",
		Authors:   []any{"Vela Admins <vela@target.com>"},
		Action:    run,
		Version:   pluginVersion.Semantic(),
		Flags: []cli.Flag{
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
				Usage: "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_LOG_LEVEL"),
					cli.EnvVar("VELA_LOG_LEVEL"),
					cli.EnvVar("GITHUB_RELEASE_LOG_LEVEL"),
					cli.File("/vela/parameters/github-release/log_level"),
					cli.File("/vela/secrets/github-release/log_level"),
				),
				Value: "info",
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
			// Config Flags
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
				Usage: "hostname to set for github instance",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_HOSTNAME"),
					cli.EnvVar("CONFIG_HOSTNAME"),
					cli.EnvVar("GH_HOST"),
					cli.EnvVar("GITHUB_HOST"),
					cli.File("/vela/parameters/github-release/config/hostname"),
					cli.File("/vela/secrets/github-release/config/hostname"),
				),
				Value: "github.com",
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
				Usage: "target branch or commit SHA",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_TARGET"),
					cli.EnvVar("CREATE_TARGET"),
					cli.File("/vela/parameters/github-release/create/target"),
					cli.File("/vela/secrets/github-release/create/target"),
				),
				Value: "main",
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
			// Download Flags
			&cli.StringFlag{
				Name:  "download.dir",
				Usage: "the directory to download files",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_DIR"),
					cli.EnvVar("DOWNLOAD_DIR"),
					cli.File("/vela/parameters/github-release/download/dir"),
					cli.File("/vela/secrets/github-release/download/dir"),
				),
				Value: ".",
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
				Usage: "maximum number of items to fetch",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_LIMIT"),
					cli.EnvVar("LIST_LIMIT"),
					cli.File("/vela/parameters/github-release/list/limit"),
					cli.File("/vela/secrets/github-release/list/limit"),
				),
				Value: 30,
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
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(ctx context.Context, c *cli.Command) error {
	// set the log level for the plugin
	switch c.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	if c.IsSet("ci") {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: false,
			PadLevelText:  true,
		})
	}

	logrus.WithFields(logrus.Fields{
		"code":     "https://github.com/go-vela/vela-github-release",
		"docs":     "https://go-vela.github.io/docs/plugins/registry/github-release/",
		"registry": "https://hub.docker.com/r/target/vela-github-release",
	}).Info("Vela Github Release Plugin")

	// capture custom gh version requested
	version := c.String("gh.version")

	// check is a custom gh version was requested
	if len(version) > 0 {
		// attempt to install the custom gh version
		if err := install(ctx, version, os.Getenv("PLUGIN_GH_VERSION")); err != nil {
			return err
		}
	}

	// create the plugin
	p := &Plugin{
		// config configuration
		Config: &Config{
			Action:   c.String("config.action"),
			Hostname: c.String("config.hostname"),
			Path:     tokenFile,
			Token:    c.String("config.token"),
		},
		// create configuration
		Create: &Create{
			Draft:      c.Bool("create.draft"),
			Files:      c.StringSlice("files"),
			Notes:      c.String("create.notes"),
			NotesFile:  c.String("create.notes_file"),
			Prerelease: c.Bool("create.prerelease"),
			Tag:        c.String("tag"),
			Target:     c.String("create.target"),
			Title:      c.String("create.title"),
		},
		// delete configuration
		Delete: &Delete{
			Yes: c.Bool("delete.yes"),
			Tag: c.String("tag"),
		},
		// download configuration
		Download: &Download{
			Directory: c.String("download.dir"),
			Patterns:  c.StringSlice("download.patterns"),
			Tag:       c.String("tag"),
		},
		// list configuration
		List: &List{
			Limit: c.Int("list.limit"),
		},
		// upload configuration
		Upload: &Upload{
			Clobber: c.Bool("upload.clobber"),
			Files:   c.StringSlice("files"),
			Tag:     c.String("tag"),
		},
		// view configuration
		View: &View{
			Tag: c.String("tag"),
			Web: c.Bool("view.web"),
		},
	}

	// validate the plugin
	if err := p.Validate(); err != nil {
		return err
	}

	// execute the plugin
	return p.Exec(ctx)
}
