package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			// value можем не указывать, оно по дефолту будет false
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human readable sizes",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			if c.NArg() == 0 {
				fmt.Println("Usage: hexlet-path-size [--flag] <path>")
				return nil
			}

			path := c.Args().Get(0)

			human := c.Bool("human")
			all := c.Bool("all")
			recursive := c.Bool("recursive")

			res, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				log.Printf("Warning: %v", err)
				return nil
			}

			fmt.Println(res)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
