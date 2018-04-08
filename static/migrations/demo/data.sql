-- +migrate Up

-- +migrate StatementBegin
DO $$
DECLARE   demoUser   UUID := uuid_generate_v4();
  DECLARE click      UUID := 'a382922d-b615-4227-b598-6d3633c397aa';
  DECLARE egg        UUID := '09579e02-e076-489a-9741-f4cf2a67ad40';
  DECLARE forma      UUID := '0d49a516-abea-430c-87a3-17992b8c6e45';
  DECLARE retry      UUID := '19ac200a-ca74-48ca-8086-511f23d0c43e';
  DECLARE passport   UUID := '4b1f1033-dd7c-452d-94bc-fe78591b7e66';
  DECLARE semaphore  UUID := '7f9b8b84-977f-4fdb-a4df-fb17919fa897';
  DECLARE clickPromo "alias"."id"%TYPE;
BEGIN
  INSERT INTO "link" ("user", "id", "name") VALUES
    (demoUser, click, 'Click! - Link Manager as a Service'),
    (demoUser, egg, 'egg - extended go get - alternative for standard "go get" with a few little but useful features'),
    (demoUser, forma, 'Forma - Data Collector as a Service'),
    (demoUser, retry, 'Functional mechanism based on channels to perform actions repetitively until successful'),
    (demoUser, passport, 'Passport - Person Identifier as a Service'),
    (demoUser, semaphore, 'Semaphore pattern implementation with timeout of lock/unlock operations based on channels');
  INSERT INTO "alias" ("link_id", "urn") VALUES
    (click, 'github/click'),
    (egg, 'github/egg'),
    (forma, 'github/forma'),
    (retry, 'github/retry'),
    (passport, 'github/passport'),
    (semaphore, 'github/semaphore');
  INSERT INTO "alias" ("link_id", "urn") VALUES (click, 'github/click!')
  RETURNING "id"
    INTO clickPromo;
  INSERT INTO "target" ("link_id", "uri", "rule") VALUES
    (click, 'https://github.com/kamilsk/click', '{
      "description": "Project location", "tags": ["src"]
    }'),
    (click, 'https://kamilsk.github.io/click/', ('{
      "description": "Promotion page", "tags": ["promo"], "alias": ' || clickPromo || ', "match": 1
    }') :: JSONB),
    (egg, 'https://github.com/kamilsk/egg', NULL),
    (forma, 'https://github.com/kamilsk/form-api', '{
      "description": "Project location", "tags": ["src"]
    }'),
    (forma, 'https://kamilsk.github.io/form-api/', '{
      "description": "Promotion page", "tags": ["promo"], "conditions": {"type": "promo"}, "match": 1
    }'),
    (retry, 'https://github.com/kamilsk/retry', NULL),
    (passport, 'https://github.com/kamilsk/passport', '{
      "description": "Project location", "tags": ["src"]
    }'),
    (passport, 'https://kamilsk.github.io/passport/', '{
      "description": "Promotion page", "tags": ["promo"]
    }'),
    (semaphore, 'https://github.com/kamilsk/semaphore', NULL);
END;
$$;
-- +migrate StatementEnd

-- +migrate Down

-- +migrate StatementBegin
DO $$
DECLARE   click     UUID := 'a382922d-b615-4227-b598-6d3633c397aa';
  DECLARE egg       UUID := '09579e02-e076-489a-9741-f4cf2a67ad40';
  DECLARE forma     UUID := '0d49a516-abea-430c-87a3-17992b8c6e45';
  DECLARE retry     UUID := '19ac200a-ca74-48ca-8086-511f23d0c43e';
  DECLARE passport  UUID := '4b1f1033-dd7c-452d-94bc-fe78591b7e66';
  DECLARE semaphore UUID := '7f9b8b84-977f-4fdb-a4df-fb17919fa897';
BEGIN
  DELETE FROM "target"
  WHERE "link_id" IN (click, egg, forma, retry, passport, semaphore);
  DELETE FROM "alias"
  WHERE "link_id" IN (click, egg, forma, retry, passport, semaphore);
  DELETE FROM "link"
  WHERE "id" IN (click, egg, forma, retry, passport, semaphore);
END;
$$;
-- +migrate StatementEnd
