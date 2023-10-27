

--
-- Name: todo_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.todo_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.todo_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: todo; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.todos (
                              id integer DEFAULT nextval('public.todo_id_seq'::regclass) NOT NULL,
                              todo character varying(255),
                              todo_active integer DEFAULT 0,
                              created_at timestamp without time zone,
                              updated_at timestamp without time zone
);


ALTER TABLE public.todos OWNER TO postgres;


--
-- Name: todo todo_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.todos
    ADD CONSTRAINT todo_pkey PRIMARY KEY (id);






