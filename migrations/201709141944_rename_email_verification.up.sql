ALTER TABLE
    email_verifications RENAME TO signin_requests;

ALTER INDEX email_verifications_pkey RENAME TO signin_requests_pkey;

ALTER INDEX email_verification_key_idx RENAME TO signin_requests_key_idx;