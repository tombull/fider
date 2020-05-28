CREATE TABLE IF NOT EXISTS LOGS (
    id serial NOT NULL,
    tag varchar(50) NOT NULL,
    LEVEL varchar(50) NOT NULL,
    text text NOT NULL,
    properties jsonb NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);