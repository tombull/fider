CREATE TABLE IF NOT EXISTS comments (
    id serial PRIMARY KEY,
    content text NULL,
    idea_id int NOT NULL,
    user_id int NOT NULL,
    created_on timestamptz NOT NULL,
    FOREIGN KEY (idea_id) REFERENCES ideas(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);