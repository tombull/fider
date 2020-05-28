ALTER TABLE
    ideas
ADD
    response text NULL;

ALTER TABLE
    ideas
ADD
    response_user_id int NULL;

ALTER TABLE
    ideas
ADD
    response_date timestamptz NULL;