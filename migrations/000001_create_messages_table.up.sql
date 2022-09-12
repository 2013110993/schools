CREATE TABLE IF NOT EXISTS  messages (
    id bigserial PRIMARY KEY,
    create_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    message text NOT NULL,
    version integer NOT NULL DEFAULT 1
);