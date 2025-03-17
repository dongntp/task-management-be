-- Create enum type "role_enum"
CREATE TYPE "public"."role_enum" AS ENUM ('Employee', 'Employer');
-- Create enum type "status_enum"
CREATE TYPE "public"."status_enum" AS ENUM ('Pending', 'InProgress', 'Complete');
-- Create "account" table
CREATE TABLE "public"."account" ("username" text NOT NULL, "password" text NOT NULL, "role" "public"."role_enum" NOT NULL, "last_updated" timestamp NOT NULL DEFAULT now(), PRIMARY KEY ("username"));
-- Create "task" table
CREATE TABLE "public"."task" ("id" text NOT NULL, "title" text NOT NULL, "assignee" text NULL, "description" text NULL, "status" "public"."status_enum" NOT NULL DEFAULT 'Pending', "last_updated" timestamp NOT NULL DEFAULT now(), PRIMARY KEY ("id"), CONSTRAINT "task_account_fk" FOREIGN KEY ("assignee") REFERENCES "public"."account" ("username") ON UPDATE NO ACTION ON DELETE SET NULL);
