ALTER TABLE
    tenants
ADD
    invitation varchar(100) NULL;

ALTER TABLE
    tenants
ADD
    welcome_message text NULL;

UPDATE
    tenants
SET
    invitation = '',
    welcome_message = '';