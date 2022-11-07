package bible

func IsValidVerseAddress(bookCode string, chapterNumber int, verseNumber int) bool {
	if chapterNumber < 1 || verseNumber < 1 {
		return false
	}

	if val, ok := totalVersesGroupByBookAndChapter[bookCode]; ok {
		totalChapters := val[0]
		if chapterNumber > totalChapters {
			return false
		}

		totalVerses := val[chapterNumber]
		if verseNumber <= totalVerses {
			return true
		}
	}

	return false
}
