DROP TABLE IF EXISTS "table_truncate1";
DROP SEQUENCE IF EXISTS table_truncate1_id_seq;
CREATE SEQUENCE table_truncate1_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."table_truncate1" (
    "id" integer DEFAULT nextval('table_truncate1_id_seq') NOT NULL,
    CONSTRAINT "table_truncate1_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "table_truncate1" ("id") VALUES (1), (2), (3);

DROP TABLE IF EXISTS "table_truncate2";
DROP SEQUENCE IF EXISTS table_truncate2_id_seq;
CREATE SEQUENCE table_truncate2_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."table_truncate2" (
    "id" integer DEFAULT nextval('table_truncate2_id_seq') NOT NULL,
    "delete_me" boolean NOT NULL,
    CONSTRAINT "table_truncate2_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "table_truncate2" ("id", "delete_me") VALUES
(1,	't'),
(2,	't'),
(3,	'f');

DROP TABLE IF EXISTS "table_update";
DROP SEQUENCE IF EXISTS table_update_id_seq;
CREATE SEQUENCE table_update_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."table_update" (
    "id" integer DEFAULT nextval('table_update_id_seq') NOT NULL,
    "col_string" character varying,
    "col_bool" boolean,
    "col_int" integer,
    CONSTRAINT "table_update_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "table_update" ("id", "col_string", "col_bool", "col_int") VALUES
(1,	'foo', 't',	1),
(2,	'bar', 'f', 2),
(3,	NULL, NULL,	NULL);
