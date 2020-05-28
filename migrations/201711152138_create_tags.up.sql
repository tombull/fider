CREATE TABLE IF NOT EXISTS tags (
    id serial PRIMARY KEY,
    tenant_id int NOT NULL,
    name varchar(30) NOT NULL,
    slug varchar(30) NOT NULL,
    color varchar(6) NOT NULL,
    is_public boolean NOT NULL,
    created_on timestamptz NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);

CREATE TABLE IF NOT EXISTS idea_tags (
    tag_id int NOT NULL,
    idea_id int NOT NULL,
    created_on timestamptz NOT NULL,
    created_by_id int NOT NULL,
    PRIMARY KEY (tag_id, idea_id),
    FOREIGN KEY (idea_id) REFERENCES ideas(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    FOREIGN KEY (created_by_id) REFERENCES users(id)
);