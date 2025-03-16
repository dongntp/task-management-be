CREATE TYPE role_enum AS ENUM (
  'Employee',
  'Employer'
);

CREATE TYPE status_enum AS ENUM (
  'Pending',
  'InProgress',
  'Complete'
);

CREATE TABLE IF NOT EXISTS account (
  id TEXT NOT NULL PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  role role_enum NOT NULL,

  last_updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS task (
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  assignee TEXT,
  description TEXT,
  status status_enum NOT NULL DEFAULT 'Pending',

  last_updated TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT task_account_fk FOREIGN KEY (assignee) REFERENCES account (id) ON DELETE SET NULL
);
