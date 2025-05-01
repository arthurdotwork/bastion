-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL DEFAULT pg_catalog.gen_random_uuid(),
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (NOW() at time zone 'UTC'),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (NOW() at time zone 'UTC'),
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

-- uq username when not empty
CREATE UNIQUE INDEX IF NOT EXISTS users_username_unique ON users (username) WHERE username != '';
CREATE UNIQUE INDEX IF NOT EXISTS users_email_unique ON users (email);

---- create above / drop below ----

DROP TABLE IF EXISTS users;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
