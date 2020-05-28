ALTER TABLE
    tenants DROP COLUMN modified_on;

ALTER TABLE
    tenants
ALTER COLUMN
    created_on DROP DEFAULT;