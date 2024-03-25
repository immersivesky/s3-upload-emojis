package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/immersivesky/s3-upload-emojis/convert"
	"github.com/immersivesky/s3-upload-emojis/image"
	"github.com/immersivesky/s3-upload-emojis/s3"
	"github.com/joho/godotenv"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	client := s3.NewS3("emojis")

	count, size := client.GetObjectCount(), client.GetSize()
	fmt.Printf("Bucket: %s | Count: %s, count: %s", client.BucketName, size, count)

	parseHTML(client)
}

type Emoji struct {
	Symbol   string
	Name     string
	PhotoURL string
}

func parseHTML(s3 *s3.S3) {
	collector := colly.NewCollector()
	collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	ch := make(chan int)
	collector.OnHTML("div.col-md-2.col-sm-3.col-xs-6.platform-grid-child.p-0", onHTML(s3, ch))
	<-ch

	if err := collector.Visit("https://emojik.com/facebook"); err != nil {
		panic(err)
	}
}

func onHTML(s3 *s3.S3, ch chan int) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		ch <- 1

		emoji := Emoji{
			Symbol:   element.ChildText("div.symbol-emoji"),
			PhotoURL: element.ChildAttr("span", "data-name"),
		}

		fmt.Println("PhotoURL:", emoji.PhotoURL)

		if link := strings.SplitAfterN(emoji.PhotoURL, ".png", -1); len(link) > 0 {
			emoji.PhotoURL = link[0]
			editName(&emoji)

			var writer bytes.Buffer
			convert.PNGToWEBP(image.GetImage(emoji.PhotoURL), &writer)

			s3.Upload(&writer, "facebook/4.0/webp/"+emoji.Name+".webp")
		} else {
			panic(link)
		}
	}
}

// TODO: 💩 (непонятно же ни*уя)
func editName(emoji *Emoji) {
	if pathName := strings.SplitAfterN(strings.ToLower(emoji.PhotoURL), "facebook\\", -1); len(pathName) > 1 {
		if emojiPathName := strings.SplitAfterN(pathName[1], "-on-facebook", -1); len(emojiPathName) > 0 {
			emoji.Name = strings.ReplaceAll(emojiPathName[0], "-on-facebook", "")
		} else {
			panic(emojiPathName)
		}
	} else {
		panic(pathName)
	}
}
