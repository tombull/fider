CREATE TABLE public.attachments (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    post_id integer NOT NULL,
    comment_id integer,
    user_id integer NOT NULL,
    attachment_bkey character varying(512) NOT NULL
);

CREATE SEQUENCE public.attachments_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.attachments_id_seq OWNED BY public.attachments.id;

CREATE TABLE public.blobs (
    id integer NOT NULL,
    KEY character varying(512) NOT NULL,
    tenant_id integer,
    size bigint NOT NULL,
    content_type character varying(200) NOT NULL,
    FILE bytea NOT NULL,
    created_at timestamp WITH time zone DEFAULT NOW() NOT NULL,
    modified_at timestamp WITH time zone DEFAULT NOW() NOT NULL
);

CREATE SEQUENCE public.blobs_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.blobs_id_seq OWNED BY public.blobs.id;

CREATE TABLE public.comments (
    id integer NOT NULL,
    content text,
    post_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    tenant_id integer NOT NULL,
    edited_at timestamp WITH time zone,
    edited_by_id integer,
    deleted_at timestamp WITH time zone,
    deleted_by_id integer
);

CREATE SEQUENCE public.comments_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;

CREATE TABLE public.email_verifications (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    email character varying(200) NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    KEY character varying(64) NOT NULL,
    verified_at timestamp WITH time zone,
    name character varying(200),
    expires_at timestamp WITH time zone NOT NULL,
    kind smallint NOT NULL,
    user_id integer
);

CREATE SEQUENCE public.email_verifications_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.email_verifications_id_seq OWNED BY public.email_verifications.id;

CREATE TABLE public.events (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    client_ip inet,
    name character varying(64) NOT NULL,
    created_at timestamp WITH time zone DEFAULT NOW() NOT NULL
);

CREATE SEQUENCE public.events_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;

CREATE TABLE public.logs (
    id integer NOT NULL,
    tag character varying(50) NOT NULL,
    LEVEL character varying(50) NOT NULL,
    text text NOT NULL,
    properties jsonb,
    created_at timestamp WITH time zone DEFAULT NOW() NOT NULL
);

CREATE SEQUENCE public.logs_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.logs_id_seq OWNED BY public.logs.id;

CREATE TABLE public.notifications (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    user_id integer NOT NULL,
    title character varying(400) NOT NULL,
    link character varying(2048),
    READ boolean NOT NULL,
    post_id integer NOT NULL,
    author_id integer NOT NULL,
    created_at timestamp WITH time zone DEFAULT NOW() NOT NULL,
    updated_at timestamp WITH time zone DEFAULT NOW() NOT NULL
);

CREATE SEQUENCE public.notifications_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.notifications_id_seq OWNED BY public.notifications.id;

CREATE TABLE public.oauth_providers (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    provider character varying(30) NOT NULL,
    display_name character varying(50) NOT NULL,
    STATUS integer NOT NULL,
    client_id character varying(100) NOT NULL,
    client_secret character varying(500) NOT NULL,
    authorize_url character varying(300) NOT NULL,
    token_url character varying(300) NOT NULL,
    profile_url character varying(300) NOT NULL,
    scope character varying(100) NOT NULL,
    json_user_id_path character varying(100) NOT NULL,
    json_user_name_path character varying(100) NOT NULL,
    json_user_email_path character varying(100) NOT NULL,
    created_at timestamp WITH time zone DEFAULT NOW() NOT NULL,
    logo_bkey character varying(512) NOT NULL
);

CREATE SEQUENCE public.oauth_providers_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.oauth_providers_id_seq OWNED BY public.oauth_providers.id;

CREATE TABLE public.post_subscribers (
    user_id integer NOT NULL,
    post_id integer NOT NULL,
    created_at timestamp WITH time zone DEFAULT NOW() NOT NULL,
    updated_at timestamp WITH time zone DEFAULT NOW() NOT NULL,
    STATUS smallint NOT NULL,
    tenant_id integer NOT NULL
);

CREATE TABLE public.post_tags (
    tag_id integer NOT NULL,
    post_id integer NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    created_by_id integer NOT NULL,
    tenant_id integer NOT NULL
);

CREATE TABLE public.post_votes (
    user_id integer NOT NULL,
    post_id integer NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    tenant_id integer NOT NULL
);

CREATE TABLE public.posts (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    description text,
    tenant_id integer NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    user_id integer,
    number integer NOT NULL,
    STATUS integer NOT NULL,
    slug character varying(100) NOT NULL,
    response text,
    response_user_id integer,
    response_date timestamp WITH time zone,
    original_id integer
);

CREATE SEQUENCE public.posts_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;

CREATE TABLE public.tags (
    id integer NOT NULL,
    tenant_id integer NOT NULL,
    name character varying(30) NOT NULL,
    slug character varying(30) NOT NULL,
    color character varying(6) NOT NULL,
    is_public boolean NOT NULL,
    created_at timestamp WITH time zone NOT NULL
);

CREATE SEQUENCE public.tags_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;

CREATE TABLE public.tenants (
    id integer NOT NULL,
    name character varying(60) NOT NULL,
    subdomain character varying(40) NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    cname character varying(100),
    invitation character varying(100),
    welcome_message text,
    STATUS integer NOT NULL,
    is_private boolean NOT NULL,
    custom_css text NOT NULL,
    logo_bkey character varying(512) NOT NULL
);

CREATE TABLE public.tenants_billing (
    tenant_id integer NOT NULL,
    trial_ends_at timestamp WITH time zone NOT NULL,
    subscription_ends_at timestamp WITH time zone,
    stripe_customer_id character varying(255),
    stripe_subscription_id character varying(255),
    stripe_plan_id character varying(255)
);

CREATE SEQUENCE public.tenants_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.tenants_id_seq OWNED BY public.tenants.id;

CREATE TABLE public.user_providers (
    user_id integer NOT NULL,
    provider character varying(40) NOT NULL,
    provider_uid character varying(100) NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    tenant_id integer NOT NULL
);

CREATE TABLE public.user_settings (
    id integer NOT NULL,
    user_id integer NOT NULL,
    KEY character varying(100) NOT NULL,
    value character varying(100),
    tenant_id integer NOT NULL
);

CREATE SEQUENCE public.user_settings_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.user_settings_id_seq OWNED BY public.user_settings.id;

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100),
    email character varying(200) NOT NULL,
    created_at timestamp WITH time zone NOT NULL,
    tenant_id integer,
    role integer NOT NULL,
    STATUS integer NOT NULL,
    api_key character varying(64),
    api_key_date timestamp WITH time zone,
    avatar_type smallint NOT NULL,
    avatar_bkey character varying(512) NOT NULL
);

CREATE SEQUENCE public.users_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;

ALTER TABLE
    ONLY public.attachments
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.attachments_id_seq' :: regclass);

ALTER TABLE
    ONLY public.blobs
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.blobs_id_seq' :: regclass);

ALTER TABLE
    ONLY public.comments
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.comments_id_seq' :: regclass);

ALTER TABLE
    ONLY public.email_verifications
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.email_verifications_id_seq' :: regclass);

ALTER TABLE
    ONLY public.events
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.events_id_seq' :: regclass);

ALTER TABLE
    ONLY public.logs
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.logs_id_seq' :: regclass);

ALTER TABLE
    ONLY public.notifications
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.notifications_id_seq' :: regclass);

ALTER TABLE
    ONLY public.oauth_providers
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.oauth_providers_id_seq' :: regclass);

ALTER TABLE
    ONLY public.posts
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.posts_id_seq' :: regclass);

ALTER TABLE
    ONLY public.tags
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.tags_id_seq' :: regclass);

ALTER TABLE
    ONLY public.tenants
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.tenants_id_seq' :: regclass);

ALTER TABLE
    ONLY public.user_settings
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.user_settings_id_seq' :: regclass);

ALTER TABLE
    ONLY public.users
ALTER COLUMN
    id
SET
    DEFAULT nextval('public.users_id_seq' :: regclass);

ALTER TABLE
    ONLY public.attachments
ADD
    CONSTRAINT attachments_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.blobs
ADD
    CONSTRAINT blobs_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.blobs
ADD
    CONSTRAINT blobs_tenant_id_key_key UNIQUE (tenant_id, KEY);

ALTER TABLE
    ONLY public.comments
ADD
    CONSTRAINT comments_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.email_verifications
ADD
    CONSTRAINT email_verifications_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.events
ADD
    CONSTRAINT events_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.logs
ADD
    CONSTRAINT logs_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.notifications
ADD
    CONSTRAINT notifications_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.oauth_providers
ADD
    CONSTRAINT oauth_providers_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.post_subscribers
ADD
    CONSTRAINT post_subscribers_pkey PRIMARY KEY (user_id, post_id);

ALTER TABLE
    ONLY public.post_tags
ADD
    CONSTRAINT post_tags_pkey PRIMARY KEY (tag_id, post_id);

ALTER TABLE
    ONLY public.post_votes
ADD
    CONSTRAINT post_votes_pkey PRIMARY KEY (user_id, post_id);

ALTER TABLE
    ONLY public.posts
ADD
    CONSTRAINT posts_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.tags
ADD
    CONSTRAINT tags_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.tenants_billing
ADD
    CONSTRAINT tenants_billing_pkey PRIMARY KEY (tenant_id);

ALTER TABLE
    ONLY public.tenants
ADD
    CONSTRAINT tenants_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.user_providers
ADD
    CONSTRAINT user_providers_pkey PRIMARY KEY (user_id, provider);

ALTER TABLE
    ONLY public.user_settings
ADD
    CONSTRAINT user_settings_pkey PRIMARY KEY (id);

ALTER TABLE
    ONLY public.users
ADD
    CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX blobs_unique_global_key ON public.blobs USING btree (KEY, tenant_id)
WHERE
    (tenant_id IS NOT NULL);

CREATE UNIQUE INDEX blobs_unique_tenant_key ON public.blobs USING btree (KEY)
WHERE
    (tenant_id IS NULL);

CREATE UNIQUE INDEX email_verifications_key_idx ON public.email_verifications USING btree (tenant_id, KEY);

CREATE UNIQUE INDEX post_id_tenant_id_key ON public.posts USING btree (tenant_id, id);

CREATE UNIQUE INDEX post_number_tenant_key ON public.posts USING btree (tenant_id, number);

CREATE UNIQUE INDEX post_slug_tenant_key ON public.posts USING btree (tenant_id, slug)
WHERE
    (STATUS <> 6);

CREATE UNIQUE INDEX tag_id_tenant_id_key ON public.tags USING btree (tenant_id, id);

CREATE UNIQUE INDEX tag_slug_tenant_key ON public.tags USING btree (tenant_id, slug);

CREATE UNIQUE INDEX tenant_cname_unique_idx ON public.tenants USING btree (cname)
WHERE
    ((cname) :: text <> '' :: text);

CREATE UNIQUE INDEX tenant_id_provider_key ON public.oauth_providers USING btree (tenant_id, provider);

CREATE UNIQUE INDEX tenant_subdomain_unique_idx ON public.tenants USING btree (subdomain);

CREATE UNIQUE INDEX user_email_unique_idx ON public.users USING btree (tenant_id, email)
WHERE
    ((email) :: text <> '' :: text);

CREATE UNIQUE INDEX user_id_tenant_id_key ON public.users USING btree (tenant_id, id);

CREATE UNIQUE INDEX user_provider_unique_idx ON public.user_providers USING btree (user_id, provider);

CREATE UNIQUE INDEX user_settings_uq_key ON public.user_settings USING btree (user_id, KEY);

CREATE UNIQUE INDEX users_api_key ON public.users USING btree (tenant_id, api_key);

ALTER TABLE
    ONLY public.attachments
ADD
    CONSTRAINT attachments_comment_id_fkey FOREIGN KEY (comment_id) REFERENCES public.comments(id);

ALTER TABLE
    ONLY public.attachments
ADD
    CONSTRAINT attachments_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id);

ALTER TABLE
    ONLY public.attachments
ADD
    CONSTRAINT attachments_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.attachments
ADD
    CONSTRAINT attachments_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);

ALTER TABLE
    ONLY public.blobs
ADD
    CONSTRAINT blobs_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.comments
ADD
    CONSTRAINT comments_deleted_by_id_fkey FOREIGN KEY (deleted_by_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.comments
ADD
    CONSTRAINT comments_edited_by_id_fkey FOREIGN KEY (edited_by_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.comments
ADD
    CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id, tenant_id) REFERENCES public.posts(id, tenant_id);

ALTER TABLE
    ONLY public.comments
ADD
    CONSTRAINT comments_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.comments
ADD
    CONSTRAINT comments_user_id_fkey FOREIGN KEY (user_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.email_verifications
ADD
    CONSTRAINT email_verifications_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.email_verifications
ADD
    CONSTRAINT email_verifications_user_id_fkey FOREIGN KEY (user_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.events
ADD
    CONSTRAINT events_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.users
ADD
    CONSTRAINT ideas_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id) ON DELETE CASCADE;

ALTER TABLE
    ONLY public.notifications
ADD
    CONSTRAINT notifications_author_id_fkey FOREIGN KEY (author_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.notifications
ADD
    CONSTRAINT notifications_post_id_fkey FOREIGN KEY (post_id, tenant_id) REFERENCES public.posts(id, tenant_id);

ALTER TABLE
    ONLY public.notifications
ADD
    CONSTRAINT notifications_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.notifications
ADD
    CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.oauth_providers
ADD
    CONSTRAINT oauth_providers_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.post_subscribers
ADD
    CONSTRAINT post_subscribers_post_id_fkey FOREIGN KEY (post_id, tenant_id) REFERENCES public.posts(id, tenant_id);

ALTER TABLE
    ONLY public.post_subscribers
ADD
    CONSTRAINT post_subscribers_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.post_subscribers
ADD
    CONSTRAINT post_subscribers_user_id_fkey FOREIGN KEY (user_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.post_tags
ADD
    CONSTRAINT post_tags_created_by_id_fkey FOREIGN KEY (created_by_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.post_tags
ADD
    CONSTRAINT post_tags_post_id_fkey FOREIGN KEY (post_id, tenant_id) REFERENCES public.posts(id, tenant_id);

ALTER TABLE
    ONLY public.post_tags
ADD
    CONSTRAINT post_tags_tag_id_fkey FOREIGN KEY (tag_id, tenant_id) REFERENCES public.tags(id, tenant_id);

ALTER TABLE
    ONLY public.post_tags
ADD
    CONSTRAINT post_tags_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.post_votes
ADD
    CONSTRAINT post_votes_post_id_fkey FOREIGN KEY (post_id, tenant_id) REFERENCES public.posts(id, tenant_id);

ALTER TABLE
    ONLY public.post_votes
ADD
    CONSTRAINT post_votes_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.post_votes
ADD
    CONSTRAINT post_votes_user_id_fkey FOREIGN KEY (user_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.posts
ADD
    CONSTRAINT posts_original_id_fkey FOREIGN KEY (original_id) REFERENCES public.posts(id);

ALTER TABLE
    ONLY public.posts
ADD
    CONSTRAINT posts_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.posts
ADD
    CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE
    ONLY public.tags
ADD
    CONSTRAINT tags_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.tenants_billing
ADD
    CONSTRAINT tenants_billing_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.user_providers
ADD
    CONSTRAINT user_providers_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.user_providers
ADD
    CONSTRAINT user_providers_user_id_fkey FOREIGN KEY (user_id, tenant_id) REFERENCES public.users(id, tenant_id);

ALTER TABLE
    ONLY public.user_settings
ADD
    CONSTRAINT user_settings_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);

ALTER TABLE
    ONLY public.user_settings
ADD
    CONSTRAINT user_settings_user_id_fkey FOREIGN KEY (user_id, tenant_id) REFERENCES public.users(id, tenant_id);