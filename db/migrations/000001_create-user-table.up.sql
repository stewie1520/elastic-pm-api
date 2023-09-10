START TRANSACTION;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid NOT NULL DEFAULT (uuid_generate_v4()),
  CONSTRAINT pk_users PRIMARY KEY (id)
);

CREATE UNIQUE INDEX ix_users_id ON "public".users (id);

COMMIT;