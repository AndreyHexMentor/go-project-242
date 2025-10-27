package flags

import (
	"context"

	"github.com/urfave/cli/v3"
)

// Flags хранит значения CLI-флагов
type Flags struct {
	HumanReadable bool
	IncludeAll    bool
}

// DefineFlags возвращает массив cli.Flag и привязывает значения к структуре Flags
func DefineFlags(f *Flags) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "human",
			Aliases: []string{"H"},
			Usage:   "human-readable sizes (auto-select unit)",
			Value:   false,
			Action: func(_ context.Context, _ *cli.Command, b bool) error {
				f.HumanReadable = b
				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "include hidden files and directories",
			Value:   false,
			Action: func(_ context.Context, _ *cli.Command, b bool) error {
				f.IncludeAll = b
				return nil
			},
		},
		// Здесь будем также добавлять остальные флаги
	}
}
