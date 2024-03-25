package repository

import (
	"context"
	"github.com/immersivesky/s3-upload-emojis/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

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

func (db *DB) GetEmojiPack(source string, sourceID int) (*domain.EmojiPack, error) {
	emojiPack := &domain.EmojiPack{}

	if err := db.pool.QueryRow(
		context.TODO(),
		"SELECT emoji_pack_id, name FROM emoji_pack WHERE source = $1 AND source_id = $1;",
		source, sourceID,
	).Scan(&emojiPack.ID, &emojiPack.Name); err != nil {
		return nil, err
	}

	return emojiPack, nil
}

func (db *DB) GetEmoji(packID, offset, count int) ([]*domain.Emoji, error) {
	emojis := []*domain.Emoji{}

	rows, err := db.pool.Query(
		context.TODO(),
		"SELECT emoji_id, photo_path FROM emoji WHERE fk_emoji_pack_id = $1 OFFSET $2 LIMIT $3;",
		packID, offset, count,
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
	emojis := []*domain.Emoji{}

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

func (db *DB) CreateEmojiPack(source string, sourceID int, name string) (int, error) {
	var emojiPackID int

	if err := db.pool.QueryRow(
		context.TODO(),
		"INSERT INTO emoji_pack(source, source_id, name) VALUES($1, $2, $3) RETURNING emoji_pack_id;",
		source, sourceID, name,
	).Scan(&emojiPackID); err != nil {
		return emojiPackID, err
	}

	return emojiPackID, nil
}

func (db *DB) CreateEmoji(emojiPackID int, photoPath string) (int, error) {
	var emojiID int

	if err := db.pool.QueryRow(
		context.TODO(),
		"INSERT INTO emoji(fk_emoji_pack_id, photo_path) VALUES ($1, $2) RETURNING emoji_id;",
		emojiPackID, photoPath,
	).Scan(&emojiID); err != nil {
		return emojiID, err
	}

	return emojiID, nil
}

func (db *DB) CreateEmojiShortCode(emojiID int, shortCode string) (int, error) {
	var emojiShortCodeID int

	if err := db.pool.QueryRow(
		context.TODO(),
		"INSERT INTO emoji_shortcode(fk_emoji_id, shortcode) VALUES ($1, $2) RETURNING ;",
		emojiID, shortCode,
	).Scan(&emojiShortCodeID); err != nil {
		return emojiShortCodeID, err
	}

	return emojiShortCodeID, nil
}
