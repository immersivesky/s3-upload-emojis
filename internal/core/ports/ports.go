package ports

import "github.com/immersivesky/s3-upload-emojis/internal/core/domain"

type EmojiRepository interface {
	GetEmojiPack(source string, sourceID int) (*domain.EmojiPack, error)
	GetEmoji(packID, offset, count int) ([]*domain.Emoji, error)
	GetEmojisByShortCode(shortCode string) ([]*domain.Emoji, error)

	CreateEmojiPack(source string, sourceID int, name string) (int, error)
	CreateEmoji(emojiPackID int, photoPath string) (int, error)
	CreateEmojiShortCode(emojiID int, shortCode string) (int, error)
}
