DO $$
DECLARE   demoAccount "account"."id"%TYPE := '10000000-2000-4000-8000-160000000001';
  DECLARE demoUser    "user"."id"%TYPE := '10000000-2000-4000-8000-160000000002';
  DECLARE demoToken   "token"."id"%TYPE := '10000000-2000-4000-8000-160000000003';
  DECLARE globalNS    "namespace"."id"%TYPE := demoAccount;
  DECLARE supportNS   "namespace"."id"%TYPE := '10000000-2000-4000-8000-160000000004';
  DECLARE click       "link"."id"%TYPE := '10000000-2000-4000-8000-160000000005';
  DECLARE aliasIssue  "alias"."id"%TYPE := '10000000-2000-4000-8000-160000000006';
  DECLARE aliasPromo  "alias"."id"%TYPE := '10000000-2000-4000-8000-160000000007';
  DECLARE aliasSrc    "alias"."id"%TYPE := '10000000-2000-4000-8000-160000000008';
  DECLARE targetIssue "target"."id"%TYPE := '10000000-2000-4000-8000-160000000009';
  DECLARE targetPromo "target"."id"%TYPE := '10000000-2000-4000-8000-160000000010';
  DECLARE targetSrc   "target"."id"%TYPE := '10000000-2000-4000-8000-160000000011';
BEGIN
  TRUNCATE TABLE "account" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "token" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "user" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "link" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "namespace" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "alias" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "target" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "event" RESTART IDENTITY RESTRICT;

  INSERT INTO "account" ("id", "name") VALUES (demoAccount, 'Demo Account');

  INSERT INTO "user" ("id", "account_id", "name") VALUES (demoUser, demoAccount, 'Demo User');

  INSERT INTO "token" ("id", "user_id", "expired_at") VALUES (demoToken, demoUser, NULL);

  INSERT INTO "link" ("id", "account_id", "name")
  VALUES (click, demoAccount, 'Click! - Link Manager as a Service');

  INSERT INTO "namespace" ("id", "account_id", "name")
  VALUES (globalNS, demoAccount, 'Global Namespace'),
         (supportNS, demoAccount, 'Support Namespace');

  INSERT INTO "alias" ("id", "account_id", "link_id", "namespace_id", "urn")
  VALUES (aliasSrc, demoAccount, click, globalNS, 'github/click');

  INSERT INTO "alias" ("id", "account_id", "link_id", "namespace_id", "urn")
  VALUES (aliasPromo, demoAccount, click, globalNS, 'github/click!');

  INSERT INTO "alias" ("id", "account_id", "link_id", "namespace_id", "urn")
  VALUES (aliasIssue, demoAccount, click, supportNS, 'github/click');

  INSERT INTO "target" ("id", "account_id", "link_id", "url", "rule", "b_rule")
  VALUES (targetSrc, demoAccount, click, 'https://github.com/kamilsk/click', '{
      "description": "Project''s source code"
    }' :: JSONB, NULL),
         (targetPromo, demoAccount, click, 'https://kamilsk.github.io/click/', ('{
      "description": "Project''s promo page",
      "tags": ["promo"], "alias": "' || aliasPromo || '", "match": 1
    }') :: JSONB, convert_to('{tag} in ["promo"] or {alias} is "' || aliasPromo || '"', 'UTF8')),
         (targetIssue, demoAccount, click, 'https://github.com/kamilsk/click/issues/new', ('{
      "description": "Project''s bug tracker",
      "alias": "' || aliasIssue || '"
    }') :: JSONB, convert_to('{alias} is "' || aliasIssue || '"', 'UTF8'));
END;
$$;
