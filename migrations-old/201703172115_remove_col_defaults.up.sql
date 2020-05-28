ALTER TABLE
    ideas DROP COLUMN modified_on;

ALTER TABLE
    ideas
ALTER COLUMN
    created_on DROP DEFAULT;

ALTER TABLE
    users DROP COLUMN modified_on;

ALTER TABLE
    users
ALTER COLUMN
    created_on DROP DEFAULT;