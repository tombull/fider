CREATE TABLE IF NOT EXISTS email_verifications (
    id serial PRIMARY KEY,
    tenant_id int NOT NULL,
    email varchar(200) NOT NULL,
    created_on timestamptz NOT NULL,
    KEY varchar(32) NOT NULL,
    verified_on timestamptz NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);