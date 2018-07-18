DO $$
DECLARE   demoAccount "account"."id"%TYPE := '00000000-0000-4000-8000-000000000001';
  DECLARE globalNS    "namespace"."id"%TYPE := demoAccount;
  DECLARE demoUser    "user"."id"%TYPE := '00000000-0000-4000-8000-000000000002';
  DECLARE demoToken   "token"."id"%TYPE := '00000000-0000-4000-8000-000000000003';
  DECLARE click       "link"."id"%TYPE := '00000000-0000-4000-8000-000000000004';
  DECLARE supportNS   "namespace"."id"%TYPE := '00000000-0000-4000-8000-000000000005';
  DECLARE promo       "alias"."id"%TYPE;
  DECLARE issue       "alias"."id"%TYPE;
BEGIN
  DELETE FROM "log"
  WHERE "link_id" = click;

  DELETE FROM "target"
  WHERE "link_id" = click;

  DELETE FROM "alias"
  WHERE "link_id" = click OR "namespace_id" IN (globalNS, supportNS);

  DELETE FROM "link"
  WHERE "id" = click OR "account_id" = demoAccount;

  DELETE FROM "namespace"
  WHERE "id" IN (globalNS, supportNS) OR "account_id" = demoAccount;

  DELETE FROM "token"
  WHERE "id" = demoToken OR "user_id" = demoUser;

  DELETE FROM "user"
  WHERE "id" = demoUser OR "account_id" = demoAccount;

  DELETE FROM "account"
  WHERE "id" = demoAccount;

  INSERT INTO "account" ("id", "name")
  VALUES (demoAccount, 'Demo account');

  INSERT INTO "user" ("id", "account_id", "name")
  VALUES (demoUser, demoAccount, 'Demo user');

  INSERT INTO "token" ("id", "user_id", "expired_at")
  VALUES (demoToken, demoUser, NULL);

  INSERT INTO "namespace" ("id", "account_id", "name")
  VALUES
    (globalNS, demoAccount, 'Global account namespace'),
    (supportNS, demoAccount, 'Support namespace');

  INSERT INTO "link" ("id", "account_id", "name")
  VALUES (click, demoAccount, 'Click! - Link Manager as a Service');

  INSERT INTO "alias" ("link_id", "namespace_id", "urn")
  VALUES (click, globalNS, 'github/click');

  INSERT INTO "alias" ("link_id", "namespace_id", "urn")
  VALUES (click, globalNS, 'github/click!')
  RETURNING "id"
    INTO promo;

  INSERT INTO "alias" ("link_id", "namespace_id", "urn")
  VALUES (click, supportNS, 'github/click')
  RETURNING "id"
    INTO issue;

  INSERT INTO "target" ("link_id", "uri", "rule", "b_rule")
  VALUES
    (click, 'https://github.com/kamilsk/click', '{
      "description": "Project''s source code",
      "tags": ["src"]
    }' :: JSONB, convert_to('{tag} in [src]', 'SQL_ASCII')),
    (click, 'https://kamilsk.github.io/click/', ('{
      "description": "Project''s promo page",
      "tags": ["promo"], "alias": "' || promo || '", "match": 1
    }') :: JSONB, convert_to('{tag} in [promo] or {alias} is "' || promo || '"', 'SQL_ASCII')),
    (click, 'https://github.com/kamilsk/click/issues/new', ('{
      "description": "Project''s bug tracker",
      "alias": "' || issue || '"
    }') :: JSONB, convert_to('{alias} is "' || issue || '"', 'SQL_ASCII'));
END;
$$;
