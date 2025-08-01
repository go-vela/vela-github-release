// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/mail"
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
		Authors: []any{
			&mail.Address{
				Name:    "Vela Admins",
				Address: "vela@target.com",
			},
		},
		Action:  run,
		Version: pluginVersion.Semantic(),
		Flags:   flags(),
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
