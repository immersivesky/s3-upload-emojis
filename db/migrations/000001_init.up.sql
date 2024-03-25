CREATE TYPE source AS ENUM(
    'user',
    'community',
    'chat'
);

CREATE TABLE emoji_pack(
    emoji_pack_id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    source source NOT NULL,
    source_id SMALLSERIAL NOT NULL
);

CREATE TABLE emoji(
    emoji_id BIGSERIAL PRIMARY KEY,
    photo_path TEXT NOT NULL,
    fk_emoji_pack_id SERIAL REFERENCES emoji_pack(emoji_pack_id)
);

CREATE TABLE emoji_shortcode(
    shortcode_id BIGSERIAL PRIMARY KEY,
    shortcode VARCHAR(64) NOT NULL,
    fk_emoji_id BIGSERIAL REFERENCES emoji(emoji_id)
);