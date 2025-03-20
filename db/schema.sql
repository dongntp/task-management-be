CREATE TYPE role AS ENUM (
  'Employee',
  'Employer'
);

CREATE TYPE status AS ENUM (
  'Pending',
  'InProgress',
  'Completed'
);

CREATE TABLE IF NOT EXISTS account (
  username TEXT NOT NULL PRIMARY KEY,
  password TEXT NOT NULL,
  role role NOT NULL,
  active BOOLEAN NOT NULL DEFAULT true,

  last_updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS task (
  id TEXT NOT NULL PRIMARY KEY,
  title TEXT NOT NULL,
  assignee TEXT,
  description TEXT,
  status status NOT NULL DEFAULT 'Pending',

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT task_account_fk FOREIGN KEY (assignee) REFERENCES account (username) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_task_creation_date ON task (created_at);
