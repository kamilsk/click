-- +migrate Up

CREATE TABLE "event" (
  "id"           BIGSERIAL PRIMARY KEY,
  "account_id"   UUID          NOT NULL,
  "namespace_id" UUID          NOT NULL,
  "link_id"      UUID          NULL DEFAULT NULL,
  "alias_id"     UUID          NULL DEFAULT NULL,
  "target_id"    UUID          NULL DEFAULT NULL,
  "code"         INTEGER       NOT NULL,
  "url"          VARCHAR(1024) NOT NULL,
  "identifier"   UUID          NOT NULL,
  "context"      JSONB         NOT NULL,
  "created_at"   TIMESTAMP     NOT NULL DEFAULT now()
);

CREATE TRIGGER "immutable_event"
  BEFORE UPDATE
  ON "event"
  FOR EACH ROW EXECUTE PROCEDURE ignore_update();



-- +migrate Down

DROP TRIGGER "immutable_event"
ON "event";

DROP TABLE "event";
