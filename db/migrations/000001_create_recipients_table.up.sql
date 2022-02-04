BEGIN;

CREATE TABLE recipient (
    id           INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    twitter_id   TEXT,
    facebook_id  TEXT,
    instagram_id TEXT,
    display_name TEXT,
    username     TEXT,

    platform_created_at TIMESTAMP WITHOUT TIME ZONE      ,        
    updated_at          TIMESTAMP WITHOUT TIME ZONE,
    created_at          TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

COMMIT;