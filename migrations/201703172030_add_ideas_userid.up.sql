ALTER TABLE
    ideas
ADD
    COLUMN user_id INT;

ALTER TABLE
    ideas
ADD
    CONSTRAINT users_fk FOREIGN KEY user_id REFERENCES users (id) ON DELETE CASCADE;