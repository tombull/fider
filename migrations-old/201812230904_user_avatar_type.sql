ALTER TABLE
    users
ADD
    COLUMN avatar_type smallint NULL;

ALTER TABLE
    users
ADD
    COLUMN avatar_bkey varchar(512) NULL;

UPDATE
    users
SET
    avatar_type = 2;

-- gravatar
UPDATE
    users
SET
    avatar_bkey = '';

ALTER TABLE
    users
ALTER COLUMN
    avatar_type
SET
    NOT NULL;

ALTER TABLE
    users
ALTER COLUMN
    avatar_bkey
SET
    NOT NULL;