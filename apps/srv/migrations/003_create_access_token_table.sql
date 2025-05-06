-- Write your migrate up statements here

CREATE TABLE access_tokens (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    token_identifier UUID NOT NULL,
    issued_at TIMESTAMP WITH TIME ZONE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON COLUMN access_tokens.token_identifier IS 'The token identifier identifies a unique token (JTI).';

---- create above / drop below ----

DROP TABLE IF EXISTS access_tokens;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
