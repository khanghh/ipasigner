package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	gitDate   = ""
	// The app that holds all commands and flags.
	app *cli.App
)

func init() {
	app = cli.NewApp()
	app.Action = run
	app.Name = filepath.Base(os.Args[0])
	app.Usage = "IPA Signer"
	app.Version = fmt.Sprintf("%s - %s ", gitCommit, gitDate)
	app.Commands = []*cli.Command{
		inspectCommand,
	}
	app.Flags = []cli.Flag{
		pkcs12FileFlag,
		pkcs12PasswdFlag,
	}
}

func run(ctx *cli.Context) error {
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
