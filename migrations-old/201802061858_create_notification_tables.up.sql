CREATE TABLE IF NOT EXISTS idea_subscribers (
    user_id int NOT NULL,
    idea_id int NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    updated_on timestamptz NOT NULL DEFAULT NOW(),
    STATUS smallint NOT NULL,
    PRIMARY KEY (user_id, idea_id),
    FOREIGN KEY (idea_id) REFERENCES ideas(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_settings (
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    KEY varchar(100) NOT NULL,
    value varchar(100) NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX user_settings_uq_key ON user_settings (user_id, KEY);