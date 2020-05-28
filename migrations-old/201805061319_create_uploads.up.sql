CREATE TABLE IF NOT EXISTS uploads (
    id serial NOT NULL,
    tenant_id int NOT NULL,
    size int NOT NULL,
    content_type varchar(200) NOT NULL,
    FILE bytea NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);