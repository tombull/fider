UPDATE
    users
SET
    email = lower(email)
WHERE
    email != lower(email)