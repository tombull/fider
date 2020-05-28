ALTER TABLE
    user_providers DROP COLUMN modified_on;

ALTER TABLE
    user_providers
ALTER COLUMN
    created_on DROP DEFAULT;