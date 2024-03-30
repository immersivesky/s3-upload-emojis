package repository

import (
	"context"
	"errors"
	"github.com/immersivesky/s3-upload-emojis/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

var (
	errLessThan = errors.New("returned number less than 0")
)

func NewDB(dsn string) (*DB, error) {
	pool, err := pgxpool.New(context.TODO(), dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.TODO()); err != nil {
		return nil, err
	}

	return &DB{
		pool: pool,
	}, nil
}

func (db *DB) GetEmojiPacksBySource(source string, sourceID int) ([]*domain.EmojiPack, error) {
	emojiPacks := make([]*domain.EmojiPack, 0, 4)

	rows, err := db.pool.Query(
		context.TODO(),
		"SELECT emoji_pack_id, name FROM emoji_pack WHERE source_type = $1 AND source_id = $2;",
		source, sourceID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		emojiPack := new(domain.EmojiPack)

		if err := rows.Scan(&emojiPack.ID, &emojiPack.Name); err != nil {
			return nil, err
		}

		emojiPacks = append(emojiPacks, emojiPack)
	}

	return emojiPacks, nil
}

func (db *DB) GetEmojisByEmojiPack(emojiPackID, offset, count int) ([]*domain.Emoji, error) {
	emojis := make([]*domain.Emoji, 0, 4)

	rows, err := db.pool.Query(
		context.TODO(),
		"SELECT emoji_id, photo_path FROM emoji WHERE fk_emoji_pack_id = $1 OFFSET $2 LIMIT $3;",
		emojiPackID, offset, count,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		emoji := new(domain.Emoji)

		if err := rows.Scan(&emoji.ID, &emoji.PhotoPath); err != nil {
			return nil, err
		}

		emojis = append(emojis, emoji)
	}

	return emojis, nil
}

func (db *DB) GetEmojisByShortCode(shortCode string) ([]*domain.Emoji, error) {
	emojis := make([]*domain.Emoji, 0, 4)

	rows, err := db.pool.Query(
		context.TODO(),
		"SELECT photo_path FROM emoji WHERE emoji_id = (SELECT fk_emoji_id FROM emoji_shortcode WHERE shortcode = $1);",
		shortCode,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		emoji := new(domain.Emoji)

		if err := rows.Scan(&emoji.ID, &emoji.PhotoPath); err != nil {
			return nil, err
		}

		emojis = append(emojis, emoji)
	}

	return emojis, nil
}

func (db *DB) CreateEmojiPack(sourceType string, sourceID int, name, version string) (int, error) {
	var emojiPackID int

	if err := db.pool.QueryRow(
		context.TODO(),
		"INSERT INTO emoji_pack(source_type, source_id, name, version) VALUES($1, $2, $3, $4) RETURNING emoji_pack_id;",
		sourceType, sourceID, name, version,
	).Scan(&emojiPackID); err != nil {
		return emojiPackID, err
	}

	if emojiPackID <= 0 {
		return 0, errLessThan
	}

	return emojiPackID, nil
}

// CreateEmoji - запись эмодзи в БД
// Принимает пак эмодзи и относительный путь к фотографии
// Возвращает ID эмодзи
func (db *DB) CreateEmoji(emojiPackID int, photoPath string) (int, error) {
	var emojiID int

	if err := db.pool.QueryRow(
		context.TODO(),
		"INSERT INTO emoji(photo_path, fk_emoji_pack_id) VALUES ($1, $2) RETURNING emoji_id;",
		photoPath, emojiPackID,
	).Scan(&emojiID); err != nil {
		return emojiID, err
	}

	if emojiID <= 0 {
		return 0, errLessThan
	}

	return emojiID, nil
}

// CreateEmojiShortCode - запись короткого кода эмодзи в БД
// Принимает ID эмодзи и его символ
func (db *DB) CreateEmojiShortCode(emojiID int, shortCode string) (int, error) {
	var emojiShortCodeID int

	if err := db.pool.QueryRow(
		context.TODO(),
		"INSERT INTO emoji_shortcode(fk_emoji_id, shortcode) VALUES ($1, $2) RETURNING emoji_shortcode_id;",
		emojiID, shortCode,
	).Scan(&emojiShortCodeID); err != nil {
		return emojiShortCodeID, err
	}

	if emojiShortCodeID <= 0 {
		return 0, errLessThan
	}

	return emojiShortCodeID, nil
}
