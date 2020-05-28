CREATE TABLE IF NOT EXISTS tenants (
    id serial PRIMARY KEY,
    name varchar(60) NOT NULL,
    domain varchar(40) NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    modified_on timestamptz NOT NULL DEFAULT NOW()
);