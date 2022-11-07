package main

import (
	"log"
	"os"
	"time"

	"github.com/ansidev/verse-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	now := time.Now().Local()
	app := &cli.App{
		Name:      "verse",
		Usage:     "Verse CLI",
		UsageText: "Get verses by month and day. Example command: verse -m=12 -d=12 -f=md",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "month",
				Aliases: []string{"m"},
				Value:   int(now.Month()),
				DefaultText: "current local month",
				Usage:   "Month to fetch verse for. Value range: [1-12]",
			},
			&cli.IntFlag{
				Name:    "day",
				Aliases: []string{"d"},
				Value:   int(now.Day()),
				DefaultText: "current local day",
				Usage:   "Day to fetch verse for. Value range: [1-31]",
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "dm",
				Usage:   "Verse address format. dm: {book} {month}:{day}, md: {book} {day}:{month}.",
			},
		},
		Action: cmd.VerseCommandHandler,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
