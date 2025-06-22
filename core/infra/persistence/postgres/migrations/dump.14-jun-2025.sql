--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5 (Debian 17.5-1.pgdg120+1)
-- Dumped by pg_dump version 17.5 (Debian 17.5-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: update_updated_at_column(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_updated_at_column() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, created_at, updated_at, deleted_at) FROM stdin;
1	Alice Johnson	alice.johnson@example.com	hashed_password_1	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
2	Bob Smith	bob.smith@example.com	hashed_password_2	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
3	Charlie Davis	charlie.davis@example.com	hashed_password_3	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
4	Diana Carter	diana.carter@example.com	hashed_password_4	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
5	Ethan Williams	ethan.williams@example.com	hashed_password_5	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
6	Fiona Lewis	fiona.lewis@example.com	hashed_password_6	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
7	George Miller	george.miller@example.com	hashed_password_7	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
8	Hannah Brown	hannah.brown@example.com	hashed_password_8	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
9	Ian Thompson	ian.thompson@example.com	hashed_password_9	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
10	Julia Wilson	julia.wilson@example.com	hashed_password_10	2025-06-08 20:53:06.726637	2025-06-08 20:53:06.726637	\N
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 209, true);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users trigger_update_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_update_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- PostgreSQL database dump complete
--

