package html

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/immersivesky/s3-upload-emojis/convert"
	"github.com/immersivesky/s3-upload-emojis/image"
	"github.com/immersivesky/s3-upload-emojis/internal/adapters/repository"
	"github.com/immersivesky/s3-upload-emojis/location"
	"github.com/immersivesky/s3-upload-emojis/s3"
	"go.uber.org/zap"
	"strings"
)

type Emoji struct {
	Symbol   string
	Name     string
	PhotoURL string
}

func newEmoji(symbolEmoji, name, photoURL string) Emoji {
	return Emoji{
		Symbol:   symbolEmoji,
		Name:     name,
		PhotoURL: photoURL,
	}
}

func On(s3 *s3.S3, db *repository.DB, location location.Location, logger *zap.Logger) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		var (
			spanImg = element.ChildAttr("span", "data-name")
			emoji   = newEmoji(
				element.ChildText("div.symbol-emoji"),
				getName(location.Name, spanImg),
				getPhotoURL(spanImg),
			)
			writer   bytes.Buffer
			path     = getPath(location.Name, location.Version, emoji.Name)
			img, err = image.GetImage(emoji.PhotoURL)
		)
		if err != nil {
			logger.Error("Error",
				zap.String("Message", err.Error()),
				zap.String("Name", emoji.Name),
				zap.String("S3 Path", path),
				zap.String("URL", emoji.PhotoURL),
			)
			return
		}

		convert.PNGToWEBP(img, &writer)
		s3.Upload(&writer, path)

		// чтобы получить ссылку на URL, необходимо сначала разрезать строку
		// при помощи символа -|/ или -| и получить срез 2 элементов
		// 1 элемент - название S3 бакета, 2 - путь к файлу
		// нужно это, чтобы в случае перехода на другой домен не пришлось изменять записи
		// мне к примеру удобно обходить CDN, поскольку он кэширует на 4 дня
		// а пользователям наоборот, нужен близлежащий к ним кэш
		// т.е. мы обращаемся к разным сервисам на разных поддоменах cdn.emojis и emojis
		emojiID, err := db.CreateEmoji(location.EmojiPackID, "emojis-|/"+path)
		if err != nil {
			panic(err)
		}

		_, err = db.CreateEmojiShortCode(emojiID, emoji.Symbol)
		if err != nil {
			panic(err)
		}


		/*logger.Info("Success",
			zap.String("Name", emoji.Name),
			zap.String("URL", emoji.PhotoURL),
			zap.String("S3 Path", path),
			zap.String("EmojiID", fmt.Sprint(emojiID)),
			zap.String("ShortCodeID", fmt.Sprint(emojiShortCodeID)),
			zap.String("ShortCode", emoji.Symbol),
		)*/ // это я убрал потому, что мне впадлу разбираться в ELK
	}
}

// TODO: 💩 (непонятная ху*ня)
func getName(vendorName, emojiPhotoURL string) (newPhotoURL string) {
	var (
		onVendor     = fmt.Sprintf("-on-%s", vendorName)
		lowerOnVedor = strings.ToLower(onVendor)
	)

	newPhotoURL = strings.SplitAfterN(strings.ToLower(emojiPhotoURL), fmt.Sprintf("%s\\", strings.ToLower(vendorName)), -1)[1]
	newPhotoURL = strings.SplitAfterN(newPhotoURL, lowerOnVedor, -1)[0]
	newPhotoURL = strings.ReplaceAll(newPhotoURL, lowerOnVedor, "")

	return newPhotoURL
}

func getPhotoURL(spanImg string) string {
	return strings.ReplaceAll(strings.SplitAfterN(spanImg, ".png", -1)[0], "\\", "/")
}

func getPath(name, version, emojiName string) string {
	return fmt.Sprintf("%s/%s/webp/%s.webp", strings.ToLower(name), version, emojiName)
}
