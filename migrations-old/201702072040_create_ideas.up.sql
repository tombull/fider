CREATE TABLE IF NOT EXISTS ideas (
    id serial PRIMARY KEY,
    title varchar(100) NOT NULL,
    description text NULL,
    tenant_id int NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    modified_on timestamptz NOT NULL DEFAULT NOW(),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);