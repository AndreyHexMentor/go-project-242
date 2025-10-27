package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"code/code"
	"code/flags"

	"github.com/urfave/cli/v3"
)

func main() {
	var f = &flags.Flags{}

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: flags.DefineFlags(f),
		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				fmt.Println("Usage: hexlet-path-size [--flag] <path>")
				return nil
			}

			path := cmd.Args().Get(0)

			human := f.HumanReadable
			all := f.IncludeAll
			recursive := f.Recursive

			res, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return fmt.Errorf("cannot get size: %w", err)
			}

			fmt.Println(res)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
