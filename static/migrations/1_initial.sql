-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate StatementBegin
CREATE FUNCTION update_timestamp()
  RETURNS TRIGGER AS $update_timestamp$
BEGIN
  IF NEW.* IS DISTINCT FROM OLD.*
  THEN
    NEW.updated_at := current_timestamp;
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$update_timestamp$
LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TYPE STATUS AS ENUM ('active', 'hidden');

CREATE TABLE "link" (
  "id"         UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user"       UUID         NOT NULL,
  "name"       VARCHAR(512) NOT NULL,
  "status"     STATUS       NOT NULL             DEFAULT 'active',
  "created_at" TIMESTAMP    NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP    NULL                 DEFAULT NULL
);

CREATE TRIGGER "link_updated"
  BEFORE UPDATE
  ON "link"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TABLE "alias" (
  "id"         SERIAL PRIMARY KEY,
  "link_id"    UUID         NOT NULL,
  "namespace"  VARCHAR(128) NOT NULL             DEFAULT 'global',
  "urn"        VARCHAR(512) NOT NULL,
  "created_at" TIMESTAMP    NOT NULL             DEFAULT now(),
  "deleted_at" TIMESTAMP    NULL                 DEFAULT NULL,
  UNIQUE ("namespace", "urn")
);

CREATE TABLE "target" (
  "id"         SERIAL PRIMARY KEY,
  "link_id"    UUID          NOT NULL,
  "uri"        VARCHAR(1024) NOT NULL,
  "rule"       JSONB         NULL     DEFAULT NULL,
  "created_at" TIMESTAMP     NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP     NULL     DEFAULT NULL
);

CREATE TRIGGER "target_updated"
  BEFORE UPDATE
  ON "target"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TABLE "log" (
  "id"         BIGSERIAL PRIMARY KEY,
  "link_id"    UUID          NOT NULL,
  "alias_id"   INTEGER       NOT NULL,
  "target_id"  INTEGER       NOT NULL,
  "uri"        VARCHAR(1024) NOT NULL,
  "code"       SMALLINT      NOT NULL,
  "context"    JSONB         NOT NULL,
  "created_at" TIMESTAMP     NOT NULL DEFAULT now()
);

-- +migrate Down

DROP TABLE "log";

DROP TRIGGER "target_updated"
ON "target";

DROP TABLE "target";

DROP TABLE "alias";

DROP TRIGGER "link_updated"
ON "link";

DROP TABLE "link";

DROP TYPE STATUS;

DROP FUNCTION update_timestamp();
