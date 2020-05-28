ALTER TABLE
    ideas
ADD
    COLUMN user_id INT;

ALTER TABLE
    ideas
ADD
    CONSTRAINT ideas_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;