CREATE TABLE IF NOT EXISTS users (
	_id SERIAL PRIMARY KEY,
	uuid VARCHAR(36) NOT NULL,
	login VARCHAR(250) NOT NULL,
	password VARCHAR(250) NOT NULL,
	is_active BOOLEAN DEFAULT true
);

CREATE UNIQUE INDEX IF NOT EXISTS login_idx ON users (login);

CREATE TABLE IF NOT EXISTS secure_notes (
	_id SERIAL,
	uuid VARCHAR(250) PRIMARY KEY NOT NULL,
	type   NUMERIC DEFAULT 0,
	title VARCHAR(250) NOT NULL,
	description  VARCHAR(250),
	username VARCHAR(250),
	password VARCHAR(250),
	text_note VARCHAR(250),
	raw_note BYTEA,
	card_number VARCHAR(250),
	card_holder VARCHAR(250),
	card_cvc VARCHAR(250),
	card_expiry TIMESTAMPTZ NOT NULL,
	created_by VARCHAR(36) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	last_edited_at TIMESTAMPTZ
);
