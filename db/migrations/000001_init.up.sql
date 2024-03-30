CREATE TYPE source_type AS ENUM(
    'user',
    'community',
    'chat'
);

CREATE TABLE emoji_pack(
    emoji_pack_id SERIAL UNIQUE,
    source_type source_type NOT NULL,
    source_id BIGSERIAL NOT NULL,
    name VARCHAR(128) NOT NULL,
    version VARCHAR(128) NOT NULL,
    PRIMARY KEY (emoji_pack_id, source_id)
);

CREATE TABLE emoji(
    emoji_id BIGSERIAL PRIMARY KEY,
    photo_path TEXT NOT NULL,
    fk_emoji_pack_id BIGSERIAL REFERENCES emoji_pack(emoji_pack_id)
);

CREATE TABLE emoji_shortcode(
    emoji_shortcode_id BIGSERIAL PRIMARY KEY,
    shortcode VARCHAR(64) NOT NULL,
    fk_emoji_id BIGSERIAL REFERENCES emoji(emoji_id)
);