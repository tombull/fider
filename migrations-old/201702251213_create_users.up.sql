CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(100) NULL,
    email varchar(200) NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    modified_on timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_providers (
    user_id int NOT NULL,
    provider varchar(40) NOT NULL,
    provider_uid varchar(100) NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    modified_on timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, provider),
    FOREIGN KEY (user_id) REFERENCES users(id)
);