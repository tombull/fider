CREATE TABLE IF NOT EXISTS tenants_billing (
    tenant_id int NOT NULL,
    trial_ends_at timestamptz NOT NULL,
    subscription_ends_at timestamptz NULL,
    stripe_customer_id varchar(255) NULL,
    stripe_subscription_id varchar(255) NULL,
    stripe_plan_id varchar(255) NULL,
    PRIMARY KEY (tenant_id),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
);