DO $$
DECLARE   demoUser "link"."user"%TYPE := '00000000-0000-4000-8000-000000000000';
  DECLARE click    "link"."id"%TYPE := 'a382922d-b615-4227-b598-6d3633c397aa';
  DECLARE promo    "alias"."id"%TYPE;
  DECLARE issue    "alias"."id"%TYPE;
BEGIN
  DELETE FROM "log"
  WHERE "link_id" = click;

  DELETE FROM "target"
  WHERE "link_id" = click;

  DELETE FROM "alias"
  WHERE "link_id" = click;

  DELETE FROM "link"
  WHERE "id" = click;

  INSERT INTO "link" ("id", "user", "name")
  VALUES (click, demoUser, 'Click! - Link Manager as a Service');

  INSERT INTO "alias" ("link_id", "urn")
  VALUES (click, 'github/click');

  INSERT INTO "alias" ("link_id", "urn")
  VALUES (click, 'github/click!')
  RETURNING "id"
    INTO promo;

  INSERT INTO "alias" ("link_id", "namespace", "urn")
  VALUES (click, 'issue', 'github/click')
  RETURNING "id"
    INTO issue;

  INSERT INTO "target" ("link_id", "uri", "rule")
  VALUES
    (click, 'https://github.com/kamilsk/click', '{
      "description": "Project''s source code",
      "tags": [
        "src"
      ]
    }' :: JSONB),
    (click, 'https://kamilsk.github.io/click/', ('{
      "description": "Project''s promo page", "tags": ["promo"], "alias": ' || promo || ', "match": 1
    }') :: JSONB),
    (click, 'https://github.com/kamilsk/click/issues/new', ('{
      "description": "Project''s bug tracker", "alias": ' || issue || '
    }') :: JSONB);
END;
$$;
