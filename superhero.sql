-- Table: public.superhero

-- DROP TABLE public.superhero;

CREATE TABLE public.superhero
(
    uuid character varying(200) COLLATE pg_catalog."default" NOT NULL,
    name character varying(200) COLLATE pg_catalog."default" NOT NULL,
    fullname character varying(200) COLLATE pg_catalog."default" NOT NULL,
    intelligence character varying(200) COLLATE pg_catalog."default" NOT NULL,
    power character varying(200) COLLATE pg_catalog."default" NOT NULL,
    occupation character varying(200) COLLATE pg_catalog."default" NOT NULL,
    image character varying(200) COLLATE pg_catalog."default" NOT NULL,
    group_affiliation character varying(200) COLLATE pg_catalog."default",
    number_relatives integer,
    CONSTRAINT superhero_pkey PRIMARY KEY (uuid)
)

TABLESPACE pg_default;

ALTER TABLE public.superhero
    OWNER to postgres;