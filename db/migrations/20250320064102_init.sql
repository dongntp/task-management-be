-- Create enum type "role"
CREATE TYPE "public"."role" AS ENUM ('Employee', 'Employer');
-- Create enum type "status"
CREATE TYPE "public"."status" AS ENUM ('Pending', 'InProgress', 'Completed');
-- Create "account" table
CREATE TABLE "public"."account" ("username" text NOT NULL, "password" text NOT NULL, "role" "public"."role" NOT NULL, "active" boolean NOT NULL DEFAULT true, "last_updated" timestamp NOT NULL DEFAULT now(), PRIMARY KEY ("username"));
-- Create "task" table
CREATE TABLE "public"."task" ("id" text NOT NULL, "title" text NOT NULL, "assignee" text NULL, "description" text NULL, "status" "public"."status" NOT NULL DEFAULT 'Pending', "created_at" timestamp NOT NULL DEFAULT now(), PRIMARY KEY ("id"), CONSTRAINT "task_account_fk" FOREIGN KEY ("assignee") REFERENCES "public"."account" ("username") ON UPDATE NO ACTION ON DELETE SET NULL);
