-- Modify "account" table
ALTER TABLE "public"."account" ADD COLUMN "active" boolean NOT NULL DEFAULT true;
