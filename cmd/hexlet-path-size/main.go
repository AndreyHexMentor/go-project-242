package main

import (
	"code/code"
	"code/flags"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var f = &flags.Flags{}

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: flags.DefineFlags(f),
		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				fmt.Println("Usage: hexlet-path-size [--human] <path>")
				return nil
			}

			path := cmd.Args().Get(0)

			human := f.HumanReadable
			all := f.IncludeAll

			res, err := code.GetPathSize(path, false, human, all)
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
