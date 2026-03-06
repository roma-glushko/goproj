package cmd

import (
	"fmt"
	"log/slog"

	"github.com/urfave/cli/v2"
)

var HelloCommand = &cli.Command{
	Name:    "hello",
	Aliases: []string{"h"},
	Usage:   "A sample CLI command",
	Flags:   []cli.Flag{},
	Action: func(_ *cli.Context) error {
		slog.Debug("running hello command")
		fmt.Println("Hey")

		return nil
	},
}
