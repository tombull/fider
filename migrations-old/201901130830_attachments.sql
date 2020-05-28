CREATE TABLE IF NOT EXISTS attachments (
    id serial NOT NULL,
    tenant_id int NOT NULL,
    post_id int NOT NULL,
    comment_id int NULL,
    user_id int NOT NULL,
    attachment_bkey varchar(512) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (comment_id) REFERENCES comments(id)
);