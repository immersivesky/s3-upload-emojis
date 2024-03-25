package image

import (
	"image"
	"image/png"
	"net/http"
	"strings"
)

func GetImage(photoURL string) image.Image {
	res, err := http.Get(strings.ReplaceAll(photoURL, "\\", "/"))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	img, err := png.Decode(res.Body)
	if err != nil {
		panic(err)
	}

	return img
}
