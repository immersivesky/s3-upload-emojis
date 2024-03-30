package location

type Location struct {
	EmojiPackID int
	URL         string
	Name        string
	Version     string
}

func GetLocations() []Location {
	locations := []Location{
		{
			URL:         "https://emojik.com/apple",
			Name:        "Apple",
			Version:     "13.3",
			EmojiPackID: 1,
		},
		{
			URL:         "https://emojik.com/microsoft",
			Name:        "Microsoft",
			Version:     "Windows 10 May 2019 Update",
			EmojiPackID: 2,
		},
		{
			URL:         "https://emojik.com/google",
			Name:        "Google",
			Version:     "11.0",
			EmojiPackID: 3,
		},
		{
			URL:         "https://emojik.com/facebook",
			Name:        "Facebook",
			Version:     "4.0",
			EmojiPackID: 4,
		},
		{
			URL:         "https://emojik.com/mozilla",
			Name:        "Mozilla",
			Version:     "2.5",
			EmojiPackID: 5,
		},
		{
			URL:         "https://emojik.com/messenger",
			Name:        "Messenger",
			Version:     "1.0",
			EmojiPackID: 6,
		},
		{
			URL:         "https://emojik.com/whatsapp",
			Name:        "WhatsApp",
			Version:     "2.19.352",
			EmojiPackID: 7,
		},
		{
			URL:         "https://emojik.com/twitter",
			Name:        "Twitter",
			Version:     "13.0",
			EmojiPackID: 8,
		},
		{
			URL:         "https://emojik.com/samsung",
			Name:        "Samsung",
			Version:     "2.5",
			EmojiPackID: 9,
		},
		{
			URL:         "https://emojik.com/lg",
			Name:        "LG",
			Version:     "G5",
			EmojiPackID: 10,
		},
		{
			URL:         "https://emojik.com/joypixels",
			Name:        "JoyPixels",
			Version:     "6.0",
			EmojiPackID: 11,
		},
		{
			URL:         "https://emojik.com/emojidex",
			Name:        "Emojidex",
			Version:     "1.0.34",
			EmojiPackID: 12,
		},
		{
			URL:         "https://emojik.com/openmoji",
			Name:        "OpenMoji",
			Version:     "12.2",
			EmojiPackID: 13,
		},
		{
			URL:         "https://emojik.com/softbank",
			Name:        "Softbank",
			Version:     "2014",
			EmojiPackID: 14,
		},
		{
			URL:         "https://emojik.com/htc",
			Name:        "HTC",
			Version:     "Sense 8",
			EmojiPackID: 15,
		},
	}

	return locations
}
