ALTER TABLE
    users
ADD
    role INT;

UPDATE
    users
SET
    role = 1;

ALTER TABLE
    users
ALTER COLUMN
    role
SET
    NOT NULL;