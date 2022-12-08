
-- -------------------------------------DB-----------------------------------------------start
CREATE DATABASE mygram
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
-- -------------------------------------TABLE::USERS-------------------------------------start
CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    username character varying COLLATE pg_catalog."default" NOT NULL,
    email character varying COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    age integer,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

CREATE SEQUENCE IF NOT EXISTS public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY users.id;

ALTER SEQUENCE public.users_id_seq
    OWNER TO postgres;
-- -------------------------------------TABLE::USERS---------------------------------------end
-- -------------------------------------TABLE::PHOTOS------------------------------------start
CREATE TABLE IF NOT EXISTS public.photos
(
    id bigint NOT NULL DEFAULT nextval('photos_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    title character varying COLLATE pg_catalog."default" NOT NULL,
    caption text COLLATE pg_catalog."default",
    photo_url text COLLATE pg_catalog."default",
    user_id bigint REFERENCES users (id),
    CONSTRAINT photos_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.photos
    OWNER to postgres;

CREATE SEQUENCE IF NOT EXISTS public.photos_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY photos.id;

ALTER SEQUENCE public.photos_id_seq
    OWNER TO postgres;
-- -------------------------------------TABLE::PHOTOS--------------------------------------end
-- -------------------------------------TABLE::COMMENTS----------------------------------start
CREATE TABLE IF NOT EXISTS public.comments
(
    id bigint NOT NULL DEFAULT nextval('comments_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    message text COLLATE pg_catalog."default",
    user_id bigint REFERENCES users (id),
    photo_id bigint REFERENCES photos (id),
    CONSTRAINT comments_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.comments
    OWNER to postgres;

CREATE SEQUENCE IF NOT EXISTS public.comments_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY comments.id;

ALTER SEQUENCE public.comments_id_seq
    OWNER TO postgres;
-- -------------------------------------TABLE::COMMENTS------------------------------------end
-- -------------------------------------TABLE::SOCIAL-MEDIAS-----------------------------start
CREATE TABLE IF NOT EXISTS public.socialmedia
(
    id bigint NOT NULL DEFAULT nextval('socialmedia_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    name character varying COLLATE pg_catalog."default" NOT NULL,
    social_media_url text COLLATE pg_catalog."default",
    user_id bigint REFERENCES users (id),
    CONSTRAINT socialmedia_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.socialmedia
    OWNER to postgres;

CREATE SEQUENCE IF NOT EXISTS public.socialmedia_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY socialmedia.id;

ALTER SEQUENCE public.socialmedia_id_seq
    OWNER TO postgres;
-- -------------------------------------TABLE::SOCIAL-MEDIAS-------------------------------end
-- -------------------------------------DB-------------------------------------------------end

