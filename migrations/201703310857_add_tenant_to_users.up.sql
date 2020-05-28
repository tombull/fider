ALTER TABLE
    users
ADD
    COLUMN tenant_id INT;

ALTER TABLE
    users
ADD
    CONSTRAINT ideas_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE;

UPDATE
    users
SET
    tenant_id = ideas.tenant_id
FROM
    ideas
WHERE
    ideas.user_id = users.id;

UPDATE
    users
SET
    tenant_id = (
        SELECT
            id
        FROM
            tenants
        LIMIT
            1
    )
WHERE
    tenant_id IS NULL