ALTER TABLE
    ideas
ADD
    supporters INT;

UPDATE
    ideas
SET
    supporters = 0;

ALTER TABLE
    ideas
ALTER COLUMN
    supporters
SET
    NOT NULL;

CREATE TABLE IF NOT EXISTS idea_supporters (
    user_id int NOT NULL,
    idea_id int NOT NULL,
    created_on timestamptz NOT NULL,
    PRIMARY KEY (user_id, idea_id),
    FOREIGN KEY (idea_id) REFERENCES ideas(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);