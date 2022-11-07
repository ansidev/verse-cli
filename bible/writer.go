package bible

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetFilePath(baseDir string, bibleVersionCode string, chapterNumber int, verseNumber int) string {
	return fmt.Sprintf(
		"%s/%s_%s_%s.md",
		baseDir,
		bibleVersionCode,
		fmt.Sprintf("%02d", chapterNumber),
		fmt.Sprintf("%02d", verseNumber),
	)
}

func WriteVerseToFile(baseDir string, bibleVersionCode string, chapterNumber int, verseNumber int, verse string, verseAddr string) {
	filePath := GetFilePath(baseDir, bibleVersionCode, chapterNumber, verseNumber)
	f, err := os.OpenFile(filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	formattedVerse := strings.ReplaceAll(verse, "\n", "\n>\n> ")
	content := fmt.Sprintf("\n> %s\n>\n> **%s**\n", formattedVerse, verseAddr)
	if _, err := f.WriteString(content); err != nil {
		log.Println(err)
	}
}
