CREATE TABLE IF NOT EXISTS blobs (
    id serial NOT NULL,
    KEY varchar(512) NOT NULL,
    tenant_id int NULL,
    size bigint NOT NULL,
    content_type varchar(200) NOT NULL,
    FILE bytea NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    modified_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    UNIQUE (tenant_id, KEY),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);