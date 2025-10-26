package main

import (
	"code/code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(_ context.Context, c *cli.Command) error {
			args := c.Args()
			if args.Len() == 0 {
				return fmt.Errorf("no path provided")
			}

			path := args.Get(0)

			res, err := code.GetPathSize(path, false, false, false)
			if err != nil {
				return err
			}

			fmt.Println(res)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
