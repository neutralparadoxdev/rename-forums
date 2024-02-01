CREATE TABLE forum(
	title TEXT PRIMARY KEY,
	description TEXT NOT NULL DEFAULT '',
	is_public BOOLEAN NOT NULL DEFAULT False
);

CREATE TABLE goforum_user(
	id INTEGER PRIMARY KEY,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	email TEXT NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_login TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	locked BOOLEAN NOT NULL DEFAULT False
);

CREATE TABLE forum_goforum_user_owner(
	forum_title TEXT NOT NULL REFERENCES forum(title),
	user_id INTEGER NOT NULL REFERENCES goforum_user(id),
	PRIMARY KEY(forum_title, user_id)
);

CREATE TABLE forum_goforum_user_joined(
	forum_title TEXT NOT NULL REFERENCES forum(title),
	user_id INTEGER NOT NULL REFERENCES goforum_user(id),
	PRIMARY KEY(forum_title, user_id)
);
