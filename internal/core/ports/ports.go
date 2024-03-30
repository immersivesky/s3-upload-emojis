package ports

import "github.com/immersivesky/s3-upload-emojis/internal/core/domain"

type EmojiRepository interface {
	GetEmojiPacksBySource(sourceType string, sourceID int) ([]*domain.EmojiPack, error)
	GetEmojisByEmojiPack(emojiPackID, offset, count int) ([]*domain.Emoji, error)
	GetEmojisByShortCode(shortCode string) ([]*domain.Emoji, error)

	CreateEmojiPack(sourceType string, sourceID int, name, version string) (int, error)
	CreateEmoji(emojiPackID int, photoPath string) (int, error)
	CreateEmojiShortCode(emojiID int, shortCode string) (int, error)
}
