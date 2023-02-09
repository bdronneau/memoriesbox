BEGIN;
CREATE SEQUENCE mbox.memories_seq;
CREATE TABLE mbox.memories (
	id integer NOT NULL DEFAULT nextval('mbox.memories_seq'),
  author text NOT NULL,
	content text NOT NULL,
  append timestamp with time zone NOT NULL,
	CONSTRAINT pk_memories_id PRIMARY KEY (id)
);
ALTER SEQUENCE mbox.memories_seq OWNED BY mbox.memories.id;
COMMIT;
