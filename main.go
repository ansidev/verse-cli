package main

import (
	"log"
	"os"

	"github.com/ansidev/verse-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "verse",
		Usage:     "Verse CLI",
		UsageText: "Get verses by chapter and verse number. Example command: verse -c=1 -v=2",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "chapter",
				Aliases: []string{"c"},
				Value:   1,
				Usage:   "Chapter number. Value range: [1-150].",
			},
			&cli.IntFlag{
				Name:    "verse",
				Aliases: []string{"v"},
				Value:   1,
				Usage:   "Verse number. Value range: [1-176].",
			},
		},
		Action: cmd.VerseCommandHandler,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
