package utils

import "github.com/urfave/cli/v2"

const (
	ValidChapterAndVerse = iota
	InvalidChapterNumber
	InvalidVerseNumber
)

/*
ParseInputArguments takes in a cli.Context object and returns three integers:
the chapter number, the verse number, and an error code.

Parameters:
- ctx: a cli.Context object containing the user's input arguments

Returns:
  - chapter: an integer representing the chapter number provided by the user
  - verse: an integer representing the verse number provided by the user
  - errorCode: an integer representing any errors encountered during parsing
    (e.g. InvalidChapterNumber, InvalidVerseNumber, ValidChapterAndVerse)
*/
func ParseInputArguments(ctx *cli.Context) (int, int, int) {
	chapter := ctx.Int("chapter")
	verse := ctx.Int("verse")

	if chapter < 1 || chapter > 150 {
		return -1, -1, InvalidChapterNumber
	}

	if verse < 1 || verse > 176 {
		return -1, -1, InvalidVerseNumber
	}

	return chapter, verse, ValidChapterAndVerse
}
