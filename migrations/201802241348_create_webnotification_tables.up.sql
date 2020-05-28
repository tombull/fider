CREATE TABLE IF NOT EXISTS notifications (
    id serial NOT NULL,
    tenant_id int NOT NULL,
    user_id int NOT NULL,
    title varchar(160) NOT NULL,
    link varchar(2048) NULL,
    READ boolean NOT NULL,
    idea_id int NOT NULL,
    author_id int NOT NULL,
    created_on timestamptz NOT NULL DEFAULT NOW(),
    updated_on timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (author_id) REFERENCES users(id),
    FOREIGN KEY (idea_id) REFERENCES ideas(id)
);