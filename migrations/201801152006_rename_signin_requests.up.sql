ALTER TABLE
    signin_requests RENAME TO email_verifications;

ALTER INDEX signin_requests_pkey RENAME TO email_verifications_pkey;

ALTER INDEX signin_requests_key_idx RENAME TO email_verifications_key_idx;