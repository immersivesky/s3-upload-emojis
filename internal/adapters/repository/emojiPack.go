package repository

func (db *DB) PushEmojiPack() {
	if _, err := db.CreateEmojiPack("community", 2, "Apple", "13.3"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 3, "Microsoft", "Windows 10 May 2019 Update"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 4, "Android", "11.0"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 5, "Facebook", "4.0"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 6, "Firefox", "2.5"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 7, "Messenger", "1.0"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 8, "Whats App", "2.19.352"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 9, "Twemoji", "13.0"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 10, "One UI", "2.5"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 11, "LG", "G5"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 12, "JoyPixels", "6.0"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 13, "Emojidex", "1.0.34"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 14, "OpenMoji", "12.2"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 15, "Softbank", "2014"); err != nil {
		panic(err)
	}

	if _, err := db.CreateEmojiPack("community", 16, "HTC", "Sense 8"); err != nil {
		panic(err)
	}
}
