ALTER TABLE
    signin_requests
ADD
    name VARCHAR(200) NULL;

ALTER TABLE
    signin_requests
ADD
    expires_on TIMESTAMPTZ NOT NULL DEFAULT NOW();

UPDATE
    signin_requests
SET
    expires_on = created_on + INTERVAL '15 minute';

ALTER TABLE
    signin_requests
ALTER COLUMN
    expires_on DROP DEFAULT;