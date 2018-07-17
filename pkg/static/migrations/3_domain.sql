-- +migrate Up

CREATE TABLE "namespace" (
  "id"         UUID        NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "account_id" UUID        NOT NULL,
  "name"       VARCHAR(32) NOT NULL,
  "created_at" TIMESTAMP   NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP   NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP   NULL                 DEFAULT NULL,
  UNIQUE ("account_id", "name")
);

CREATE TABLE "link" (
  "id"         UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "account_id" UUID         NOT NULL,
  "name"       VARCHAR(256) NOT NULL,
  "created_at" TIMESTAMP    NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP    NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP    NULL                 DEFAULT NULL,
  UNIQUE ("account_id", "name")
);

CREATE TABLE "alias" (
  "id"           UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "link_id"      UUID         NOT NULL,
  "namespace_id" UUID         NOT NULL,
  "urn"          VARCHAR(512) NOT NULL,
  "created_at"   TIMESTAMP    NOT NULL             DEFAULT now(),
  "updated_at"   TIMESTAMP    NULL                 DEFAULT NULL,
  "deleted_at"   TIMESTAMP    NULL                 DEFAULT NULL,
  UNIQUE ("namespace_id", "urn")
);

CREATE TABLE "target" (
  "id"         UUID          NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "link_id"    UUID          NOT NULL,
  "uri"        VARCHAR(1024) NOT NULL,
  "rule"       JSONB         NULL                 DEFAULT NULL,
  "b_rule"     BYTEA         NULL                 DEFAULT NULL,
  "created_at" TIMESTAMP     NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP     NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP     NULL                 DEFAULT NULL,
  UNIQUE ("link_id", "uri")
);

CREATE TRIGGER "namespace_updated"
  BEFORE UPDATE
  ON "namespace"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "link_updated"
  BEFORE UPDATE
  ON "link"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "alias_updated"
  BEFORE UPDATE
  ON "alias"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "target_updated"
  BEFORE UPDATE
  ON "target"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();



-- +migrate Down

DROP TRIGGER "target_updated"
ON "target";

DROP TRIGGER "alias_updated"
ON "alias";

DROP TRIGGER "link_updated"
ON "link";

DROP TRIGGER "namespace_updated"
ON "namespace";

DROP TABLE "target";

DROP TABLE "alias";

DROP TABLE "link";

DROP TABLE "namespace";
