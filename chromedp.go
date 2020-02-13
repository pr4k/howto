// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func chromedpHandler(query string) []post {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	var title string
	if err := chromedp.Run(ctx,
		searchQuery(`https://www.google.com/`, `#main`, query),
		fullScreenshot(90, &buf),
		getSearchContent(`#cnt > div:nth-child(12)`, &title),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	return getPostDetails(title)

}

func fullScreenshot(quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func searchQuery(urlstr string, sel string, query string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.SendKeys(`#tsf > div:nth-child(2) > div.A8SBwf > div.RNNXgb > div > div.a4bIc > input`, query+" site:stackoverflow.com "),
		chromedp.Click(`#tsf > div:nth-child(2) > div.A8SBwf > div.FPdoLc.tfB0Bf > center > input.gNO89b`),
		chromedp.WaitVisible(sel, chromedp.ByID),
	}
}
func getSearchContent(sel string, title *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(`#fbar > div.fbar.b2hzT`, chromedp.ByID),
		chromedp.InnerHTML(sel, title),
	}
}

func getPostDetails(body string) []post {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	var temp []post
	doc.Find(".r").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		title := linkTag.Find("h3").Text()
		temp = append(temp, post{title, link, "Cannot Fetch", "Scraped from google"})

	})
	return temp
}
