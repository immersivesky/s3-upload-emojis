package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/immersivesky/s3-upload-emojis/convert"
	"github.com/immersivesky/s3-upload-emojis/image"
	"github.com/immersivesky/s3-upload-emojis/internal/adapters/repository"
	"github.com/immersivesky/s3-upload-emojis/s3"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	client := s3.NewS3("emojis")

	count, size := client.GetObjectCount(), client.GetSize()
	fmt.Printf("Bucket: %s | Count: %s, count: %s\n", client.BucketName, size, count)

	db, err := repository.NewDB(os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		panic(err)
	}

	appleEmojiPackID, err := db.CreateEmojiPack("community", 2, "Apple")
	if err != nil {
		panic(err)
	}

	microsoftEmojiPackID, err := db.CreateEmojiPack("community", 3, "Microsoft")
	if err != nil {
		panic(err)
	}

	googleEmojiPackID, err := db.CreateEmojiPack("community", 4, "Google")
	if err != nil {
		panic(err)
	}

	facebookEmojiPackID, err := db.CreateEmojiPack("community", 5, "Facebook")
	if err != nil {
		panic(err)
	}

	mozillaEmojiPackID, err := db.CreateEmojiPack("community", 6, "Mozilla")
	if err != nil {
		panic(err)
	}

	messengerEmojiPackID, err := db.CreateEmojiPack("community", 7, "Messenger")
	if err != nil {
		panic(err)
	}

	whatsAppEmojiPackID, err := db.CreateEmojiPack("community", 8, "Whats App")
	if err != nil {
		panic(err)
	}

	xEmojiPackID, err := db.CreateEmojiPack("community", 9, "Twitter")
	if err != nil {
		panic(err)
	}

	samsungEmojiPackID, err := db.CreateEmojiPack("community", 10, "Samsung")
	if err != nil {
		panic(err)
	}

	lgEmojiPackID, err := db.CreateEmojiPack("community", 11, "LG")
	if err != nil {
		panic(err)
	}

	joyPixelsEmojiPackID, err := db.CreateEmojiPack("community", 12, "JoyPixels")
	if err != nil {
		panic(err)
	}

	emojiDexEmojiPackID, err := db.CreateEmojiPack("community", 13, "EmojiDex")
	if err != nil {
		panic(err)
	}

	openMojiEmojiPackID, err := db.CreateEmojiPack("community", 14, "OpenMoji")
	if err != nil {
		panic(err)
	}

	softBankEmojiPackID, err := db.CreateEmojiPack("community", 15, "SoftBank")
	if err != nil {
		panic(err)
	}

	htcEmojiPackID, err := db.CreateEmojiPack("community", 16, "HTC")
	if err != nil {
		panic(err)
	}

	parseHTML(client, db, []Location{
		{
			URL:         "https://emojik.com/apple",
			EmojiPackID: appleEmojiPackID,
		},
		{
			URL:         "https://emojik.com/microsoft",
			EmojiPackID: microsoftEmojiPackID,
		},
		{
			URL:         "https://emojik.com/google",
			EmojiPackID: googleEmojiPackID,
		},
		{
			URL:         "https://emojik.com/facebook",
			EmojiPackID: facebookEmojiPackID,
		},
		{
			URL:         "https://emojik.com/messenger",
			EmojiPackID: messengerEmojiPackID,
		},
		{
			URL:         "https://emojik.com/whatsapp",
			EmojiPackID: whatsAppEmojiPackID,
		},
		{
			URL:         "https://emojik.com/mozilla",
			EmojiPackID: mozillaEmojiPackID,
		},
		{
			URL:         "https://emojik.com/twitter",
			EmojiPackID: xEmojiPackID,
		},
		{
			URL:         "https://emojik.com/samsung",
			EmojiPackID: samsungEmojiPackID,
		},
		{
			URL:         "https://emojik.com/lg",
			EmojiPackID: lgEmojiPackID,
		},
		{
			URL:         "https://emojik.com/joypixels",
			EmojiPackID: joyPixelsEmojiPackID,
		},
		{
			URL:         "https://emojik.com/emojidex",
			EmojiPackID: emojiDexEmojiPackID,
		},
		{
			URL:         "https://emojik.com/openmoji",
			EmojiPackID: openMojiEmojiPackID,
		},
		{
			URL:         "https://emojik.com/softbank",
			EmojiPackID: softBankEmojiPackID,
		},
		{
			URL:         "https://emojik.com/htc",
			EmojiPackID: htcEmojiPackID,
		},
	})
}

type Location struct {
	URL         string
	EmojiPackID int
}

type Emoji struct {
	Symbol   string
	Name     string
	PhotoURL string
}

func parseHTML(s3 *s3.S3, db *repository.DB, locations []Location) {
	for _, location := range locations {
		collector := colly.NewCollector()
		collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

		collector.OnHTML("div.col-md-2.col-sm-3.col-xs-6.platform-grid-child.p-0", onHTML(s3, db, location.EmojiPackID))

		if err := collector.Visit(location.URL); err != nil {
			panic(err)
		}
	}
}

func onHTML(s3 *s3.S3, db *repository.DB, emojiPackID int) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		emoji := Emoji{
			Symbol:   element.ChildText("div.symbol-emoji"),
			PhotoURL: element.ChildAttr("span", "data-name"),
		}

		if link := strings.SplitAfterN(emoji.PhotoURL, ".png", -1); len(link) > 0 {
			emoji.PhotoURL = link[0]
			editName(&emoji)

			emoji.PhotoURL = strings.ReplaceAll(emoji.PhotoURL, "\\", "/")

			img, err := image.GetImage(emoji.PhotoURL)
			if err != nil {
				fmt.Printf("❌ Error: %v\n\t➜ Name: %s\n\t➜ URL: %s\n", err, emoji.Name, emoji.PhotoURL)
			} else {
				var writer bytes.Buffer
				convert.PNGToWEBP(img, &writer)

				location := s3.Upload(&writer, "facebook/4.0/webp/"+emoji.Name+".webp")
				emojiID, err := db.CreateEmoji(emojiPackID, location)
				if err != nil {
					panic(err)
				}

				if emojiID > 0 {
					emojiShortCodeID, err := db.CreateEmojiShortCode(emojiID, emoji.Symbol)
					if err != nil {
						panic(err)
					}

					if emojiShortCodeID <= 0 {
						panic("Emoji ShortCode ID <= 0")
					}

					fmt.Printf("✅ Success! | Name: %s\n\t➜ URL: %s\n\t➜ S3 Location: %s\n\t➜ Emoji ID: %d\n\t➜ Emoji ShortCode: %d,\n", emoji.Name, emoji.PhotoURL, location, emojiID, emoji.Symbol)
				} else {
					panic("EmojiID <= 0")
				}
			}
		} else {
			panic(link)
		}
	}
}

// TODO: 💩 (непонятная ху*ня)
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
