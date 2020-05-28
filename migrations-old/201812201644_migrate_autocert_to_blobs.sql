CREATE UNIQUE INDEX blobs_unique_global_key ON blobs (KEY, tenant_id)
WHERE
    tenant_id IS NOT NULL;

CREATE UNIQUE INDEX blobs_unique_tenant_key ON blobs (KEY)
WHERE
    tenant_id IS NULL;