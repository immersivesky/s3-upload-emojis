package convert

import (
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	"io"
)

func PNGToWEBP(img image.Image, writer io.Writer) {
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 100)
	if err != nil {
		panic(err)
	}

	if err := webp.Encode(writer, img, options); err != nil {
		panic(err)
	}
}
