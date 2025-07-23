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
				Sources: cli.EnvVars("PARAMETER_FILES", "GITHUB_RELEASE_FILES"),
				Name:    "files",
				Usage:   "files name used for action",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_LOG_LEVEL", "VELA_LOG_LEVEL", "GITHUB_RELEASE_LOG_LEVEL"),
				Name:    "log.level",
				Usage:   "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
				Value:   "info",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_TAG", "GITHUB_RELEASE_TAG"),
				Name:    "tag",
				Usage:   "tag name used for action",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_VERSION", "VELA_GH_VERSION", "GH_VERSION"),
				Name:    "gh.version",
				Usage:   "set gh version for plugin",
			},
			// Config Flags
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_ACTION", "CONFIG_ACTION"),
				Name:    "config.action",
				Usage:   "action to perform against github instance",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_HOSTNAME", "CONFIG_HOSTNAME", "GH_HOST", "GITHUB_HOST"),
				Name:    "config.hostname",
				Usage:   "hostname to set for github instance",
				Value:   "github.com",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_TOKEN", "CONFIG_TOKEN", "GH_TOKEN", "GITHUB_TOKEN"),
				Name:    "config.token",
				Usage:   "token to set to authenticate to github instance",
			},
			// Create Flags
			&cli.BoolFlag{
				Sources: cli.EnvVars("PARAMETER_DRAFT", "CREATE_DRAFT"),
				Name:    "create.draft",
				Usage:   "save the release as a draft instead of publishing it",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_NOTES", "CREATE_NOTES"),
				Name:    "create.notes",
				Usage:   "create release notes",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_NOTES_FILE", "CREATE_NOTES_FILE"),
				Name:    "create.notes_file",
				Usage:   "read release notes from file",
			},
			&cli.BoolFlag{
				Sources: cli.EnvVars("PARAMETER_PRERELEASE", "CREATE_PRERELEASE"),
				Name:    "create.prerelease",
				Usage:   "mark the release as a prerelease",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_TARGET", "CREATE_TARGET"),
				Name:    "create.target",
				Usage:   "target branch or commit SHA",
				Value:   "main",
			},
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_TITLE", "CREATE_TITLE"),
				Name:    "create.title",
				Usage:   "Release title",
			},
			// Delete Flags
			&cli.BoolFlag{
				Sources: cli.EnvVars("PARAMETER_YES", "DELETE_YES"),
				Name:    "delete.yes",
				Usage:   "skip the confirmation prompt",
				//  TODO: should this be set with default to bypass the prompt? : Value: true,
			},
			// Download Flags
			&cli.StringFlag{
				Sources: cli.EnvVars("PARAMETER_DIR", "DOWNLOAD_DIR"),
				Name:    "download.dir",
				Usage:   "the directory to download files",
				Value:   ".",
			},
			&cli.StringSliceFlag{
				Sources: cli.EnvVars("PARAMETER_PATTERNS", "DOWNLOAD_PATTERNS"),
				Name:    "download.patterns",
				Usage:   "download only assets that match glob patterns",
			},
			// List Flags
			&cli.IntFlag{
				Sources: cli.EnvVars("PARAMETER_LIMIT", "LIST_LIMIT"),
				Name:    "list.limit",
				Usage:   "maximum number of items to fetch",
				Value:   30,
			},
			// Upload Flags
			&cli.BoolFlag{
				Sources: cli.EnvVars("PARAMETER_CLOBBER", "UPLOAD_CLOBBER"),
				Name:    "upload.clobber",
				Usage:   "overwrite existing assets of the same name",
			},
			// View Flags
			&cli.BoolFlag{
				Sources: cli.EnvVars("PARAMETER_WEB", "VIEW_WEB"),
				Name:    "view.web",
				Usage:   "open the release in the browser",
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
	return p.Exec()
}
