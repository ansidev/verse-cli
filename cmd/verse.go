package cmd

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ansidev/verse-cli/bible"
	"github.com/ansidev/verse-cli/utils"
	"github.com/gocolly/colly/v2"
	"github.com/gosuri/uilive"
	"github.com/urfave/cli/v2"
)

const (
	BaseURL            = "https://www.bible.com/bible"
	BibleVersionCode   = "VIE2010"
	BibleVersionId     = 151
	BaseDir            = "./data"
	AddrFormatMonthDay = "md"
	AddrFormatDayMonth = "dm"
	CtxKeyJobNumber    = "job_number"
)

func VerseCommandHandler(ctx *cli.Context) error {
	month := ctx.Int("month")
	day := ctx.Int("day")
	addrFormat := ctx.String("format")

	resultCode := utils.ValidateMonthAndDay(month, day)

	if resultCode == utils.InvalidMonth {
		return fmt.Errorf("invalid month: %d", month)
	}

	if resultCode == utils.InvalidDay {
		return fmt.Errorf("invalid day: %d", day)
	}

	chapterNumber, verseNumber := day, month
	if addrFormat == AddrFormatMonthDay {
		chapterNumber = month
		verseNumber = day
	}

	done := make(chan bool, 1)
	totalJobs := 0
	successJobs := 0
	failedJobs := 0
	messages := sync.Map{}
	jobMap := sync.Map{}
	writer := uilive.New()

	// start listening for updates and render
	writer.Start()

	c := colly.NewCollector(
		// allow only bible.com to be crawled, will visit all links if not set
		colly.AllowedDomains("bible.com", "www.bible.com"),
		// sets the recursion depth for links to visit, goes on forever if not set
		colly.MaxDepth(3),
		// enables asynchronous network requests
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		url := r.URL.String()
		verseAddr := strings.ReplaceAll(url, fmt.Sprintf("%s/%d/", BaseURL, BibleVersionId), "")
		v, b := jobMap.Load(verseAddr)

		if !b {
			return
		}

		jobNumber := v.(int)

		r.Ctx.Put(CtxKeyJobNumber, jobNumber)

		messages.Store(jobNumber, fmt.Sprintf("Job #%02d: Fetching verse %s\n", jobNumber, verseAddr))
	})

	c.OnHTML("#__next > div > main > div > div > div > div:nth-child(1)", func(e *colly.HTMLElement) {
		verse := e.ChildText("p")
		verseAddr := e.ChildText("h2")

		jobNumber := e.Request.Ctx.GetAny(CtxKeyJobNumber).(int)
		if len(verse) == 0 || len(verseAddr) == 0 {
			messages.Store(jobNumber, fmt.Sprintf("Job #%02d: Invalid fetched data", jobNumber))
			failedJobs++
		} else {
			messages.Store(jobNumber, fmt.Sprintf("Job #%02d: Writing verse %s\n", jobNumber, verseAddr))
			bible.WriteVerseToFile(BaseDir, BibleVersionCode, chapterNumber, verseNumber, verse, verseAddr)
			messages.Store(jobNumber, fmt.Sprintf("Job #%02d: Finished\n", jobNumber))
			successJobs++
		}

		if successJobs+failedJobs == totalJobs {
			delay(1)
			fmt.Fprintf(writer, "Finished %d jobs, success: %d, failed: %d\n", totalJobs, successJobs, failedJobs)
			writer.Stop()
			done <- true
		}
	})

	go func() {
		for {
			delay(1)
			for i := 0; i < totalJobs; i++ {
				item, b := messages.Load(i)
				if !b {
					continue
				}
				m := item.(string)
				if len(m) != 0 {
					fmt.Fprint(writer, m)
				}
			}
		}
	}()

	c.Wait()
	utils.CreateDirIfNotExists(BaseDir)

	filePath := bible.GetFilePath(BaseDir, BibleVersionCode, chapterNumber, verseNumber)
	fmt.Fprintf(writer, "Truncating file %s\n", filePath)
	err := utils.TruncateFile(filePath)
	if err != nil {
		fmt.Fprint(writer, err.Error())
	}

	for _, bookCode := range append(bible.OldStatementBookCodes, bible.NewStatementBookCodes...) {
		if !bible.IsValidVerseAddress(bookCode, chapterNumber, verseNumber) {
			continue
		}

		url := fmt.Sprintf("%s/%d/%s.%d.%d", BaseURL, BibleVersionId, bookCode, chapterNumber, verseNumber)
		err1 := c.Visit(url)
		if err1 != nil {
			return err1
		}

		totalJobs++
		verseAddr := fmt.Sprintf("%s.%d.%d", bookCode, chapterNumber, verseNumber)
		jobMap.Store(verseAddr, totalJobs)
	}

	<-done
	return nil
}

func delay(i int) {
	time.Sleep(time.Millisecond * time.Duration(100*i))
}
