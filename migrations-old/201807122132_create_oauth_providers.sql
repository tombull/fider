CREATE TABLE IF NOT EXISTS oauth_providers (
    id serial NOT NULL,
    tenant_id int NOT NULL,
    logo_id int NULL,
    provider varchar(30) NOT NULL,
    display_name varchar(50) NOT NULL,
    STATUS int NOT NULL,
    client_id varchar(100) NOT NULL,
    client_secret varchar(500) NOT NULL,
    authorize_url varchar(300) NOT NULL,
    token_url varchar(300) NOT NULL,
    profile_url varchar(300) NOT NULL,
    scope varchar(100) NOT NULL,
    json_user_id_path varchar(100) NOT NULL,
    json_user_name_path varchar(100) NOT NULL,
    json_user_email_path varchar(100) NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    FOREIGN KEY (logo_id, tenant_id) REFERENCES uploads(id, tenant_id)
);

CREATE UNIQUE INDEX tenant_id_provider_key ON oauth_providers (tenant_id, provider);