INSERT INTO
    blobs (
        KEY,
        tenant_id,
        size,
        content_type,
        FILE,
        created_at,
        modified_at
    )
SELECT
    'logos/' || md5(cast(id AS varchar)) || '.' || split_part(content_type, '/', 2) AS KEY,
    tenant_id,
    size,
    content_type,
    FILE,
    created_at,
    created_at AS modified_at
FROM
    uploads;

ALTER TABLE
    tenants
ADD
    COLUMN logo_bkey varchar(512) NULL;

ALTER TABLE
    oauth_providers
ADD
    COLUMN logo_bkey varchar(512) NULL;

UPDATE
    tenants
SET
    logo_bkey = 'logos/' || md5(cast(u.id AS varchar)) || '.' || split_part(u.content_type, '/', 2)
FROM
    uploads u
WHERE
    u.id = tenants.logo_id;

UPDATE
    oauth_providers
SET
    logo_bkey = 'logos/' || md5(cast(u.id AS varchar)) || '.' || split_part(u.content_type, '/', 2)
FROM
    uploads u
WHERE
    u.id = oauth_providers.logo_id;

ALTER TABLE
    tenants DROP COLUMN logo_id;

ALTER TABLE
    oauth_providers DROP COLUMN logo_id;

UPDATE
    tenants
SET
    logo_bkey = ''
WHERE
    logo_bkey IS NULL;

UPDATE
    oauth_providers
SET
    logo_bkey = ''
WHERE
    logo_bkey IS NULL;

ALTER TABLE
    tenants
ALTER COLUMN
    logo_bkey
SET
    NOT NULL;

ALTER TABLE
    oauth_providers
ALTER COLUMN
    logo_bkey
SET
    NOT NULL;

DROP TABLE uploads;