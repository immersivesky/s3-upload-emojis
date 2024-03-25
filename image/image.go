package image

import (
	"image"
	"image/png"
	"net/http"
)

func GetImage(photoURL string) (image.Image, error) {
	res, err := http.Get(photoURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, err := png.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
