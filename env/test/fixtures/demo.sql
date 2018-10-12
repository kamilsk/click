DO $$
DECLARE   demoAccount "account"."id"%TYPE := '10000000-2000-4000-8000-160000000001';
  DECLARE demoUser    "user"."id"%TYPE := '10000000-2000-4000-8000-160000000002';
  DECLARE demoToken   "token"."id"%TYPE := '10000000-2000-4000-8000-160000000003';
  DECLARE globalNS    "namespace"."id"%TYPE := demoAccount;
  DECLARE supportNS   "namespace"."id"%TYPE := '10000000-2000-4000-8000-160000000004';
  DECLARE click       "link"."id"%TYPE := '10000000-2000-4000-8000-160000000005';
  DECLARE promo       "alias"."id"%TYPE;
  DECLARE issue       "alias"."id"%TYPE;
BEGIN
  TRUNCATE TABLE "account" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "token" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "user" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "link" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "target" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "namespace" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "alias" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "log" RESTART IDENTITY RESTRICT;

  INSERT INTO "account" ("id", "name")
  VALUES (demoAccount, 'Demo Account');

  INSERT INTO "user" ("id", "account_id", "name")
  VALUES (demoUser, demoAccount, 'Demo User');

  INSERT INTO "token" ("id", "user_id", "expired_at")
  VALUES (demoToken, demoUser, NULL);

  INSERT INTO "namespace" ("id", "account_id", "name")
  VALUES
    (globalNS, demoAccount, 'Global account namespace'),
    (supportNS, demoAccount, 'Support namespace');

  INSERT INTO "link" ("id", "account_id", "name")
  VALUES (click, demoAccount, 'Click! - Link Manager as a Service');

  INSERT INTO "alias" ("account_id", "link_id", "namespace_id", "urn")
  VALUES (demoAccount, click, globalNS, 'github/click');

  INSERT INTO "alias" ("account_id", "link_id", "namespace_id", "urn")
  VALUES (demoAccount, click, globalNS, 'github/click!')
  RETURNING "id"
    INTO promo;

  INSERT INTO "alias" ("account_id", "link_id", "namespace_id", "urn")
  VALUES (demoAccount, click, supportNS, 'github/click')
  RETURNING "id"
    INTO issue;

  INSERT INTO "target" ("account_id", "link_id", "uri", "rule", "b_rule")
  VALUES
    (demoAccount, click, 'https://github.com/kamilsk/click', '{
      "description": "Project''s source code",
      "tags": ["src"]
    }' :: JSONB, convert_to('{tag} in ["src"]', 'UTF8')),
    (demoAccount, click, 'https://kamilsk.github.io/click/', ('{
      "description": "Project''s promo page",
      "tags": ["promo"], "alias": "' || promo || '", "match": 1
    }') :: JSONB, convert_to('{tag} in ["promo"] or {alias} is "' || promo || '"', 'UTF8')),
    (demoAccount, click, 'https://github.com/kamilsk/click/issues/new', ('{
      "description": "Project''s bug tracker",
      "alias": "' || issue || '"
    }') :: JSONB, convert_to('{alias} is "' || issue || '"', 'UTF8'));
END;
$$;
